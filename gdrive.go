// Package gdrivestore is a Google Drive driver for togo storage. It implements
// togo.Storage and overrides the default filesystem storage when installed.
// Files are keyed by their path (used as the Drive file name) within an optional
// parent folder.
//
//	togo install togo-framework/storage-gdrive
//
// Env: GOOGLE_APPLICATION_CREDENTIALS (service-account JSON, required),
// GDRIVE_FOLDER_ID (optional parent folder).
package gdrivestore

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("storage-gdrive", togo.PriorityService+10, func(k *togo.Kernel) error {
		creds := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
		if creds == "" {
			if k.Log != nil {
				k.Log.Warn("storage-gdrive: GOOGLE_APPLICATION_CREDENTIALS not set; skipping")
			}
			return nil
		}
		svc, err := drive.NewService(context.Background(),
			option.WithCredentialsFile(creds), option.WithScopes(drive.DriveScope))
		if err != nil {
			return fmt.Errorf("storage-gdrive: drive service: %w", err)
		}
		k.Storage = &store{svc: svc, folder: os.Getenv("GDRIVE_FOLDER_ID")}
		return nil
	})
}

type store struct {
	svc    *drive.Service
	folder string
}

func (s *store) find(name string) (string, error) {
	q := fmt.Sprintf("name = '%s' and trashed = false", strings.ReplaceAll(name, "'", `\'`))
	if s.folder != "" {
		q += fmt.Sprintf(" and '%s' in parents", s.folder)
	}
	r, err := s.svc.Files.List().Q(q).Fields("files(id)").PageSize(1).Do()
	if err != nil {
		return "", err
	}
	if len(r.Files) == 0 {
		return "", nil
	}
	return r.Files[0].Id, nil
}

func (s *store) Put(path string, data []byte) error {
	name := key(path)
	id, err := s.find(name)
	if err != nil {
		return err
	}
	if id != "" {
		_, err = s.svc.Files.Update(id, &drive.File{}).Media(bytes.NewReader(data)).Do()
		return err
	}
	f := &drive.File{Name: name}
	if s.folder != "" {
		f.Parents = []string{s.folder}
	}
	_, err = s.svc.Files.Create(f).Media(bytes.NewReader(data)).Do()
	return err
}

func (s *store) Get(path string) ([]byte, error) {
	id, err := s.find(key(path))
	if err != nil {
		return nil, err
	}
	if id == "" {
		return nil, fmt.Errorf("storage-gdrive: not found: %s", path)
	}
	resp, err := s.svc.Files.Get(id).Download()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (s *store) Delete(path string) error {
	id, err := s.find(key(path))
	if err != nil {
		return err
	}
	if id == "" {
		return nil
	}
	return s.svc.Files.Delete(id).Do()
}

func (s *store) Path(path string) string {
	id, _ := s.find(key(path))
	if id == "" {
		return ""
	}
	return "https://drive.google.com/file/d/" + id + "/view"
}

func key(p string) string { return strings.TrimPrefix(p, "/") }

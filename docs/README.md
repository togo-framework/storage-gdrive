# storage-gdrive — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package gdrivestore is a Google Drive driver for togo storage. It implements
togo.Storage and overrides the default filesystem storage when installed.
Files are keyed by their path (used as the Drive file name) within an optional
parent folder.

	togo install togo-framework/storage-gdrive

## Install

```bash
togo install togo-framework/storage-gdrive
```

Set `STORAGE_DRIVER=gdrive`.

## Configuration

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `G` | _see provider docs_ |
| `GDRIVE_FOLDER_ID` | _see provider docs_ |
| `GOOGLE_APPLICATION_CREDENTIALS` | _see provider docs_ |

## Usage

```go
st := k.Storage
st.Put(ctx, "path/file.txt", data)
b, _ := st.Get(ctx, "path/file.txt")
url := st.Path("path/file.txt")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/storage-gdrive
- README: ../README.md

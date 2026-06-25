# storage-gdrive — documentation

togo storage driver

## Overview

Package gdrivestore is a Google Drive driver for togo storage. It implements
togo.Storage and overrides the default filesystem storage when installed.
Files are keyed by their path (used as the Drive file name) within an optional
parent folder.

## Install

```bash
togo install togo-framework/storage-gdrive
```

Set `STORAGE_DRIVER=gdrive`.

## Configuration

Environment variables read by this plugin (extracted from the source — see the gateway/provider docs for each value):

| Env var |
|---|
| `GDRIVE_FOLDER_ID` |
| `GOOGLE_APPLICATION_CREDENTIALS` |

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
- Full README: ../README.md

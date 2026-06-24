# storage-gdrive

A **Google Drive** driver for [togo](https://to-go.dev) storage. Implements
`togo.Storage` and overrides the default filesystem storage when installed. Files
are keyed by their path (the Drive file name) inside an optional parent folder.

## Install

```bash
togo install togo-framework/storage-gdrive
```

## Configure (`.env`)

```ini
GOOGLE_APPLICATION_CREDENTIALS=/path/to/service-account.json
GDRIVE_FOLDER_ID=optional-parent-folder-id
```

Share the target Drive folder with the service-account email. If credentials are
unset the plugin no-ops and the default storage stays active.

MIT © togo-framework

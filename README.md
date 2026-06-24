<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/storage-gdrive</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/storage-gdrive"><img src="https://pkg.go.dev/badge/github.com/togo-framework/storage-gdrive.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/storage-gdrive
```

<!-- /togo-header -->

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

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->

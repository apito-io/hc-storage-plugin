# hc-storage-plugin

AWS S3, Google Cloud Storage, and local file storage for the Apito API platform.

## Features

- AWS S3 integration
- Google Cloud Storage
- Local file storage
- Image optimization
- File upload validation

## Installation

```bash
go build -o hc-storage-plugin .
```

## Configuration

See `config.yml` for environment variables: `STORAGE_PROVIDER`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_S3_BUCKET`.

## How to Develop Plugins

See the [CDN Plugin Development Guide](https://github.com/apito-io/plugins/blob/main/CDN-PLUGIN-DEVELOPMENT.md).

## How to Submit Your Plugin

1. Create your plugin repo with naming convention `hc-{name}-plugin`
2. Ensure: `config.yml`, `README.md`, `logo.png`
3. Submit a PR to [github.com/apito-io/plugins](https://github.com/apito-io/plugins) adding your entry to `plugins.json`

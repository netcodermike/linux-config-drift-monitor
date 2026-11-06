# Linux Config Drift Monitor

Compare configuration file contents against an explicit SHA-256 baseline.

[![Build and verify](https://github.com/netcodermike/linux-config-drift-monitor/actions/workflows/verify.yml/badge.svg)](https://github.com/netcodermike/linux-config-drift-monitor/actions/workflows/verify.yml)

**Language:** Go · **Platform:** Linux-focused · **App type:** Linux configuration change detector

First-release learning implementation, with explicit scope below. This is one independently runnable app, not a placeholder in a combined repository.

## Features

- Recursive regular-file hashing
- Added/changed/deleted reports
- Exclusive baseline creation
- No configuration writes

## Requirements

Go 1.22+ and the documented Linux facilities. No external Go modules.

## Quick start

```sh
git clone https://github.com/netcodermike/linux-config-drift-monitor.git
cd linux-config-drift-monitor
go build -o app .
./app baseline /path/to/config /another/path/baseline.json
./app check /path/to/config /another/path/baseline.json
```

Commands run from the repository root. Lines marked for a particular operating system apply only there. Language tools are prerequisites, not installed by the app. Close desktop windows normally; terminal/server apps stop with Ctrl+C unless their usage notes say otherwise.

## Example workflow

Create a small folder, run ./app baseline FOLDER ../baseline.json, edit a file, then ./app check FOLDER ../baseline.json. Confirm CHANGED and later DELETED/ADDED reports.

## Architecture, configuration, and data

main.go walks regular files, skips symlinks, hashes contents, and stores relative-path hashes in a JSON manifest outside the scanned folder. Individual files are limited to 4 MiB and scans to 10,000 files. baseline refuses an existing manifest. check sorts drift messages and exits nonzero on drift or errors. Tests modify/delete synthetic files.

## Verification

Build with the quick-start command where needed, then run:

```sh
go test ./...
```

Manual acceptance: Create a small folder, run ./app baseline FOLDER ../baseline.json, edit a file, then ./app check FOLDER ../baseline.json. Confirm CHANGED and later DELETED/ADDED reports. Also test malformed inputs and an unwritable output destination.

See [validation record](docs/VALIDATION.md) for observed checks and their limits. CI results verify only what the workflow runs; they do not establish native GUI correctness or production readiness.

## Limitations and next steps

Content only: no owner/mode/ACL checks, scheduled service, automatic remediation, encryption, or signed baseline. Active writers can make a scan inconsistent. The manifest must stay outside the source tree and should be protected separately. Baseline creation may leave a partial file after a write failure.

## Troubleshooting

Use the exact platform/toolchain above. Preserve originals and use synthetic fixtures first. Errors are displayed in the UI or stderr. A headless machine can run the model tests but cannot verify native GUI/audio behavior. Output paths are chosen explicitly; keep backups of important data.

## Source map

- `go.mod`
- `main.go`
- `main_test.go`

## Development and license

Read [contributing](CONTRIBUTING.md) before changing the app, and [security/privacy](SECURITY.md) before supplying private data. Use synthetic fixtures for tests. No telemetry, paid API key, or account is required by the app itself. Dependencies, when present, are stated in Requirements.

[MIT license](LICENSE), copyright 2026 netcodermike.

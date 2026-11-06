# Validation

Authoring date: November 6, 2026. Host: Windows. Target: Linux-focused.

## Local result

Target-platform source CI is recorded after publication. This host has Python, Node.js, and .NET; other native toolchains are verified through CI. No unobserved GUI/device behavior is claimed.

## Reproducible checks

The GitHub Actions workflow builds/runs the following on `ubuntu-latest`:

```sh
gofmt -w main.go main_test.go
go vet ./...
go test -race ./...
go build -o app .
```

## Manual acceptance

Create a small folder, run ./app baseline FOLDER ../baseline.json, edit a file, then ./app check FOLDER ../baseline.json. Confirm CHANGED and later DELETED/ADDED reports. Also test malformed inputs and an unwritable output destination.

A workflow file is not evidence of a passing run. Observed CI run links and source commits are recorded after publication. Native GUI, sound, sleep/resume, and platform permissions require direct target-platform acceptance even when model tests pass.

## Observed CI result

Passed: [target-platform build/check](https://github.com/netcodermike/linux-config-drift-monitor/actions/runs/42205349975). Source commit: `edb7735e42f5da7eaa92803ade71e6c350cc79a3`. Later documentation-only changes do not alter the tested implementation. The reproducible commands above identify test scope; native GUI/audio/permissions are not implied by compilation.

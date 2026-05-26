---
name: degg-release
description: Release guidance for GoReleaser, GitHub Actions release tags, binary archives, Devbox tooling, and CLI version metadata in degg.
---

# Degg Release

Use this skill when changing GoReleaser, GitHub Actions, Devbox, `.justfile`,
or CLI version metadata.

## Release Shape

Releases are tag-driven. Pushing a `v*` tag runs
`.github/workflows/release.yml`, which uses GoReleaser v2 to publish binary
archives for `degg`.

This repo publishes CLI binaries only. Do not add Docker or GHCR publishing
unless a future task introduces a container runtime requirement.

## Version Metadata

Release builds inject these variables in `cmd/degg/cli`:

- `ProgramVersion`
- `ProgramCommitSHA`
- `ProgramBuildTime`

The local `just build degg` path injects development metadata from the current
commit and UTC timestamp.

## Validation

Run:

```bash
goreleaser check
goreleaser release --snapshot --clean
```

If GoReleaser is not installed globally, use:

```bash
devbox run release-check
devbox run release-snapshot
```

## Risk Checks

- Keep `go-version-file: go.mod` in the workflow unless the module policy
  changes.
- Keep `CGO_ENABLED=0`; `degg` currently has no cgo dependency.
- Do not publish release artifacts from local snapshot validation.

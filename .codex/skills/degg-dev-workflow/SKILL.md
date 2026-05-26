---
name: degg-dev-workflow
description: Local development workflow for degg. Use when orienting in the workspace, choosing just/go/GoReleaser commands, or validating general changes.
---

# Degg Dev Workflow

Use this skill for repo orientation, command selection, and everyday
implementation hygiene across the workspace.

## Repo Shape

`degg` is a single-module Go project. The main binary is `cmd/degg`, supported
by repo-local packages in `internal/`.

- `cmd/degg/` - CLI entrypoint, flags, validation, and program metadata.
- `internal/declaration/` - declaration format parsing and validation.
- `internal/generator/` - code generation and template execution.
- `internal/system/` - filesystem helpers.
- `examples/` - declaration examples and generated outputs.

## Commands

Prefer `devbox shell` when available.

```bash
just list
just build degg
just run degg -- -i examples/color/color.yml -o /tmp/color_enum.go
go test ./...
```

Use `just test` when you want to match the repository recipe that excludes
examples.

## Implementation Rules

- Prefer existing package boundaries over new abstractions.
- Keep validation behavior in `internal/declaration/` and CLI flag/path checks
  in `cmd/degg/`.
- Keep generator template changes synchronized with tests and examples.
- Do not hand-edit local output such as `.devbox/`, `bin/`, `dist/`, or
  temporary generated scratch files.

## Validation

Run the smallest useful checks first. Before finishing a meaningful change,
prefer:

```bash
go test ./...
just build degg
```

For release behavior, add:

```bash
goreleaser check
goreleaser release --snapshot --clean
```

## Commit Hygiene

Recent history uses concise subjects such as `feat: ...` and `chore: ...`.
Keep commits focused and do not stage unrelated user changes.

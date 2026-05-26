# Repository Guidelines

## Project Structure & Module Organization

`degg` is a Go CLI that reads YAML or JSON enum declarations and generates
formatted Go enum code.

- `cmd/degg/` contains the executable entrypoint, CLI wiring, and flag
  validation.
- `internal/declaration/` parses, formats, and validates declaration files.
- `internal/generator/` renders Go code from declarations using templates.
- `internal/generator/templates/` contains generator templates.
- `internal/system/` holds local filesystem helpers.
- `examples/` contains source declarations and expected generated Go examples.
- `.github/workflows/release.yml`, `.goreleaser.yaml`, `devbox.json`, and
  `.justfile` define build, test, and release workflows.

## Build, Test, And Development Commands

Use `devbox shell` when possible. It pins Go, `just`, and GoReleaser.

- `just list` lists binaries under `cmd/`.
- `just build degg` builds `cmd/degg` into `bin/degg` with version metadata.
- `just run degg -- -i <input> -o <output>` builds and runs the CLI.
- `just test` runs tests outside `examples/`.
- `go test ./...` is the default validation command.
- `goreleaser check` validates release configuration.
- `goreleaser release --snapshot --clean` builds release artifacts without
  publishing.

## Generator Contract

Declaration files must provide `name`, `package`, `type`, and exactly one of
`values` or `named-values`. The supported enum types are `string` and `int`.

Generated code is formatted with `go/format` before it is written. If formatting
fails, the unformatted output is written to the requested path so the generated
source can be inspected.

Keep template changes coordinated with tests in `internal/generator/` and with
examples under `examples/`.

## Release

Releases are tag-driven. Pushing a `v*` tag runs
`.github/workflows/release.yml`, which invokes GoReleaser and publishes binary
archives for `degg`.

The release build injects these variables in `cmd/degg/cli`:

- `ProgramVersion`
- `ProgramCommitSHA`
- `ProgramBuildTime`

Do not add Docker or GHCR publishing unless the repository gains a runtime use
case for containers.

## Coding Style & Safety

- Prefer small, direct Go changes that match existing package boundaries.
- Keep generated examples intentional; do not refresh them casually unless the
  generator behavior changed.
- Do not commit local output: `.devbox/`, `bin/`, `dist/`, coverage files, or
  temporary generated scratch files.
- If a task requires committing, use focused commits that match repo history,
  such as `build: ...`, `docs: ...`, `feat: ...`, `fix: ...`, or
  `refactor: ...`.

## Agent-Specific Instructions

`AGENTS.md` is the canonical instruction layer for this repository. `CLAUDE.md`
is a Claude Code pointer and quick reference. Local Codex skills live in
`.codex/skills/`; Claude Code should read those skills from `.codex/skills/`
rather than duplicating them under `.claude/skills/`. Specialist subagents are
mirrored under `.codex/agents/` and `.claude/agents/`.

When changing repository guidance, keep `AGENTS.md`, `CLAUDE.md`, skills, and
subagents consistent in the same change set.

## BRAIN - Obsidian Vault

Vault root: `/Users/upsetbit/Projects/_me/upsetbit/BRAIN`.

Use this personal Obsidian vault when the user asks to save to BRAIN, put in
BRAIN, "guarde no brain", or clearly means storing a note, capture, or memo in
the vault rather than in this project tree.

- Default generic captures go under `inbox/`.
- Investigation work goes under `projects/<slug>/` with frontmatter.
- `knowledge/` is read-only. To propose durable knowledge, write a proposal
  under `inbox/` with `type: proposal`.
- Daily notes under `daily/` are append-only.
- Prefer Markdown. If no filename is specified, choose a sensible slug.
- Include YAML frontmatter with `type`, `created`, `agent`, and `source`.
- Never relocate the vault or assume a different root without explicit user
  instruction.
- Do not overwrite important notes without confirmation if the target path
  already exists and has content.

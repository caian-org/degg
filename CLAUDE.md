# Claude Notes

Pointer doc for Claude Code agents working in `degg`.

**Read `AGENTS.md` first**. It is the canonical instruction layer. This file
covers only what is specific to Claude Code or to quick orientation.

## Bootstrap

1. `AGENTS.md` - canonical repository rules, structure, commands, generator
   contract, release, and agent-specific instructions.
2. `.codex/skills/degg-dev-workflow/SKILL.md` - default local workflow guidance
   for this repo.
3. A subsystem skill when relevant:
   - `.codex/skills/degg-generator/SKILL.md`
   - `.codex/skills/degg-release/SKILL.md`

`AGENTS.md` is authoritative. If `CLAUDE.md` disagrees with it, follow
`AGENTS.md` and reconcile this file in the same change set.

## Claude Code Specifics

- Use Claude Code's native subagent feature for delegated lanes when the session
  policy allows it. Specialists live under `.claude/agents/`, mirroring
  `.codex/agents/`.
- Skills are not duplicated for Claude. The canonical home is `.codex/skills/`;
  read skills from there.
- Claude Code settings live under `.claude/settings.json`. Hooks are currently
  empty.
- Use the local checkout as the source of truth. Do not use GitHub API reads as
  a substitute for inspecting files in this repo unless it is a narrow one-off
  check.
- Serialize shared-checkout mutations to one owner: edits, generated output,
  patch application, staging, committing, rebasing, branch switching, and
  pushing.
- After editing `AGENTS.md`, `.codex/skills/`, `.codex/agents/`, or
  `.claude/agents/`, update this file if the Claude-facing guidance changes.

## Quick Reference

Project purpose:

- Read YAML or JSON enum declarations.
- Validate declaration names, package names, enum types, and value sets.
- Generate formatted Go enum code with helper methods.

Command surface:

- `devbox shell` - enter the pinned development environment.
- `just list` - list binaries under `cmd/`.
- `just build degg` - build the CLI into `bin/degg`.
- `just run degg -- -i examples/color/color.yml -o /tmp/color_enum.go` - run a
  sample generation.
- `go test ./...` - default validation command.
- `goreleaser check` - validate release config.
- `goreleaser release --snapshot --clean` - local release dry run.

High-traffic paths:

- `cmd/degg/main.go` - orchestration of parse, validate, generate, format, and
  write.
- `cmd/degg/cli/` - CLI metadata and urfave/cli setup.
- `internal/declaration/` - declaration parsing and validation.
- `internal/generator/` - code rendering and templates.
- `examples/` - example inputs and expected generated files.

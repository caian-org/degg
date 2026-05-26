---
name: degg-generator
description: Generator guidance for declaration parsing, validation, enum template output, examples, and generated Go formatting in degg.
---

# Degg Generator

Use this skill when changing declaration semantics, validation, generated code,
templates, or examples.

## Declaration Contract

Supported inputs are YAML and JSON. A declaration must include:

- `name`
- `package`
- `type`
- exactly one of `values` or `named-values`

The supported enum types are `string` and `int`.

## Generator Contract

Generation is coordinated by `cmd/degg/main.go`:

1. Resolve input and output paths.
2. Identify input format.
3. Validate output path.
4. Parse and validate the declaration.
5. Run `internal/generator`.
6. Format with `go/format`.
7. Write the output file.

Keep generated helper names and error behavior stable unless the task explicitly
changes the public generated API.

## Template And Examples

Template output lives under `internal/generator/templates/`. The example
directories under `examples/` should represent real generated output. When the
template changes, update tests first and refresh examples only when their
behavior is intentionally changed.

## Risk Checks

- Case-insensitive `FromName` behavior is user-facing.
- `named-values` for `int` must remain parseable as integers.
- Invalid declarations should report useful validation errors before generation.
- Formatting failures intentionally leave unformatted generated code on disk.

## Validation

Run:

```bash
go test ./...
just run degg -- -i examples/color/color.yml -o /tmp/color_enum.go
```

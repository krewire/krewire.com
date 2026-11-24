# Contributing to krewire.com

Guidelines for contributing to **krewire.com** (`github.com/krewire/krewire.com`).

> All code, comments, docstrings, commit messages, and documentation are strictly in **English**.

## Invariants

1. **Bilingual Parity**: Technical documentation updates must be maintained in both English (`content/en/docs/`) and Indonesian (`content/id/docs/`).
2. **Pure Go & Zero CGO**: All Go code (`mcp`, `examples/todo-app`) must compile with `CGO_ENABLED=0`.
3. **No Root go.mod**: Modules reside under `mcp/` and `examples/todo-app/`.

## Quality Gates

```bash
gofmt -l .
go vet ./...
go test ./...
CGO_ENABLED=0 go build ./...
```

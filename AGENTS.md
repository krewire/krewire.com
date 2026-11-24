# AGENTS.md — krewire.com Landing & Docs Portal

Agent guide for AI agents working inside the `krewire.com` repository.

`krewire.com` is the official public website and documentation portal for Krewire.
It is 100% dogfooded using the ecosystem itself: `project.kind: site`, `.kiw` DSL,
`pages/`, `components/`, `layouts/`, and built with `krew build`.

## Multi-Module Layout

- `mcp`: Public MCP server implementation and daemon (`github.com/krewire/krewire.com/mcp`).
- `examples/todo-app`: Reference fullstack monolith application.
- Coordinated via `go.work` without root `go.mod`.

## Language

- **Write all stored artifacts in English** — pages, components, docs content, commit messages.
- If the human writes in Indonesian or another language, reply in the human's language in chat, but keep all repository files in English.

## Quality Gates

```bash
gofmt -l .
go vet ./...
go test ./...
CGO_ENABLED=0 go build ./...
```

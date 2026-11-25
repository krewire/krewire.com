# `krewire.com` — Go Modules & Examples Reference

While the primary `krewire.com` portal is a static `site` workload built with `krew build` and `.kiw` templates, the repository also contains two self-contained Go modules coordinated through [`go.work`](../../go.work):

- [**`mcp`**](./mcp.md) (`github.com/krewire/krewire.com/mcp`) — The public **Model Context Protocol (MCP)** server (`krewire-mcp-server`) that indexes the bilingual `content/` documentation tree and exposes Krewire's ecosystem, UI component catalog, architecture rules, and scaffolding tools over both `stdio` and HTTP/SSE.
- [**`todo-app`**](./todo-app.md) (`todo-app`) — The companion **full-stack Go + SQLite + `.kiw` SSR** tutorial application located in `examples/todo-app/`, demonstrating `app.Container` dependency injection, `orm` with `hub/sqlite`, CSRF-protected HTML forms, and embedded migrations/assets.

---

## Verification

From the `krewire.com/` repository root:

```bash
gofmt -l ./mcp ./examples/todo-app
go vet ./mcp/... ./examples/todo-app/...
CGO_ENABLED=0 go test ./mcp/... ./examples/todo-app/...
```

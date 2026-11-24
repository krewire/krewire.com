# krewire.com — Official Ecosystem Website & Docs Portal

Official public website and documentation portal for the Krewire digital SDLC ecosystem (published at **https://krewire.com**).

Demonstrates real-world Static Site Generation (SSG), multi-language book compilation, scoped styling, and client-side modal search with zero Node.js/npm dependencies.

The repository operates as a **Pure Multi-Module Monorepo** comprising **2 independent Go modules** linked via `go.work`:
- **`mcp` (`github.com/krewire/krewire.com/mcp`)**: Public Model Context Protocol server exposing Krewire documentation, component catalog, and ecosystem rules.
- **`examples/todo-app` (`todo-app`)**: End-to-end reference monolith application demonstrating `.kiw` views, UI components, and SQLite storage via `github.com/krewire/hub/sqlite`.

## Features

- **Unified Build Engine:** Built via `krew build`, powered by `.kiw` template engine, documentation manuscript engine, and static site generator.
- **Bilingual Documentation:** Full 1:1 parity between English (`content/en/docs/`) and Indonesian (`content/id/docs/`).
- **MCP Public Server:** Built-in SSE and stdio MCP server for IDEs.

## Quality Gates

Before submitting changes:
```bash
gofmt -l .
go vet ./...
go test ./...
CGO_ENABLED=0 go build ./...
```

## License

[MIT License](LICENSE) © 2026 Krewire Contributors

# `mcp` (`github.com/krewire/krewire.com/mcp`)

## Purpose & Workspace Role

The `mcp` module implements Krewire's public **Model Context Protocol (MCP)** server (`ProtocolVersion = "2024-11-05"`), powering both the hosted endpoints at `https://krewire.com/mcp/*` and local IDE integrations (Cursor, Claude, Antigravity) via `--stdio`. It gives AI coding assistants live, structured access to Krewire's bilingual documentation (`en` and `id`), Pop-Brutalist `.kiw` component catalog, 6-Layer Pyramid architecture rules, layering validator, and project/migration generators.

---

## Architectural Position

- **Ecosystem Layer**: Layer 5 (`krewire.com` public portal & tooling).
- **Dependencies**: Pure Go Standard Library only (`bufio`, `context`, `crypto/rand`, `encoding/hex`, `encoding/json`, `io/fs`, `net/http`, `os`, `path/filepath`, `sort`, `strings`, `sync`, `time`). Zero third-party dependencies and `CGO_ENABLED=0`.
- **Module Layout**:
  - `mcp/mcp/` — Core MCP library (`Server`, `DocIndexer`, `ComponentCatalog`, `HTTPTransport`, `StdioTransport`, ecosystem/architecture helpers).
  - `mcp/krewire-mcp-server/` — CLI/daemon entrypoint (`main.go`).

---

## Core Concepts & APIs

### MCP Server & JSON-RPC Dispatch (`server.go`)

- **`Server`** & **`NewServer(docs *DocIndexer) *Server`**: Handles JSON-RPC 2.0 `Request`s via `HandleRequest(ctx, req) *Response` (returning `nil` for `notifications/*`).
- **Supported Methods**: `initialize`, `ping`, `resources/list`, `resources/read`, `prompts/list`, `prompts/get`, `tools/list`, `tools/call`.
- **Resources**:
  - `krewire://docs/toc` — Documentation table of contents (`en`).
  - `krewire://ecosystem` — Authoritative repository & layer breakdown (`GetEcosystem`).
  - `krewire://components` — Pop-Brutalist `.kiw` component catalog.
  - `krewire://rules` — 6-Layer Pyramid and Three Pillars invariants (`GetArchitectureRules`).
- **Prompts**: `krewire-architect`, `kiw-component-builder`, `krewire-migrator`.
- **Tools (11 total)**: `search_docs`, `get_doc`, `list_docs`, `list_components`, `get_component`, `get_ecosystem`, `get_architecture_rules`, `check_layering_rule`, `generate_migration`, `scaffold_project_config`, `get_cli_commands`.

### Documentation Indexer (`docs.go`)

- **`DocIndexer`** & **`NewDocIndexer(contentFS fs.FS) (*DocIndexer, error)`**: Walks markdown files, parses YAML frontmatter (`title`, `description`, `tags`), normalizes paths (`CleanPath` strips numeric prefixes like `02-`), and builds canonical `https://krewire.com/docs/...` and `https://krewire.com/id/docs/...` URLs.
- **`(*DocIndexer).Search(query, locale, limit)`**: Tokenizes queries and scores matches across title (`+15`), tags (`+10`), description (`+8`), path (`+6`), and body occurrences (`+2` per hit, capped), returning `[]SearchResult` with excerpts.
- **`(*DocIndexer).Get(rawPath, locale)`** & **`(*DocIndexer).List(locale)`**: Resolves full paths, clean paths, directory indexes (`00-index`), and full URLs.

### Component Catalog & Ecosystem Helpers (`components.go`, `ecosystem.go`)

- **`ComponentCatalog`** & **`NewComponentCatalog()`**: Seeds metadata, props, variants, and `.kiw` examples for `Button`, `Card`, `Badge`, `Modal`, `Input`, `Navbar`, `Terminal`, `Alert`, and `Tabs`, exposed via `List(category)` and `Get(name)`.
- **`CheckLayering(from, to string) LayeringEvaluation`**: Validates downward-only imports (`Layer N -> Layer < N`) and lateral peer independence at Layers 4 and 5.
- **`GenerateMigrationTemplate`**, **`ScaffoldKrewireYAML`**, **`GetCLICommands`**: Code-generation and CLI reference helpers.

### Transports (`http.go`, `stdio.go`)

- **`HTTPTransport`** & **`NewHTTPTransport(srv *Server)`**: `Mount(mux, "/mcp")` attaches `/mcp/sse` (Server-Sent Events with 15s keepalive pings and `event: endpoint`), `/mcp/message` (async SSE dispatch when `sessionId` is present, or synchronous JSON-RPC fallback), `/mcp/health`, and `/mcp/` metadata.
- **`StdioTransport`** & **`NewStdioTransport(srv, in, out)`**: Line-delimited JSON-RPC over `stdin`/`stdout` supporting payloads up to 4 MB.

---

## Realistic Usage Examples

### Running Locally in Stdio or HTTP Mode

```bash
cd /home/krewire/workspace/krewire/krewire.com/mcp

# Stdio mode for local IDEs (auto-detects ../content)
go run ./krewire-mcp-server --stdio

# HTTP/SSE mode on 127.0.0.1:8088
go run ./krewire-mcp-server --host 127.0.0.1 --port 8088 --docs ../content
```

---

## Verification

```bash
cd /home/krewire/workspace/krewire/krewire.com/mcp
gofmt -l .
go vet ./...
CGO_ENABLED=0 go test -v ./...
CGO_ENABLED=0 go build ./krewire-mcp-server
```

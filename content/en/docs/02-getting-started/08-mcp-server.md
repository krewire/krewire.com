---
title: "Model Context Protocol (MCP) Server"
description: "Connect AI coding agents directly to Krewire's documentation, architecture invariants, and component specifications via MCP."
order: 8
tags:
  - mcp
  - ai
  - agents
  - ide
---

# Model Context Protocol (MCP) Server

Krewire provides native, first-class support for Anthropic's **Model Context Protocol (MCP)** across two complementary tiers:

1. **Local Developer MCP Server (`krew mcp`)**: Built directly into the `krew` developer toolchain. Inspects your local project routes, runs test suites, monitors RAM/goroutine metrics, and invokes scaffolding generators over `stdio` or local HTTP/SSE.
2. **Official Remote MCP Server (`https://krewire.com/mcp/sse`)**: Hosted globally on Krewire's official portal. Provides AI assistants (Cursor, Claude Desktop, Antigravity, Windsurf) with real-time semantic search over 70+ chapters of documentation, Pop-Brutalist `.kiw` component specifications, and 6-Layer Pyramid Architecture rules.

---

## 1. Quick Connection Setup

### A. Cursor Configuration (`.cursor/mcp.json`)

To connect Cursor to the live Krewire remote documentation server, add the following to `.cursor/mcp.json` in your project root:

```json
{
  "mcpServers": {
    "krewire": {
      "url": "https://krewire.com/mcp/sse"
    }
  }
}
```

Or for offline local project inspection using `krew`:

```json
{
  "mcpServers": {
    "krewire-local": {
      "command": "krew",
      "args": ["mcp", "--stdio"]
    }
  }
}
```

### B. Claude Desktop Configuration (`claude_desktop_config.json`)

Add Krewire to your Claude Desktop configuration file:

- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Linux**: `~/.config/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "krewire": {
      "url": "https://krewire.com/mcp/sse"
    }
  }
}
```

---

## 2. Available MCP Tools

When your AI assistant connects to `https://krewire.com/mcp/sse`, it gains access to the following high-confidence tools:

- `search_docs(query, locale, limit)`: Scored full-text search across all documentation guides.
- `get_doc(path, locale)`: Fetches complete markdown text of any documentation chapter.
- `list_docs(locale)`: Returns the hierarchical table of contents.
- `list_components(category)`: Lists Pop-Brutalist `.kiw` UI components with descriptions and variants.
- `get_component(name)`: Returns property tables, variants, and copy-paste code snippets for a component.
- `get_ecosystem()`: Returns the 10-repository matrix and layer definitions (Layer 0 through Layer 5).
- `get_architecture_rules()`: Supplies the 6-Layer Pyramid rules, Three Pillars (S1-S3), and pure-Go invariants.

---

## 3. Available MCP Resources & Prompts

### Resources
- `krewire://docs/toc`: Complete table of contents of the documentation manuscript.
- `krewire://ecosystem`: Authoritative repository mapping and ownership scopes.
- `krewire://components`: Comprehensive Pop-Brutalist `.kiw` component specifications.
- `krewire://rules`: Strict architectural filters and constraints.

### Prompts
- `krewire-architect`: Guides agents on architecting applications within Krewire's 6-Layer Pyramid.
- `kiw-component-builder`: Directs agents on writing accessible, brutalist `.kiw` components.
- `krewire-migrator`: Explains how to construct Laravel-style Go database migrations with `schema.Builder`.

---
title: "Autonomous AI Agent Engine"
description: "Architecture, ReAct execution loop, safe sandboxed tools, sliding-window memory, personas, multi-protocol transport (HTTP/SSE & MCP), and embedding API in krewire/krew."
date: "2026-11-05"
---

# Autonomous AI Agent Engine (`github.com/krewire/krew`)

`krewire/krew` is the standalone, platform-agnostic autonomous AI agent monorepo of the Krewire ecosystem. It delivers a full-spectrum agent runtime featuring multi-turn ReAct reasoning, sandboxed workspace operations, sliding-window memory compaction, native Model Context Protocol (MCP) support, and a high-ergonomics embedding API.

---

## Subsystems & Architecture

### 1. `pkg/agent` — ReAct Reasoning Loop
- **Multi-Turn Reasoning:** Autonomous observe-orient-decide-act loop with bounded execution steps (`MaxSteps`).
- **Concurrent Tool Calls:** Independent LLM tool calls execute concurrently with aggregate error recovery.
- **Sandboxed Workspace Tools:** Hermetic file system and execution tools (`read_file`, `write_file`, `edit_file`, `list_dir`, `run_command`, `fetch_url`).
- **Security & Authorization:** Pluggable `Authorizer` and `Auditor` interfaces for granular tool permissions and compliance logging.

### 2. `pkg/persona` — Agent Specializations
Configurable personas tailoring system prompts, temperature, and behavioral constraints:
- `CodingAssistant`: Pair programming, structural refactoring, and test-driven implementation.
- `DevOpsRunner`: Infrastructure provisioning, CI/CD pipeline diagnosis, and container telemetry.
- `AutonomousWorker`: Unattended task completion, batch script automation, and bulk migrations.
- `GeneralAssistant`: Conversational assistant for documentation, research, and project planning.

### 3. `pkg/memory` — Token-Bounded Context Buffer
- **Sliding-Window Buffer:** `TokenBuffer` tracking message history and rolling token consumption.
- **Context Compaction:** Automated summarization when approaching LLM context window limits.
- **Persistent Sessions:** JSON-serialized session storage with resume capability (`krew chat --resume=<id>`).
- **Transcript Export:** Markdown transcript renderer for audits and pull request reviews.

### 4. `pkg/transport` — Multi-Protocol Interfaces
- **HTTP / REST Server:** Streaming execution endpoints with Server-Sent Events (`POST /api/agent/run?stream=true`).
- **Model Context Protocol (MCP):** Stdio MCP server enabling instant integration with IDEs (Cursor, Claude Desktop, Antigravity).
- **Authentication & RBAC:** Session token verification with tenant isolation and bearer token middleware.

### 5. `pkg/embed` — In-Process Go SDK
Embed the autonomous agent into any Krewire service or CLI in three lines:

```go
package main

import (
	"context"
	"fmt"

	"github.com/krewire/krew/pkg/embed"
	"github.com/krewire/krew/pkg/persona"
)

func main() {
	agent, err := embed.New(embed.Options{
		Workspace: ".",
		Persona:   &persona.CodingAssistant,
	})
	if err != nil {
		panic(err)
	}

	result, err := agent.Run(context.Background(), "Audit codebase for hardcoded literals")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Content)
}
```

---

## CLI Commands (`cmd/krew`)

The standalone `krew` CLI provides dedicated commands decomposed by Single Responsibility Principle:

- `krew chat`: Start an interactive conversational coding session in the terminal.
- `krew run <prompt>`: Execute a single instruction or automation task headless to completion.
- `krew serve`: Launch the HTTP/REST and Server-Sent Events (SSE) server daemon.
- `krew mcp`: Start the Model Context Protocol stdio server for IDE copilot workflows.
- `krew sessions`: List past conversation sessions with message and token counts.
- `krew export <id>`: Export a conversation transcript into clean Markdown.

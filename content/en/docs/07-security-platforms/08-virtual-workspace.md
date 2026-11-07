---
title: "AI-Native Virtual Workspace"
description: "Architecture, sandbox process isolation, WebSocket PTY terminal, pair programmer assistant, collaborative session hub, and daemon in krewire/dash."
date: "2026-11-05"
---

# AI-Native Virtual Workspace (`github.com/krewire/dash`)

`krewire/dash` provides an AI-powered virtual workspace engine and cloud IDE platform. It combines hermetic process isolation, bidirectional terminal streaming, Monaco-based web editing, and an embedded `krew` pair programmer into a single-binary server.

---

## Subsystems & Architecture

### 1. `pkg/sandbox` — Hermetic Isolation Engine
- **Process Jail:** Lightweight container-like sandboxing using Linux Bubblewrap (`bwrap`) with unprivileged user namespaces.
- **Resource Boundaries:** Hard memory quotas, temporary file limits, and process tree isolation per workspace.
- **Automated Idle Sweeper:** Background daemon worker monitoring workspace activity and gracefully pausing sandboxes after 30 minutes of inactivity.
- **Cross-Platform Fallback:** Safe process isolation with graceful OS-specific sysprocattr separation across Linux and macOS.

### 2. `pkg/terminal` — Bidirectional WebSocket PTY
- **Real PTY Allocation:** Kernel pseudo-terminal allocation (`pty`) spawning real shell instances (`/bin/bash` or `/bin/sh`).
- **WebSocket Streaming:** Sub-10ms latency streaming between xterm.js in the browser and the host PTY runner.
- **Window Resizing:** Dynamic terminal window column and row synchronization.

### 3. `pkg/session` — Real-Time Collaboration Hub
- **Multi-Tenant Hub:** Manages active workspace sessions and connected client connections.
- **Presence & Broadcast:** Synchronizes cursor positions, file selection changes, and live agent thoughts across collaborators.

### 4. `pkg/ai` — Embedded Agent Assistant Pool
- **Workspace-Aware Context:** Connects the embedded `krew` agent runtime directly to the active workspace sandbox filesystem.
- **Semantic Code Search:** AST-based code navigation allowing the assistant to resolve symbols and declarations across packages.
- **Interactive Diff Review:** Generates unified diffs for proposed code edits before committing them to the workspace.

### 5. `pkg/api` — Unified REST & Streaming API
- `POST /api/v1/workspaces`: Provision a new hermetic workspace.
- `GET /api/v1/workspaces`: List running and paused tenant workspaces.
- `DELETE /api/v1/workspaces/{id}`: Destroy and cleanup workspace resources.
- `GET /api/v1/workspaces/{id}/files`: Browse workspace directory tree.
- `GET /api/v1/workspaces/{id}/terminal/ws`: Connect to the interactive terminal WebSocket.
- `GET /api/v1/workspaces/{id}/session/ws`: Connect to the real-time presence hub.
- `GET /metrics`: Prometheus metrics reporting active sandboxes, CPU/RAM utilization, and terminal sessions.

---

## Workspace Access & Platform Integration

Krewire Workspace is hosted as a collaborative cloud service at [dash.krewire.com](https://dash.krewire.com).

### Cloud Workspace Portal

Developers can spin up on-demand development environments directly from their browser:
- **Web Portal:** `https://dash.krewire.com`
- **Features:** In-browser VS Code / Monaco editor, bidirectional terminal streaming, and embedded `krew` AI assistant.
- **Single Sign-On:** Integrated with Krewire Auth (`auth.krewire.com`).

### Local Development Alternatives

For local development without cloud dependencies, developers use the native `kiw` CLI tools:
- `kiw dev`: Run applications locally with auto-restart and hot reload.
- `kiw dashboard`: Run a local observability dashboard at `http://127.0.0.1:4000`.
- `kiw mcp --stdio`: Connect AI coding assistants (such as Antigravity or Cursor) directly to your project via the Model Context Protocol.

### Enterprise On-Premise Edition

Enterprise workspaces with custom VPC peering, air-gapped sandboxes, and centralized audit logging are available under enterprise contract. Contact `enterprise@krewire.com`.

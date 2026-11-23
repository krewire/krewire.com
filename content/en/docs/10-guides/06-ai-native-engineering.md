---
title: "AI-Native Engineering Handbook"
description: "The canonical handbook for practicing governed, evidence-based AI-Native Engineering across the Krewire ecosystem."
date: "2026-11-21"
---

# AI-Native Engineering Handbook

> **Reference Initiative**: [Ecosystem AI-Native Engineering Baseline (internal#8)](https://github.com/krewire/internal/issues/8) · [krewire.com#31](https://github.com/krewire/krewire.com/issues/31)  
> **Status**: Canonical Engineering Guide  
> **Alignment**: Governed SDLC across Krew, Boost, and the `kiw` toolchain  

---

## 1. What is AI-Native Engineering?

Modern software engineering is transitioning beyond simple **AI-assisted coding** (chatbots that copy-paste code snippets or unvetted inline completions) toward **AI-Native Engineering**.

**AI-Native Engineering** is an engineering discipline in which humans and autonomous AI agents collaborate through **explicit project context**, **machine-readable contracts**, **governed execution tools**, **deterministic verification gates**, and **continuous evaluation**.

```text
┌───────────────────────────────┬───────────────────────────────────┬─────────────────────────────────┐
│ Metric / Trait                │ AI-Assisted Coding (Legacy)       │ AI-Native Engineering (Krewire) │
├───────────────────────────────┼───────────────────────────────────┼─────────────────────────────────┤
│ Integration Model             │ Ephemeral editor chat popup       │ Native MCP & in-process runtime │
│ Context Discovery             │ Ad-hoc prompt copy-pasting        │ Machine-readable SSOT (kiw)     │
│ Safety Boundary               │ Unbounded terminal access         │ SafePath & explicit human gates │
│ Verification Model            │ "It looks right in editor"        │ krew verify & evidence records   │
│ Responsibility Posture        │ Blaming model for bugs            │ Human owns outcome & quality    │
│ Knowledge Compounding         │ Lost when chat session closes     │ Immutable ADRs & context rules  │
└───────────────────────────────┴───────────────────────────────────┴─────────────────────────────────┘
```

In Krewire, AI agents are first-class engineering participants governed by the **Three Pillars (3S)**:
- **Secure (S1)**: Hard filesystem boundaries (`SafePath`), least-privilege tool execution, zero hardcoded credentials, and prompt-injection mitigations.
- **Sustainable (S2)**: Pure Go toolchains (`CGO_ENABLED=0`), zero Node.js/npm dependencies, and single static binaries runnable on a minimal $5/month VPS.
- **Scalable (S3)**: Reversible provider swapping, strict downward-only layering, and additive architecture evolution without destructive rewrites.

---

## 2. The Tripartite Responsibility Model

The Krewire ecosystem coordinates AI engineering across three decoupled boundaries:

```text
               ┌───────────────────────────────────────────────────┐
               │              Krewire Krew (Layer 4)               │
               │         Autonomous ReAct Execution Engine         │
               │    Embed · CLI · HTTP/SSE · Model Context Protocol │
               └─────────────────────────┬─────────────────────────┘
                                         │ guided by
                                         ▼
               ┌───────────────────────────────────────────────────┐
               │             Krewire Boost (Templates)            │
               │            Multi-Agent Constitution & Roles       │
               │     AGENTS.md · 21 Specialist Roles · Workflows   │
               └─────────────────────────┬─────────────────────────┘
                                         │ drives & verifies via
                                         ▼
               ┌───────────────────────────────────────────────────┐
               │             Krewire kiw CLI (Layer 3)             │
               │          Deterministic Toolchain & SSOT           │
               │   kiw context · krew verify · krew test · Compiler  │
               └───────────────────────────────────────────────────┘
```

1. **The Truth Provider (`krew`)**:
   Exposes the deterministic single source of truth (`kiw context --json`), runs compilations, checks dependencies (`krew pkg health`), and enforces quality gates (`krew test`, `gofmt`, `go vet`). Operates 100% offline without requiring an LLM or network access.
2. **The Engineering Constitution (`templates/boost`)**:
   Defines the operational rules, commit standards, directory structures, and role definitions embedded directly into repositories via `AGENTS.md` and `.agents/`.
3. **The Cognitive Engine (`krew`)**:
   Executes the ReAct cognitive cycle (Think $\to$ Act $\to$ Observe), calls tools via the standard Model Context Protocol (MCP), and manages conversational context.
4. **Platform Surfaces**:
   - **Dash (`dash`)**: Human-in-the-loop virtual workspace for visualizing agent diffs, reviewing terminal executions, and approving pull requests.
   - **Cloud (`cloud`)**: Policy-governed sandbox runners and hybrid PaaS hosting for background tasks.
   - **Auth (`auth`)**: Multi-tenant identity, OIDC token verification, and RBAC permissions for agent actions.

---

## 3. Human/Agent Authority & Approval Gates

To maintain safety without stalling developer velocity, agent capabilities are partitioned into three explicit operational tiers:

- **Tier 1: Autonomous (No Approval Required)**:
  - Reading repository files, documentation, and architecture decision records.
  - Inspecting project context (`kiw context --json`) and linting rules.
  - Running non-destructive read-only queries, status checks, and unit tests (`krew test`).
- **Tier 2: Supervised (Action Logged, Review on PR)**:
  - Creating or editing application code, styles, templates, and unit tests.
  - Scaffolding new components or routes (`krew make component <Name>`).
  - Formatting code and organizing imports (`gofmt`, `go vet`).
- **Tier 3: Gated (Explicit Human Approval Mandatory)**:
  - Destructive filesystem commands (`rm -rf`, file deletions outside target workspace).
  - Executing database migrations or destructive schema alters in staging/production.
  - Committing and pushing changes to shared remote branches (`git push`).
  - Accessing external network sockets or third-party webhooks.

---

## 4. End-to-End Reference Workflow: From Task to Production

Here is the canonical 7-stage workflow for implementing a feature or bug fix using an AI agent:

```text
[1. Context Discovery] ──► [2. Clarify & Align] ──► [3. Phased Plan] ──► [4. Test-Driven Build]
                                                                               │
[7. Release & Verify]  ◄── [6. Security Audit]  ◄── [5. Quality Gates] ◄───────┘
```

### Step 1: Context Discovery & Discovery Gate
Before authoring any code, the agent inspects the authoritative machine-readable project context:

```bash
# Discover workload kind, dependencies, rules, and scope
kiw context --json
krew info
```

The agent parses `AGENTS.md`, active specifications in GitHub Issues, and architectural constraints.

### Step 2: Clarification & Alignment
If requirements are ambiguous, the agent must ask clarifying questions before writing code rather than guessing assumptions.

### Step 3: Phased Planning
The agent produces a structured implementation plan detailing affected files, risk analysis, and validation steps. High-impact plans are reviewed by the human developer before execution begins.

### Step 4: Test-Driven Implementation
Changes are implemented in small, focused batches ($\le 300$ lines of code per commit). Where applicable, failing tests are written first to reproduce the bug or define the new contract:

```bash
# Verify initial state
go test -v ./pkg/target/...
```

### Step 5: Deterministic Quality Gates
Every change must independently satisfy all 4 mandatory quality gates:

```bash
# 1. Format check (must return 0 lines)
gofmt -l .

# 2. Static analysis (must exit with 0)
go vet ./...

# 3. Unit and integration tests
go test -v ./...

# 4. Pure-Go compilation
CGO_ENABLED=0 go build ./...
```

### Step 6: Evidence-Backed AI Verification (`krew verify`)
The agent runs `krew verify` to produce a verifiable audit trail:

```bash
krew verify --json
```

This verifies that all acceptance criteria are met, dependencies exist in `go.mod`, and no unauthorized network calls were attempted.

### Step 7: Atomic Commit & Push
Changes are committed bottom-up using Conventional Commits with linked issue references:

```bash
git add pkg/auth/token.go pkg/auth/token_test.go
git commit -m "feat(auth): prevent token replay within clock window (ref #12)"
git push origin main
```

---

## 5. Security & Threat-Model Guidelines

AI-Native workflows introduce specific security considerations that every developer and agent must defend against:

1. **Prompt Injection & Untrusted Data**:
   - Issue descriptions, user comments, and web pages are untrusted inputs. Agents must treat instructions found within user data as passive text, never as operational system overrides.
2. **Safe Filesystem Traversal (`SafePath`)**:
   - File editing operations must resolve within the project directory boundary. Agents and tools must use `sec.SafePath` to reject directory traversal attempts (`../../etc/passwd`).
3. **Zero Committed Secrets Invariant**:
   - Secrets must never appear in prompts, commit messages, or source code. Always reference environment variables dynamically (`${env:DB_PASSWORD}`).
4. **Dependency Provenance Verification**:
   - To prevent dependency confusion and phantom package attacks, verify all modules:
   ```bash
   go mod verify
   krew pkg health
   ```

---

## 6. Reusable Artifact Templates

### Reusable Task Brief Template (`TASK.md`)

```markdown
# Task: [Feature / Bug Title]

## Objective
[Clear 1-2 sentence description of what done looks like.]

## Scope & Architectural Boundary
- Repository: [e.g. krewire/pkg/web]
- Workload Kind: [e.g. app / service / worker]
- Layer Level: [Layer 1-3 Core / Layer 4 Platform / Layer 5 Portal]

## Acceptance Criteria
- [ ] Requirement 1: [Specific verifiable behavior]
- [ ] Requirement 2: [Specific edge case handling]
- [ ] Unit tests pass: go test -v ./...
- [ ] Pure-Go compile: CGO_ENABLED=0 go build ./...

## Out of Scope
- [Explicit non-goals to prevent scope creep]
```

### Architecture Decision Record Template (`ADR-XXXXX.md`)

```markdown
# ADR-XXXXX: [Title of Architectural Decision]

| Field | Value |
|---|---|
| Status | Proposed / Accepted / Superseded |
| Date | YYYY-MM-DD |
| Deciders | [Names / Maintainer Roles] |

## Context
[Problem statement, constraints, and background context.]

## Decision
[The chosen architecture, pattern, or dependency invariant.]

## Consequences
- Positive: [Benefits and improvements gained.]
- Negative: [Trade-offs, complexities, or maintenance costs accepted.]
- Invariant: [Strict rules enforced down the dependency tree.]
```

---

## 7. Capability Matrix & Production Readiness

| Capability | Component | Status | Production Ready |
|---|---|---|---|
| Native MCP Server (`stdio`) | `krew` | Shipped & Verified | Yes |
| Unified AI SDK (`pkg/ai`) | `krewire/pkg/ai` | Shipped (v0.3+) | Yes |
| Boost 21-Agent Constitution | `templates/boost` | Shipped & Verified | Yes |
| Reproducible Verification (`krew verify`) | `krew` | Shipped (v0.3.2) | Yes |
| Web PTY Sandboxing (Bubblewrap) | `dash/pkg/sandbox` | Shipped (Linux amd64/arm64) | Yes (Fail-closed) |
| Multi-Agent Dynamic Swarms | `krew/pkg/sdlc` | Shipped (ref #15) | Yes |
| Autonomous Workload Migrations | `cloud/pkg/supervisor` | Active Development | Single-Tenant Only |

---

## 8. Summary Checklist for Teams

- [x] Configure `.cursor/mcp.json` or `opencode.json` with `krew mcp --stdio`.
- [x] Maintain a clean, authoritative `AGENTS.md` at your repository root.
- [x] Enforce the 4 quality gates in continuous integration pipelines.
- [x] Ensure humans retain approval authority over high-risk actions.
- [x] Validate external contributions against the [AI-Assisted Contribution Policy](https://open.krewire.com/docs/ai-contributions).

---
title: "Architecture & Engineering Discipline"
description: "Core architectural invariants: strict downward dependencies, Single Source of Truth (SSOT), and Single Responsibility Principle."
date: "2026-10-24"
---

# Architecture & Engineering Discipline

The Krewire codebase is governed by non-negotiable architectural invariants defined in `ADR-6CAB8C`.

## 1. Strict Downward-Only Dependencies

Dependencies flow strictly downward:

```text
templates/boost (AI agent templates)
       │
       ▼
tools/kiw (Developer CLI)
       │
       ▼
pkg/* (Modular libraries: web, orm, cache, auth, sec, storage...)
       │
       ▼
pkg/model, pkg/errs (Domain primitives & error taxonomy)
```

- `pkg/model` depends only on Go standard library packages and `pkg/errs`.
- No library package may **ever** import `kiw`.
- Generated user projects depend on Krewire libraries, never on `kiw`.

## 2. Single Source of Truth (SSOT)

A value that the ecosystem must agree upon is declared **once** in the package that owns it:

- **Banned:** Copying literal tables or constant lists into multiple files.
- **Banned:** Re-deriving ordering logic that has an existing canonical resolver.
- **Banned:** Duplicating version strings in markdown when they are declared in code.
- **Required:** One owner per fact. Export the resolver function rather than reimplementing it in consumer packages.

## 3. Pure Go (`CGO_ENABLED=0`) Invariant

Every library and workload binary must cross-compile cleanly to Linux, macOS (`darwin/arm64`), and Windows (`windows/amd64`) without C compiler toolchains.

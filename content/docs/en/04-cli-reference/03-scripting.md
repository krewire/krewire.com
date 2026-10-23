---
title: "Task Automation & Scripting"
description: "Declare custom workflow tasks, build pipelines, and automation scripts directly in krewire.yaml with argument forwarding and cross-platform execution."
date: "2026-10-08"
---

# Task Automation & Scripting

Krewire includes a built-in task automation engine directly within `kiw`, eliminating the need for `Makefile`, `npm run`, or external task runners.

Tasks are declared in the `scripts:` block of `krewire.yaml` and executed via `kiw run <task>`.

---

## Defining Scripts in `krewire.yaml`

```yaml
scripts:
  # Simple single-line task
  lint: "golangci-lint run ./..."

  # Sequential multi-command pipeline
  ci:
    - "kiw fmt"
    - "kiw vet"
    - "kiw test"

  # Task with environment variables
  seed:
    cmd: "go run ./cmd/seed/main.go"
    env:
      DB_MODE: "test"

  # Database migration runner
  migrate: "go run ./cmd/migrate/main.go up"
```

---

## Executing Scripts

Run any named script using `kiw run`:

```bash
# Run the lint script
kiw run lint

# Run the complete CI validation chain
kiw run ci
```

### Argument Forwarding

Arguments passed after `--` are automatically forwarded to the target command:

```bash
kiw run migrate -- --steps 2
```

The underlying command executes:
```bash
go run ./cmd/migrate/main.go up --steps 2
```

---

## Built-In Ecosystem Tasks

If no script with the specified name is found in `krewire.yaml`, `kiw run` attempts to execute matching Go files or entry points:

```bash
kiw run ./cmd/server/main.go
```

This guarantees a frictionless workflow for both scripted and ad-hoc Go development tasks.

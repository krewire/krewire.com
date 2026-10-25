---
title: "Lifecycle Supervision & Kernel"
description: "Imperative control plane of Krewire: kernel supervisor, module registry, dependency ordering, and graceful process shutdown."
date: "2026-10-25"
---

# Lifecycle Supervision & Kernel (`pkg/lifecycle`)

Import path: `github.com/krewire/krewire/pkg/lifecycle`

`lifecycle.Kernel` is the imperative control plane: it binds a validated `model.Project` to registered modules and executes a declared workload.

---

## Kernel Execution Example

```go
package main

import (
	"context"
	"os"

	"github.com/krewire/krewire/pkg/lifecycle"
	"github.com/krewire/krewire/pkg/model"
)

func main() {
	ctx := context.Background()

	k, err := lifecycle.New(project)
	if err != nil {
		panic(err)
	}

	k.Use(bootModule, webModule)          // modules initialize in dependency order
	k.RegisterHandler(model.KindApp, run) // one handler per workload kind

	if err := k.Boot(ctx); err != nil {
		panic(err)
	}
	defer k.Shutdown(ctx)

	os.Exit(k.Execute(ctx, workload).Int())
}
```

---

## Supporting Types

- **`Registry`** — `NewRegistry`, `Register`, `Ordered`, `Resolve`; ordering is resolved from each module's `DependsOn() []string`, never from arbitrary map iteration.
- **`Supervisor`** — `NewSupervisor`, `Add`, `Start`, `Stop`, `Health`, `Reload`, `WithShutdownTimeout`; supervises long-running background processes and services.
- **Capability Interfaces** — `Module`, `Startable`, and `Depender` opt-in interfaces that allow modular subsystems to hook into initialization and cleanup phases.\n
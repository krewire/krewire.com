---
title: "Kernel & Model (kern)"
description: "Foundational control plane of Krewire: workload model primitives, lifecycle supervision, error taxonomy, environment detection, and version contracts."
date: "2026-10-08"
---

# Kernel & Model (`kern`)

Import path: `github.com/krewire/krewire/packages/kern`

The **`kern`** package is the foundational control plane and substrate for the entire Krewire ecosystem. It defines domain primitives, workload kinds, lifecycle supervision, error taxonomy, environment detection, and semantic versioning contracts.

Per ADR `ADR-B7D3A`, `kern/model` may depend **only on Go standard library packages** and its internal subpackages—never on higher-level packages and never on `kiw`.

---

## Subpackage Layout

```text
packages/kern/
├── kern.go               # Top-level kernel supervisor and executor
├── model/                # Workload kinds, project scopes, and domain rules
│   ├── workload.go       # Kind definitions (app, cli, site, book, worker, service, infra, kernel)
│   └── scope.go          # Ecosystem scope levels (ARCH, BUILD, SEC, etc.)
├── lifecycle/            # Service supervisor, graceful startup & teardown hooks
├── errs/                 # Domain error classification and POSIX exit codes
├── env/                  # Environment detection (Development, Staging, Production)
├── version/              # Semantic version parsing and compatibility validation
└── log/                  # Structured logging adapters over Go's slog
```

---

## Workload Model (`model.Kind`)

The eight official workload kinds are declared as typed constants:

```go
package model

type Kind string

const (
	KindApp     Kind = "app"
	KindCLI     Kind = "cli"
	KindSite    Kind = "site"
	KindBook    Kind = "book"
	KindWorker  Kind = "worker"
	KindService Kind = "service"
	KindInfra   Kind = "infra"
	KindKernel  Kind = "kernel"
)

// IsValid reports whether the kind is one of the 8 canonical kinds.
func (k Kind) IsValid() bool { ... }
```

---

## Error Handling & Exit Codes (`kern/errs`)

Krewire standardizes errors across all packages using structured taxonomy and exit codes:

```go
import "github.com/krewire/krewire/packages/kern/errs"

// Predefined exit code constants
const (
	ExitCodeSuccess = 0
	ExitCodeFailure = 1
	ExitCodeUsage   = 2
)

// Domain error creation
err := errs.New(errs.CodeNotFound, "user resource not found")
```

---

## Environment Detection (`kern/env`)

```go
import "github.com/krewire/krewire/packages/kern/env"

current := env.Current()

if current.IsDevelopment() {
    // Enable debug diagnostics and hot reload
}

if current.IsProduction() {
    // Enforce strict TLS and secure cookies
}
```

---

## Version Contract (`kern/version`)

Semantic versions are validated at compile and runtime:

```go
import "github.com/krewire/krewire/packages/kern/version"

v := version.MustParse("0.1.0")
fmt.Println(v.Major, v.Minor, v.Patch) // 0 1 0
```

---
title: "Model, Lifecycle & Errors (pkg/model, pkg/lifecycle, pkg/errs)"
description: "Foundational control plane of Krewire: workload model primitives, lifecycle supervision, error taxonomy, environment detection, and version contracts."
date: "2026-10-25"
---

# Model, Lifecycle & Errors

Import paths:

```text
github.com/krewire/krewire/pkg/model
github.com/krewire/krewire/pkg/lifecycle
github.com/krewire/krewire/pkg/errs
github.com/krewire/krewire/pkg/env
github.com/krewire/krewire/pkg/version
```

These five packages form the foundation of the ecosystem. They define domain
primitives, workload kinds, lifecycle supervision, the error taxonomy,
environment resolution, and semantic version contracts.

Per ADR `ADR-B7D3A`, `pkg/model` may depend **only on Go standard library
packages** and its internal `errs` sibling — never on higher-level packages and
never on `kiw`.

---

## Package Layout

```text
pkg/
├── model/        # Workload kinds, project scopes, tiers, opt-in rules
├── lifecycle/    # Kernel supervisor, module registry, graceful shutdown
├── errs/         # Error classification and POSIX exit codes
├── env/          # Target environments (local, production, testing)
└── version/      # Semantic version parsing and compatibility validation
```

---

## Workload Model (`pkg/model`)

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
	KindGateway Kind = "gateway"
	KindInfra   Kind = "infra"
)

// AllKinds lists every valid Kind in canonical order.
var AllKinds = []Kind{KindApp, KindCLI, KindSite, KindBook, KindWorker, KindService, KindGateway, KindInfra}

// IsValid reports whether k is one of the eight known kinds.
func (k Kind) IsValid() bool

// ParseKind parses s as a Kind, returning a usage error on unknown input.
func ParseKind(s string) (Kind, error)
```

`AllKinds` is the single owner of the kind list: every other package (including
`kiw`'s shape detector) derives from it instead of restating the eight names.

`pkg/model` also declares `Project` (name, module path, kind, config path) with
`Validate`, the `Scope` ladder (`workspace` → `module` → `domain` → `service` →
`unit`), the open-core `Tier` ladder, and the generic `OptInRule` /
`ViolatesOptIn` constraints that keep batteries out of kinds that did not opt in.

---

## Lifecycle Supervision (`pkg/lifecycle`)

`lifecycle.Kernel` is the imperative control plane: it binds a validated
`model.Project` to registered modules and executes a declared workload.

```go
k, err := lifecycle.New(project)
if err != nil {
	return err
}

k.Use(bootModule, webModule)          // modules initialize in dependency order
k.RegisterHandler(model.KindApp, run) // one handler per workload kind

if err := k.Boot(ctx); err != nil {
	return err
}
defer k.Shutdown(ctx)

os.Exit(k.Execute(ctx, workload).Int())
```

Supporting types:

- **`Registry`** — `NewRegistry`, `Register`, `Ordered`, `Resolve`; ordering is
  resolved from each module's `DependsOn() []string`, never from map iteration.
- **`Supervisor`** — `NewSupervisor`, `Add`, `Start`, `Stop`, `Health`,
  `Reload`, `WithShutdownTimeout`; supervises long-running workloads.
- **`Module` / `Startable` / `Depender`** — opt-in capability interfaces.

---

## Error Handling & Exit Codes (`pkg/errs`)

Krewire standardizes errors across all packages with a structured taxonomy and
POSIX exit codes:

```go
import "github.com/krewire/krewire/pkg/errs"

// Exit codes
errs.ExitCodeSuccess // 0
errs.ExitCodeFailure // 1
errs.ExitCodeUsage   // 2

// Error construction
err := errs.UsageError("kiw build: unknown target")
err = errs.FailureError("build failed")
err = errs.NewError("custom", errs.ExitCode(3))

// Context that survives the error chain
err = errs.WithHint(err, "run 'kiw init' first")
err = errs.WithAttrs(err, errs.Attr{Key: "path", Value: "krewire.yaml"})
err = errs.WithStack(err)

fmt.Println(errs.FormatTree(err))
```

`errs.HintOf`, `errs.AttrsOf`, and `errs.StackOf` read that context back, so a
single top-level handler can render a complete diagnostic tree.

---

## Environment Resolution (`pkg/env`)

`pkg/env` owns the environment vocabulary; resolution itself comes from
configuration (`env:` in `krewire.yaml`, `KIW_ENV`, or `--env`).

```go
import "github.com/krewire/krewire/pkg/env"

current, err := env.ParseEnv(cfg.Env) // empty input yields env.DefaultEnv
if err != nil {
	return err
}

switch current {
case env.EnvLocal:
	// debug diagnostics and hot reload
case env.EnvProduction:
	// strict TLS and secure cookies
case env.EnvTesting:
	// deterministic fixtures
}
```

`env.Envs()` exposes the canonical order (`local`, `production`, `testing`).

---

## Version Contract (`pkg/version`)

Semantic versions are parsed, compared, and validated:

```go
import "github.com/krewire/krewire/pkg/version"

v := version.MustParseVersion("0.1.0")
fmt.Println(v.String()) // 0.1.0

if !v.IsCompatible(version.MustParseVersion("0.1.0")) {
	return fmt.Errorf("incompatible version %s", v)
}

// Ecosystem-wide compatibility gate used by `kiw compat`.
err := version.CheckEcosystemCompatibility(required, actual)

// Build watermark printed by `kiw version`.
sig := version.BuildSignature()
```

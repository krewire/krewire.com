---
title: "Workload Model & Domain Primitives"
description: "Foundational domain model of Krewire: eight workload kinds, project metadata, scope hierarchy, opt-in rules, error taxonomy, environment detection, and version contracts."
date: "2026-10-25"
---

# Workload Model & Domain Primitives (`pkg/model`, `pkg/errs`)

Import paths:

```text
github.com/krewire/krewire/pkg/model
github.com/krewire/krewire/pkg/errs
github.com/krewire/krewire/pkg/env
github.com/krewire/krewire/pkg/version
```

As a foundational core package, `pkg/model` may depend **only on Go standard library packages** and its internal `errs` sibling — never on higher-level packages and never on `krew`.

---

## Declarative Models & Boundaries (`pkg/model`)

The `pkg/model` package establishes the foundational models, architectural scopes, licensing tiers, and invariant validation rules for the Krewire ecosystem:

```go
package model

type Project struct {
	Name       string `json:"name"`
	ModulePath string `json:"modulePath"`
	ConfigPath string `json:"configPath"`
}

// Validate verifies project invariants (kebab-case name, safe module path).
func (p Project) Validate() error
```

`pkg/model` also declares the hierarchical `Scope` ladder (`Workspace` → `Module` → `Domain` → `Service` → `Unit`), the open-core licensing `Tier` ladder (`free`, `pro`, `team`, `enterprise`), immutable `DomainEvent` primitives, and the baseline security controls.

---

## Error Handling & Exit Codes (`pkg/errs`)

Krewire standardizes errors across all packages with a structured taxonomy and POSIX exit codes:

```go
import "github.com/krewire/krewire/pkg/errs"

// Exit codes
errs.ExitCodeSuccess // 0
errs.ExitCodeFailure // 1
errs.ExitCodeUsage   // 2

// Error construction
err := errs.UsageError("krew build: unknown target")
err = errs.FailureError("build failed")
err = errs.NewError("custom", errs.ExitCode(3))

// Context that survives the error chain
err = errs.WithHint(err, "run 'krew init' first")
err = errs.WithAttrs(err, errs.Attr{Key: "path", Value: "krewire.yaml"})
err = errs.WithStack(err)

fmt.Println(errs.FormatTree(err))
```

`errs.HintOf`, `errs.AttrsOf`, and `errs.StackOf` read that context back, so a single top-level handler can render a complete diagnostic tree.

---

## Environment Resolution (`pkg/env`)

`pkg/env` owns the environment vocabulary; resolution itself comes from configuration (`env:` in `krewire.yaml`, `KIW_ENV`, or `--env`).

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

// Ecosystem-wide compatibility gate used by `krew compat`.
err := version.CheckEcosystemCompatibility(required, actual)

// Build watermark printed by `krew version`.
sig := version.BuildSignature()
```\n
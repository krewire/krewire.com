---
title: "Scope Levels"
description: "The five architectural scope levels of Krewire — Workspace, Module, Domain, Service, Unit — as codified by model.Scope, and how specifications, tests, and extraction decisions use them."
date: "2026-11-09"
---

# Scope Levels

<nav class="track-crumbs" aria-label="Architecture Deep-Dive track">
<span class="track-crumbs-label">Deep-Dive track</span>
<ol>
<li><a href="/architecture">Overview</a></li>
<li><a href="/architecture/layered-pyramid">1 · Pyramid</a></li>
<li><a href="/architecture/ecosystem-dag">2 · Ecosystem DAG</a></li>
<li aria-current="step"><span>3 · Scope levels</span></li>
<li><a href="/architecture/data-flow">4 · Data flow</a></li>
</ol>
</nav>

Layers answer "who may import whom". **Scope levels** answer a different question: "how big is the thing we are talking about?" Every specification, test, and design document in the Krewire repositories names the scope it targets, so that a reviewer immediately knows whether a change touches one function or the whole workspace.

## 1. The Five Levels

```text
Workspace  ───►  Module  ───►  Domain  ───►  Service  ───►  Unit
 (level 0)      (level 1)     (level 2)     (level 3)     (level 4)
```

| Level | Scope | What it is | Example |
| :--- | :--- | :--- | :--- |
| 0 | **Workspace** | The multi-repository Go workspace rooted at a `go.work` file, with the `kiw` binary and the sibling repositories. | `krewire/` + `ui/` + `krew/` checked out side by side |
| 1 | **Module** | A Go module — the unit you version, release, and import. | `github.com/krewire/krewire`, `github.com/krewire/ui` |
| 2 | **Domain** | A bounded context inside a module: a cohesive set of services, usually `internal/<domain>/`. Before extraction it is a group of packages; after extraction it can become its own module. | `internal/catalog`, `internal/user` |
| 3 | **Service** | A runtime deployable implemented as a Go `main` package. | `cmd/api/`, `service/billing/` |
| 4 | **Unit** | The smallest piece of code: a package, file, type, or function. | `pkg/ui`, `ui.Button` |

Older names are no longer accepted: use **Module** instead of "Project", and **Unit** instead of "Package" or "Func".

## 2. Scope as Code (`pkg/model`)

The hierarchy is not just documentation — it is a type in the hermetic core package `pkg/model`:

```go
package main

import (
	"fmt"

	"github.com/krewire/krewire/pkg/model"
)

func main() {
	for _, s := range model.AllScopes {
		fmt.Println(s, s.Level()) // Workspace 0 … Unit 4
	}

	s, err := model.ParseScope(" domain ") // case-insensitive, trims spaces
	fmt.Println(s, err)                    // Domain <nil>

	fmt.Println(model.ScopeService.Less(model.ScopeUnit)) // true: Service is the wider scope

	_, err = model.ParseScope("package")
	fmt.Println(err) // unknown scope "package": want one of Workspace, Module, Domain, Service, Unit
}
```

| API | Behaviour |
| :--- | :--- |
| `model.AllScopes` | All five scopes in canonical order. |
| `Scope.IsValid()` | Reports whether the value is one of the five scopes. |
| `Scope.Level()` | `0` (Workspace) to `4` (Unit); `-1` for an invalid scope. |
| `Scope.Less(other)` | Whether the scope is ordered before (wider than) `other`. |
| `model.ParseScope(s)` | Parses a name case-insensitively; unknown names return a usage error. |

## 3. Using Scope Levels in Practice

- **Specs and issues** state their scope, so "add a `Button` variant" (Unit) is never reviewed like "split the catalog domain into its own module" (Domain → Module).
- **Tests** are organised by scope: unit tests next to the code, service tests against a running `main` package, and workspace-level checks such as `layering_test.go` across modules.
- **Extraction decisions** move code *up* one level at a time: a Unit grows into a Domain inside a monolith; a Domain becomes a Service (its own `main`) when it needs to scale or deploy independently; a Service becomes its own Module when another team must version it separately. Each step keeps the code; only its packaging changes — the [progressive pipeline](/overview/why-krewire) from a single `app` to `service` and `worker` workloads.

---

**Previous:** [11.2 Ecosystem DAG](/architecture/ecosystem-dag) · **Next:** [**11.4 Data Flow: SSG vs Runtime →**](/architecture/data-flow)

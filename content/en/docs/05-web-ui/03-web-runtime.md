---
title: "Web Runtime & WASM (pkg/runtime)"
description: "Author client-side reactive web applications in Go compiled to WebAssembly: virtual DOM diffing, hooks, component registry, SSR hydration markers, and multi-framework island hydration."
date: "2026-10-25"
---

# Web Runtime & WASM (`pkg/runtime`)

Import path: `github.com/krewire/krewire/pkg/runtime`

The `pkg/runtime` library delivers a client-side WebAssembly (WASM) execution engine, Virtual DOM (VDOM) diffing and reconciliation algorithm, Fiber component lifecycle scheduler, and progressive multi-framework island hydration contracts.

It allows developers to write interactive, reactive client-side widgets in 100% pure Go compiled to WebAssembly (`GOOS=js GOARCH=wasm`), or seamlessly embed lightweight modular frontend framework islands (Svelte 5, Vue 3) while keeping initial page loads static and zero-JS by default.

---

## 1. Subpackage Layout

```text
pkg/runtime/
├── vdom/         # Immutable virtual DOM: VNode, El, Text, Component, Diff, RenderHTML
├── component/    # Component model, Fiber scheduler, state hooks, event registry
├── mount/        # Progressive hydration contract (data-kiw-mount, load, idle, visible)
├── boot/         # Scans data-kiw-mount markers and hydrates registered components
├── schedule/     # Per-frame render batching (requestAnimationFrame on js/wasm)
└── build/        # BuildWASM: pure-Go wasm compiler wrapper with content hashing
```

---

## 2. Progressive Island Hydration (`pkg/runtime/mount`)

Krewire implements the **Islands Architecture**: pages are rendered statically on the server (SSG/SSR) with zero JavaScript bloat. Interactive islands are hydrated selectively on the client based on user visibility or browser idle cycles.

### Hydration HTML Contract

Mount points emit standard HTML markers:

```html
<div data-kiw-mount="Counter"
     data-kiw-framework="wasm"
     data-kiw-hydrate="visible"
     data-kiw-props='{"initial":10}'>
  <!-- Server-rendered fallback HTML -->
  <button class="btn">Count: 10</button>
</div>
```

### Hydration Strategies (`mount.Hydrate`)

- **`load`** (`mount.Load`): Hydrates immediately as soon as the client module or WASM bundle boots.
- **`idle`** (`mount.Idle`): Hydrates during browser idle periods via `window.requestIdleCallback`.
- **`visible`** (`mount.Visible`): Hydrates only when the island scrolls into the browser viewport via `IntersectionObserver`.

### Multi-Framework Island Support

Through `pkg/runtime/mount` and the Krewire plugin ecosystem, mount points support:
- **`wasm`** (Default): Pure-Go compiled WebAssembly island with zero node/npm dependencies.
- **`svelte`**: Svelte 5 reactive component island installed via `krew add svelte`.
- **`vue`**: Vue 3 reactive component island installed via `krew add vue`.

---

## 3. Virtual DOM Engine (`pkg/runtime/vdom`)

```go
import "github.com/krewire/krewire/pkg/runtime/vdom"

tree := vdom.El("section", map[string]string{"class": "hero"},
	vdom.El("h1", nil, vdom.Text("Hello from Go")),
	vdom.Component("Counter", "counter-1", map[string]string{"initial": "0"}),
)

html := vdom.RenderHTML(tree)        // Server-side rendering (SSR)
patches := vdom.Diff(oldTree, tree) // Client-side reconciliation
```

`VNode` is an immutable description. Both `RenderHTML` and `Diff` funnel props through `NormalizeProps`, ensuring server output and client reconciliation agree on attribute names, classes, and event markers.

---

## 4. Component Model & Reactive Hooks (`pkg/runtime/component`)

A component is anything implementing the `Component` interface:

```go
type Component interface {
	Render() *vdom.VNode
}
```

Stateful logic uses generic hooks that bind to the active `Fiber` and enforce stable execution order:

```go
package main

import (
	"context"
	"fmt"

	"github.com/krewire/krewire/pkg/runtime/component"
	"github.com/krewire/krewire/pkg/runtime/vdom"
)

type Counter struct{}

func (c *Counter) Render() *vdom.VNode {
	count, setCount := component.UseState(0)

	component.UseEvent("click", "inc", func(_ context.Context, _ component.Event) {
		setCount(count + 1)
	})

	return vdom.El("div", map[string]string{"class": "counter"},
		vdom.El("h2", nil, vdom.Text(fmt.Sprintf("Count: %d", count))),
		vdom.El("button", map[string]string{"data-kiw-on": "click:inc"}, vdom.Text("Increment")),
	)
}

func init() {
	component.MustRegister("Counter", func(component.Props) component.Component {
		return &Counter{}
	})
}
```

Available hooks include `UseState`, `UseRef`, `UseMemo`, `UseEffect`, and `UseEvent`. Calling a hook outside a render cycle panics with `*component.HookRuleError`, immediately surfacing incorrect hook ordering.

---

## 5. Client Boot Scanner (`pkg/runtime/boot`)

The client runtime automatically boots and mounts components upon DOM ready:

```go
package main

import (
	"github.com/krewire/krewire/pkg/runtime/boot"
	"github.com/krewire/krewire/pkg/runtime/component"
)

func main() {
	// 1. Register interactive Go components
	component.Register("Counter", func() component.Component {
		return &Counter{}
	})

	// 2. Scan DOM for [data-kiw-mount] markers and mount islands
	boot.Run()
}
```

---

## 6. Building WASM Binaries (`pkg/runtime/build`)

```go
import "github.com/krewire/krewire/pkg/runtime/build"

artifacts, err := build.BuildWASM(build.Config{
	Entry:  "./frontend",
	OutDir: "public/wasm",
	Name:   "counter",
})
if err != nil {
	return err
}

fmt.Println(artifacts.WASM, artifacts.JS, artifacts.Digest)
```

From the developer CLI, the same pipeline is accessible as a build target:

```bash
krew build --target wasm
```

---

## 7. Alignment with the Three Pillars (3S)

- **Secure (S1)**: WASM executes inside the browser's hermetic WebAssembly sandbox with no direct unsafe memory access. VDOM auto-escapes all text nodes and attribute values to eliminate Cross-Site Scripting (XSS).
- **Sustainable (S2)**: Zero JavaScript runtime dependencies for pure Go WASM workloads. Progressive island hydration guarantees users download interactive scripts only for sections requiring dynamic behavior.
- **Scalable (S3)**: Unifies Go WASM, Svelte 5, and Vue 3 under a single declarative island hydration contract (`data-kiw-hydrate`).

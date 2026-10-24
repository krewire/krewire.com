---
title: "Web Runtime & WASM (pkg/runtime)"
description: "Author client-side reactive web applications in Go compiled to WebAssembly: virtual DOM diffing, hooks, component registry, SSR hydration markers, and a reproducible WASM build pipeline."
date: "2026-10-25"
---

# Web Runtime & WASM (`pkg/runtime`)

The `pkg/runtime` family compiles Go to WebAssembly (`GOOS=js GOARCH=wasm`) so
interactive islands can share the same component model as server-rendered
`.kiw` pages. It is a **library**, not a workload kind: a WASM frontend ships
inside an `app` or `site` project and is built with `kiw build --target wasm`.

---

## Subpackage Layout

```text
pkg/runtime/
├── vdom/         # Immutable virtual DOM: VNode, El, Text, Component, Diff, RenderHTML
├── component/    # Component model, hooks, event handlers, component registry
├── mount/        # SSR hydration contract shared by the emitter and the client boot
├── boot/         # Scans data-kiw-mount markers and hydrates registered components
├── schedule/     # Per-frame render batching (requestAnimationFrame on js/wasm)
└── build/        # BuildWASM: go build wrapper with hashing and diagnostics
```

---

## Virtual DOM (`pkg/runtime/vdom`)

```go
import "github.com/krewire/krewire/pkg/runtime/vdom"

tree := vdom.El("section", map[string]string{"class": "hero"},
	vdom.El("h1", nil, vdom.Text("Hello from Go")),
	vdom.Component("Counter", "counter-1", map[string]string{"initial": "0"}),
)

html := vdom.RenderHTML(tree) // server-side rendering
patches := vdom.Diff(oldTree, tree) // client-side reconciliation
```

`VNode` is an immutable description. `RenderHTML` and `Diff` both funnel props
through `NormalizeProps`, so server output and client reconciliation agree on
attribute names, classes, and event markers.

---

## Component Model (`pkg/runtime/component`)

A component is anything that renders a virtual tree:

```go
type Component interface {
	Render() *vdom.VNode
}
```

Stateful logic uses generic hooks, which bind to the `Fiber` currently being
rendered and enforce a stable call order:

```go
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

Available hooks: `UseState`, `UseRef`, `UseMemo`, `UseEffect`, `UseEvent`.
Lifecycle participation is opt-in through the `Mounter`, `Unmounter`, and
`ShouldUpdater` capability interfaces. Calling a hook outside a render panics
with `*component.HookRuleError`, which surfaces the ordering mistake instead of
silently corrupting state.

---

## SSR Hydration (`pkg/runtime/mount` and `pkg/runtime/boot`)

The server emits a fragment wrapped in a container carrying three attributes:

```html
<div data-kiw-mount="Counter" data-kiw-hydrate="load" data-kiw-props='{"initial":0}'>
  <!-- server-rendered HTML -->
</div>
```

`boot` scans for those markers on the client, instantiates the registered
component, and seeds the fiber with the server-rendered DOM so the first render
diffs instead of re-rendering. Later updates coalesce through `schedule`, which
batches notifications into one flush per animation frame.

---

## Building a WASM Module (`pkg/runtime/build`)

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

`BuildWASM` wraps `go build` with `GOOS=js GOARCH=wasm`, copies `wasm_exec.js`,
content-hashes the module, and returns actionable diagnostics for toolchain and
version mismatches.

From the CLI, the same pipeline is reachable as a build target:

```bash
kiw build --target wasm
```

The resulting module and JS glue are placed in your output directory, ready to be
mounted into any HTML page served by the `app` or `site` workload.

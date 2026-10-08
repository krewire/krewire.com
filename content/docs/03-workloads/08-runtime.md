---
title: "WebAssembly Runtime (runtime)"
description: "Author client-side reactive web applications in Go compiled directly to WebAssembly with virtual DOM diffing, component mounting, and partial hydration."
date: "2026-10-08"
---

# WebAssembly Runtime (`runtime`)

The **`runtime`** workload compiles Go code directly to WebAssembly (`GOOS=js GOARCH=wasm`), enabling rich client-side interactivity, reactive state management, and virtual DOM diffing without writing JavaScript.

Powered by `packages/runtime`, the WebAssembly engine integrates cleanly with `.kiw` server templates through **island hydration** patterns.

---

## Architectural Characteristics

1. **Virtual DOM Tree:** High-speed in-memory VDOM diffing that updates only mutated DOM elements.
2. **Selective Island Hydration:** Serve complete static HTML from SSR/SSG, then hydrate only interactive widgets (`data-kiw-island`) on the client.
3. **Standard Go Toolchain:** Compiles with standard Go compiler flags (`GOOS=js GOARCH=wasm go build`)—zero custom compiler forks or external bundlers required.
4. **Zero JS Glue Maintenance:** `kiw` automatically bundles `wasm_exec.js` and provides an ergonomic client mounting bridge.

---

## Scaffolding a `runtime` Project

```bash
kiw new interactive-widget --runtime
cd interactive-widget
```

Directory layout:

```text
interactive-widget/
├── krewire.yaml          # Project configuration
├── go.mod
└── main.go               # Client WASM application
```

---

## Interactive Widget Implementation (`main.go`)

```go
//go:build js && wasm

package main

import (
	"fmt"

	"github.com/krewire/krewire/packages/runtime"
)

type CounterState struct {
	Count int
}

func main() {
	app := runtime.NewApp()

	state := &CounterState{Count: 0}

	app.Mount("#counter-widget", func() runtime.VNode {
		return runtime.Div(
			runtime.Class("counter-box"),
			runtime.H2(runtime.Text(fmt.Sprintf("Current Count: %d", state.Count))),
			runtime.Button(
				runtime.Class("btn btn-primary"),
				runtime.OnClick(func(e runtime.Event) {
					state.Count++
					app.Render() // Re-render VDOM tree
				}),
				runtime.Text("Increment +"),
			),
			runtime.Button(
				runtime.Class("btn btn-outline"),
				runtime.OnClick(func(e runtime.Event) {
					if state.Count > 0 {
						state.Count--
						app.Render()
					}
				}),
				runtime.Text("Decrement -"),
			),
		)
	})

	// Keep runtime alive in browser thread
	select {}
}
```

---

## Configuration (`krewire.yaml`)

```yaml
project:
  name: "interactive-widget"
  kind: runtime
  version: "0.1.0"

build:
  output: "public/app.wasm"
  target: wasm
```

---

## Compiling and Running

Compile the WebAssembly binary:

```bash
kiw build --target wasm
```

The resulting `app.wasm` is placed into your output directory, ready to be mounted into any HTML page via `<script src="/assets/wasm_exec.js"></script>`.

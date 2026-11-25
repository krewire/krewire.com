---
title: "Terminal UI Toolkit"
description: "Interactive terminal user interfaces, Elm/Bubble Tea architecture, ANSI color styling, tables, and progress bars in tui."
date: "2026-10-25"
---

# Terminal UI Toolkit (`tui`)

Import path: `github.com/krewire/tui`

The `tui` package powers interactive terminal user interfaces across the `krew` CLI and developer tools using an Elm/Bubble Tea architecture.

---

## Usage Example

```go
package main

import (
	"github.com/krewire/tui/term"
	"github.com/krewire/tui"
)

type appModel struct {
	selected int
}

func (m appModel) Init() tui.Cmd { return nil }

func (m appModel) Update(msg tui.Msg) (tui.Model, tui.Cmd) {
	if key, ok := msg.(tui.KeyMsg); ok && key.String() == "q" {
		return m, tui.Quit
	}
	return m, nil
}

func (m appModel) View() string {
	return term.Bold(term.Green("Krewire Terminal Ready! Press 'q' to quit.
"))
}

func main() {
	p := tui.NewProgram(appModel{})
	p.Start()
}
```\n
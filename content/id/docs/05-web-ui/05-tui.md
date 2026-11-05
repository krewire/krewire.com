---
title: "Toolkit Terminal UI"
description: "Antarmuka pengguna terminal interaktif, arsitektur Elm/Bubble Tea, pewarnaan ANSI, tabel, dan bilah progres di pkg/tui."
date: "2026-10-25"
---

# Toolkit Terminal UI (`pkg/tui`)

Jalur impor: `github.com/krewire/krewire/pkg/tui`

Paket `tui` mentenagai antarmuka terminal interaktif untuk CLI `kiw` dan berbagai devtool menggunakan arsitektur Elm/Bubble Tea.

---

## Contoh Penggunaan

```go
package main

import (
	"github.com/krewire/krewire/pkg/term"
	"github.com/krewire/krewire/pkg/tui"
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
	return term.Bold(term.Green("Krewire Terminal Siap! Tekan 'q' untuk keluar.
"))
}

func main() {
	p := tui.NewProgram(appModel{})
	p.Start()
}
```\n
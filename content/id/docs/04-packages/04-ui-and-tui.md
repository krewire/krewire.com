---
title: "Antarmuka Pengguna (ui, tui)"
description: "Rekayasa antarmuka pengguna lintas permukaan: Single-File Component web, token desain, tema terang/gelap, dan antarmuka terminal interaktif."
date: "2026-10-08"
---

# Antarmuka Pengguna (`ui`, `tui`)

Jalur impor (*import paths*):
- `github.com/krewire/krewire/pkg/ui`
- `github.com/krewire/krewire/pkg/tui`

Krewire menyediakan tooling antarmuka terpadu yang mencakup peramban web (`pkg/ui`) maupun konsol terminal (`pkg/tui`).

---

## 1. Komponen Web UI & Token Desain (`pkg/ui`)

Paket `ui` menyediakan Single-File Component pakai ulang (`.kiw`), token desain CSS semantik, dan engine peralihan tema terang/gelap.

### Token Tema Kanonikal

Tema sepenuhnya digerakkan melalui variabel custom CSS yang terpasang pada `:root` dan `html[data-theme="dark"]`:

```css
:root {
  --primary: #39D353;
  --primary-content: #0B1F3B;
  --secondary: #00D1C1;
  --accent: #FF3B2E;
  --bg: #FAF8F4;
  --fg: #0B1F3B;
  --muted: #475569;
  --radius: 14px;
}

html[data-theme="dark"] {
  --primary: #39D353;
  --bg: #0B1F3B;
  --fg: #FAF8F4;
  --muted: #8fa2bc;
}
```

### Komponen UI Siap Pakai
Tersedia di dalam direktori `pkg/ui/components/`:
- `<Navbar />`: Bar navigasi responsif lengkap dengan pemilih bahasa, tema, dan toggle drawer ponsel.
- `<Sidebar />`: Navigasi sidebar yang dapat diciutkan (*collapsible*) untuk dokumentasi.
- `<Button />`: Tombol aksesibel dengan varian `primary`, `secondary`, dan `outline`.
- `<Terminal />`: Pratinjau blok kode terminal dengan styling sintaks.
- `<Drawer />`: Hamparan drawer geser untuk tampilan layar ponsel.

---

## 2. Toolkit Terminal UI (`pkg/tui`)

Paket `tui` mentenagai antarmuka pengguna terminal interaktif menggunakan pola arsitektur Elm/Bubble Tea:

```go
package main

import (
	"fmt"
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
	return term.Bold(term.Green("Terminal Krewire Siap! Tekan 'q' untuk keluar.\n"))
}

func main() {
	p := tui.NewProgram(appModel{})
	p.Start()
}
```

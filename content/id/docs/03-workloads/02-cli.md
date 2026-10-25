---
title: "Terminal CLI & TUI (cli)"
description: "Bangun antarmuka command-line pengembang dan aplikasi terminal interaktif di Go dengan pkg/tui, parsing flag POSIX, dan disiplin exit code yang ketat."
date: "2026-10-08"
---

# Terminal CLI & TUI (`cli`)

Workload **`cli`** dirancang khusus untuk membangun utilitas command-line berkinerja tinggi, perkakas sistem, CLI pengembang, dan terminal user interface (TUI) yang interaktif.

Ditenagai oleh `pkg/tui` dan `pkg/term`, workload `cli` menghadirkan pola **Model-Update-View (MUV)** yang terinspirasi dari Elm/Bubble Tea dengan nol dependensi eksternal serta disiplin exit code POSIX yang ketat (`pkg/errs`).

---

## Karakteristik Arsitektur

1. **Ringan & Cepat:** Memulai eksekusi dalam waktu kurang dari 2 md tanpa dependensi terminal pihak ketiga atau CGo yang berat.
2. **Arsitektur Model-Update-View:** Aplikasi terminal reaktif berbasis state dengan double-buffering otomatis dan manajemen kursor.
3. **Degradasi Anggun (Graceful Degradation):** Mendeteksi pipe non-TTY, lingkungan CI (`CI=true`), dan pengaturan `NO_COLOR`, secara otomatis menonaktifkan kode escape ANSI.
4. **Exit Code POSIX:** Mematuhi konvensi standar Unix: `0` untuk sukses, `1` untuk kesalahan umum, dan `2` untuk kesalahan sintaks command-line.

---

## Scaffolding Proyek `cli`

```bash
kiw new sysmon --cli
cd sysmon
```

Struktur direktori standar proyek:

```text
sysmon/
├── krewire.yaml          # Konfigurasi proyek
├── go.mod                # Definisi modul Go
├── cmd/
│   └── sysmon/
│       └── main.go       # Titik masuk CLI
└── internal/
    ├── commands/         # Handler sub-perintah
    ├── ui/               # Tampilan dan model TUI
    └── system/           # Pengumpul metrik inti
```

---

## Implementasi TUI Interaktif

Berikut adalah contoh dasbor terminal interaktif menggunakan `pkg/tui`:

```go
package main

import (
	"fmt"
	"os"

	"github.com/krewire/krewire/pkg/term"
	"github.com/krewire/krewire/pkg/tui"
)

type model struct {
	cursor int
	items  []string
	status string
}

func initialModel() model {
	return model{
		items:  []string{"Database Service", "Auth Gateway", "Worker Queue", "Redis Cache"},
		status: "Operational",
	}
}

func (m model) Init() tui.Cmd {
	return nil
}

func (m model) Update(msg tui.Msg) (tui.Model, tui.Cmd) {
	switch msg := msg.(type) {
	case tui.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tui.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := term.Bold("=== System Health Monitor ===\n\n")

	for i, item := range m.items {
		cursor := " "
		if m.cursor == i {
			cursor = term.Green("▸")
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, term.Cyan("OK"), item)
	}

	s += fmt.Sprintf("\nStatus: %s\n", term.Green(m.status))
	s += term.Dim("\nPress ↑/↓ to navigate, 'q' to exit.\n")
	return s
}

func main() {
	p := tui.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
```

---

## Konfigurasi (`krewire.yaml`)

```yaml
project:
  name: "sysmon"
  kind: cli
  version: "0.1.0"

build:
  output: ".krewire/build"      # binary ditempatkan di .krewire/build/bin/sysmon
```

---

## Menjalankan dan Menguji

Jalankan CLI selama tahap pengembangan:

```bash
kiw run -- --verbose
```

Build binary rilis:

```bash
kiw build
./bin/sysmon

---
title: "Terminal CLI & TUI (cli)"
description: "Build developer command-line interfaces and interactive terminal applications in Go with packages/tui, POSIX flag parsing, and strict exit code discipline."
date: "2026-10-08"
---

# Terminal CLI & TUI (`cli`)

The **`cli`** workload is engineered for authoring high-performance command-line utilities, system tools, developer CLIs, and interactive terminal user interfaces (TUIs).

Powered by `packages/tui` and `packages/term`, the `cli` workload delivers an Elm/Bubble Tea-inspired **Model-Update-View (MUV)** pattern with zero external dependencies and strict POSIX exit code discipline (`packages/kern/errs`).

---

## Architectural Characteristics

1. **Lightweight & Fast:** Starts in under 2ms with zero heavy CGo or third-party terminal dependencies.
2. **Model-Update-View Architecture:** Stateful, reactive terminal applications with automatic double-buffering and cursor management.
3. **Graceful Degradation:** Detects non-TTY pipes, CI environments (`CI=true`), and `NO_COLOR` settings, automatically disabling ANSI escape sequences.
4. **POSIX Exit Codes:** Conforms to standard Unix conventions: `0` for success, `1` for general error, and `2` for command-line syntax errors.

---

## Scaffolding a `cli` Project

```bash
kiw new sysmon --cli
cd sysmon
```

Standard project directory structure:

```text
sysmon/
├── krewire.yaml          # Project configuration
├── go.mod                # Go module definition
├── cmd/
│   └── sysmon/
│       └── main.go       # CLI entry point
└── internal/
    ├── commands/         # Subcommand handlers
    ├── ui/               # TUI views and models
    └── system/           # Core metrics collectors
```

---

## Interactive TUI Implementation

Here is an interactive terminal dashboard using `packages/tui`:

```go
package main

import (
	"fmt"
	"os"

	"github.com/krewire/krewire/packages/term"
	"github.com/krewire/krewire/packages/tui"
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

## Configuration (`krewire.yaml`)

```yaml
project:
  name: "sysmon"
  kind: cli
  version: "0.1.0"

build:
  output: "bin/sysmon"
```

---

## Running and Testing

Run the CLI during development:

```bash
kiw run -- --verbose
```

Build the release binary:

```bash
kiw build
./bin/sysmon
```

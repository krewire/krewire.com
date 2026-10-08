---
title: "User Interfaces (ui, tui)"
description: "Cross-surface user interface engineering: web Single-File Components, design tokens, light/dark themes, and interactive terminal interfaces."
date: "2026-10-08"
---

# User Interfaces (`ui`, `tui`)

Import paths:
- `github.com/krewire/krewire/packages/ui`
- `github.com/krewire/krewire/packages/tui`

Krewire provides unified interface tooling spanning both web browsers (`packages/ui`) and terminal consoles (`packages/tui`).

---

## 1. Web UI Components & Tokens (`packages/ui`)

The `ui` package provides reusable Single-File Components (`.kiw`), semantic CSS design tokens, and a theme switching engine.

### Canonical Theme Tokens

Themes are driven entirely via CSS custom properties attached to `:root` and `html[data-theme="dark"]`:

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

### Shared UI Components
Available in `packages/ui/components/`:
- `<Navbar />`: Responsive navigation bar with theme and mobile drawer toggles.
- `<Sidebar />`: Collapsible sidebar navigation for documentation.
- `<Button />`: Accessible button with `primary`, `secondary`, and `outline` variants.
- `<Terminal />`: Terminal code block preview with syntax styling.
- `<Drawer />`: Slide-out mobile drawer overlay.

---

## 2. Terminal UI Toolkit (`packages/tui`)

The `tui` package powers interactive terminal user interfaces using an Elm/Bubble Tea architecture:

```go
package main

import (
	"fmt"
	"github.com/krewire/krewire/packages/term"
	"github.com/krewire/krewire/packages/tui"
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
	return term.Bold(term.Green("Krewire Terminal Ready! Press 'q' to quit.\n"))
}

func main() {
	p := tui.NewProgram(appModel{})
	p.Start()
}
```

---
title: "Krewire UI & Design System"
description: "Zero-dependency pop-brutalist design system and component architecture for the Krewire ecosystem."
date: "2026-11-09"
---

# Krewire UI (`github.com/krewire/ui`)

Import path: `github.com/krewire/ui`  
Components Directory: `/docs/ui/components`  
Design System Stylesheet: `/styles/krewire-ui.css`

**Krewire UI (`kui`)** is a zero-dependency, neo-brutalist ("pop-brutalist") design system and Single-File Component library built natively for the Krewire Go ecosystem.

All components follow the **shadcn/ui "Copy & Own"** model: you copy `.kiw` component files directly into your project's `./components/` directory, giving you complete ownership without opaque runtime dependencies or upstream breaking changes.

---

## What is in this Chapter?

1. [**Components Directory (`/ui/components`)**](/docs/ui/components) — Complete reference catalog for all 66 Single-File Components, including buttons, cards, modals, navigation bars, forms, data displays, and layout primitives.
2. **Design Tokens & Pop-Brutalism** — High-contrast semantic color palette, 2.5px solid borders, hard offset drop shadows (`4px 4px 0px var(--border)`), and accessible dark mode inversion.
3. **Zero-Dependency Architecture** — Pure Go compilation (`CGO_ENABLED=0`), zero npm or Node.js runtime bloat, and standard library `embed.FS` distribution.

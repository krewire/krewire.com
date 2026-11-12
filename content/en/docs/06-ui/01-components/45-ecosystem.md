---
title: "Ecosystem"
description: "Interactive grid component displaying Krewire's 10-repository architecture."
date: "2026-11-09"
---

# Ecosystem

Category: **Data Display & Content**

Interactive grid component displaying Krewire's 10-repository architecture.

---

## Basic Usage

Use `<Ecosystem />` directly within your `.kiw` page or layout templates:

```html
<Ecosystem class="my-custom-ecosystem">
  <!-- Slot content or properties -->
</Ecosystem>
```

---

## Supported Props

- `badge`: Configures the `badge` attribute or behavior.
- `header`: Configures the `header` attribute or behavior.
- `lead`: Configures the `lead` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

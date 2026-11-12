---
title: "Grid"
description: "CSS Grid container with configurable columns and gap tokens."
date: "2026-11-09"
---

# Grid

Category: **Layout Primitives**

CSS Grid container with configurable columns and gap tokens.

---

## Basic Usage

Use `<Grid />` directly within your `.kiw` page or layout templates:

```html
<Grid class="my-custom-grid">
  <!-- Slot content or properties -->
</Grid>
```

---

## Supported Props

- `cols`: Configures the `cols` attribute or behavior.
- `gap`: Configures the `gap` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

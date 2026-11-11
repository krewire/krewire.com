---
title: "ButtonGroup"
description: "Grouped button segment with unified borders and active states."
date: "2026-11-09"
---

# ButtonGroup

Category: **Actions & Buttons**

Grouped button segment with unified borders and active states.

---

## Basic Usage

Use `<ButtonGroup />` directly within your `.kiw` page or layout templates:

```html
<ButtonGroup class="my-custom-button-group">
  <!-- Slot content or properties -->
</ButtonGroup>
```

---

## Supported Props

- `ariaLabel`: Configures the `ariaLabel` attribute or behavior.
- `item`: Configures the `item` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

---
title: "Tooltip"
description: "Accessible hover and focus informative tooltip."
date: "2026-11-09"
---

# Tooltip

Category: **Overlays & Dialogs**

Accessible hover and focus informative tooltip.

---

## Basic Usage

Use `<Tooltip />` directly within your `.kiw` page or layout templates:

```html
<Tooltip class="my-custom-tooltip">
  <!-- Slot content or properties -->
</Tooltip>
```

---

## Supported Props

- `id`: Configures the `id` attribute or behavior.
- `placement`: Configures the `placement` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

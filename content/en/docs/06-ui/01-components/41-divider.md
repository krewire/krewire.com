---
title: "Divider"
description: "Horizontal or vertical separator line with optional text label."
date: "2026-11-09"
---

# Divider

Category: **Data Display & Content**

Horizontal or vertical separator line with optional text label.

---

## Basic Usage

Use `<Divider />` directly within your `.kiw` page or layout templates:

```html
<Divider class="my-custom-divider">
  <!-- Slot content or properties -->
</Divider>
```

---

## Supported Props

- `label`: Configures the `label` attribute or behavior.
- `vertical`: Configures the `vertical` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

---
title: "Box"
description: "Versatile layout box supporting border, padding, and pop-brutalist shadows."
date: "2026-11-09"
---

# Box

Category: **Layout Primitives**

Versatile layout box supporting border, padding, and pop-brutalist shadows.

---

## Basic Usage

Use `<Box />` directly within your `.kiw` page or layout templates:

```html
<Box class="my-custom-box">
  <!-- Slot content or properties -->
</Box>
```

---

## Supported Props

- `border`: Configures the `border` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `rounded`: Configures the `rounded` attribute or behavior.
- `shadow`: Configures the `shadow` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

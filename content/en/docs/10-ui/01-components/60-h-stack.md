---
title: "HStack"
description: "Horizontal flex container with uniform spacing between children."
date: "2026-11-09"
---

# HStack

Category: **Layout Primitives**

Horizontal flex container with uniform spacing between children.

---

## Basic Usage

Use `<HStack />` directly within your `.kiw` page or layout templates:

```html
<HStack class="my-custom-h-stack">
  <!-- Slot content or properties -->
</HStack>
```

---

## Supported Props

- `align`: Configures the `align` attribute or behavior.
- `gap`: Configures the `gap` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `justify`: Configures the `justify` attribute or behavior.
- `noWrap`: Configures the `noWrap` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

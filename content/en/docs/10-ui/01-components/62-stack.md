---
title: "Stack"
description: "Directional layout container with configurable gap spacing."
date: "2026-11-09"
---

# Stack

Category: **Layout Primitives**

Directional layout container with configurable gap spacing.

---

## Basic Usage

Use `<Stack />` directly within your `.kiw` page or layout templates:

```html
<Stack class="my-custom-stack">
  <!-- Slot content or properties -->
</Stack>
```

---

## Supported Props

- `align`: Configures the `align` attribute or behavior.
- `direction`: Configures the `direction` attribute or behavior.
- `gap`: Configures the `gap` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `justify`: Configures the `justify` attribute or behavior.
- `wrap`: Configures the `wrap` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

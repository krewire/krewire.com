---
title: "Progress"
description: "Determinate and indeterminate progress bar indicator."
date: "2026-11-09"
---

# Progress

Category: **Feedback & Status**

Determinate and indeterminate progress bar indicator.

---

## Basic Usage

Use `<Progress />` directly within your `.kiw` page or layout templates:

```html
<Progress class="my-custom-progress">
  <!-- Slot content or properties -->
</Progress>
```

---

## Supported Props

- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `showValue`: Configures the `showValue` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.
- `value`: Configures the `value` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

---
title: "Spinner"
description: "Accessible animated SVG loader spinner for pending operations."
date: "2026-11-09"
---

# Spinner

Category: **Feedback & Status**

Accessible animated SVG loader spinner for pending operations.

---

## Basic Usage

Use `<Spinner />` directly within your `.kiw` page or layout templates:

```html
<Spinner class="my-custom-spinner">
  <!-- Slot content or properties -->
</Spinner>
```

---

## Supported Props

- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

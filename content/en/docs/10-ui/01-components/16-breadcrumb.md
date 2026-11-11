---
title: "Breadcrumb"
description: "Hierarchical navigation trail indicating current page location."
date: "2026-11-09"
---

# Breadcrumb

Category: **Navigation & Structure**

Hierarchical navigation trail indicating current page location.

---

## Basic Usage

Use `<Breadcrumb />` directly within your `.kiw` page or layout templates:

```html
<Breadcrumb class="my-custom-breadcrumb">
  <!-- Slot content or properties -->
</Breadcrumb>
```

---

## Supported Props

- `id`: Configures the `id` attribute or behavior.
- `item`: Configures the `item` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

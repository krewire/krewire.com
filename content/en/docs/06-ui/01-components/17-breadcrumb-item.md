---
title: "BreadcrumbItem"
description: "Individual segment within a breadcrumb navigation trail."
date: "2026-11-09"
---

# BreadcrumbItem

Category: **Navigation & Structure**

Individual segment within a breadcrumb navigation trail.

---

## Basic Usage

Use `<BreadcrumbItem />` directly within your `.kiw` page or layout templates:

```html
<BreadcrumbItem class="my-custom-breadcrumb-item">
  <!-- Slot content or properties -->
</BreadcrumbItem>
```

---

## Supported Props

- `current`: Configures the `current` attribute or behavior.
- `href`: Configures the `href` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `target`: Configures the `target` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

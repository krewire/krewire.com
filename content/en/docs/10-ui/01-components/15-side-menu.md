---
title: "SideMenu"
description: "Vertical navigational list for administrative sidebars."
date: "2026-11-09"
---

# SideMenu

Category: **Navigation & Structure**

Vertical navigational list for administrative sidebars.

---

## Basic Usage

Use `<SideMenu />` directly within your `.kiw` page or layout templates:

```html
<SideMenu class="my-custom-side-menu">
  <!-- Slot content or properties -->
</SideMenu>
```

---

## Supported Props

- `g`: Configures the `g` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

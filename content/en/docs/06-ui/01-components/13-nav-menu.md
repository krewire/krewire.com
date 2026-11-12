---
title: "NavMenu"
description: "Horizontal navigation menu with interactive dropdown menus."
date: "2026-11-09"
---

# NavMenu

Category: **Navigation & Structure**

Horizontal navigation menu with interactive dropdown menus.

---

## Basic Usage

Use `<NavMenu />` directly within your `.kiw` page or layout templates:

```html
<NavMenu class="my-custom-nav-menu">
  <!-- Slot content or properties -->
</NavMenu>
```

---

## Supported Props

- `dropdown`: Configures the `dropdown` attribute or behavior.
- `links`: Configures the `links` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

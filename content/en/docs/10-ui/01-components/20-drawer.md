---
title: "Drawer"
description: "Smooth sliding drawer panel for mobile navigation and overlays."
date: "2026-11-09"
---

# Drawer

Category: **Navigation & Structure**

Smooth sliding drawer panel for mobile navigation and overlays.

---

## Basic Usage

Use `<Drawer />` directly within your `.kiw` page or layout templates:

```html
<Drawer class="my-custom-drawer">
  <!-- Slot content or properties -->
</Drawer>
```

---

## Supported Props

- `Body`: Configures the `Body` attribute or behavior.
- `Class`: Configures the `Class` attribute or behavior.
- `Footer`: Configures the `Footer` attribute or behavior.
- `ID`: Configures the `ID` attribute or behavior.
- `Side`: Configures the `Side` attribute or behavior.
- `Title`: Configures the `Title` attribute or behavior.
- `actions`: Configures the `actions` attribute or behavior.
- `footer`: Configures the `footer` attribute or behavior.
- `head`: Configures the `head` attribute or behavior.
- `header`: Configures the `header` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `side`: Configures the `side` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

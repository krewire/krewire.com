---
title: "Collapse"
description: "Collapsible single-panel container for toggling visibility."
date: "2026-11-09"
---

# Collapse

Category: **Data Display & Content**

Collapsible single-panel container for toggling visibility.

---

## Basic Usage

Use `<Collapse />` directly within your `.kiw` page or layout templates:

```html
<Collapse class="my-custom-collapse">
  <!-- Slot content or properties -->
</Collapse>
```

---

## Supported Props

- `badge`: Configures the `badge` attribute or behavior.
- `header`: Configures the `header` attribute or behavior.
- `icon`: Configures the `icon` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `name`: Configures the `name` attribute or behavior.
- `open`: Configures the `open` attribute or behavior.
- `summary`: Configures the `summary` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

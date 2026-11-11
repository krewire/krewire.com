---
title: "Sidebar"
description: "Collapsible sidebar navigation for documentation and dashboard workspaces."
date: "2026-11-09"
---

# Sidebar

Category: **Navigation & Structure**

Collapsible sidebar navigation for documentation and dashboard workspaces.

---

## Basic Usage

Use `<Sidebar />` directly within your `.kiw` page or layout templates:

```html
<Sidebar class="my-custom-sidebar">
  <!-- Slot content or properties -->
</Sidebar>
```

---

## Supported Props

- `head`: Configures the `head` attribute or behavior.
- `header`: Configures the `header` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `menu`: Configures the `menu` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

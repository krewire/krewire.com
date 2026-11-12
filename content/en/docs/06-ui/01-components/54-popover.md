---
title: "Popover"
description: "Floating contextual popover container anchored to target elements."
date: "2026-11-09"
---

# Popover

Category: **Overlays & Dialogs**

Floating contextual popover container anchored to target elements.

---

## Basic Usage

Use `<Popover />` directly within your `.kiw` page or layout templates:

```html
<Popover class="my-custom-popover">
  <!-- Slot content or properties -->
</Popover>
```

---

## Supported Props

- `content`: Configures the `content` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.
- `trigger`: Configures the `trigger` attribute or behavior.
- `triggerText`: Configures the `triggerText` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

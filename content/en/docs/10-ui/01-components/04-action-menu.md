---
title: "ActionMenu"
description: "Compact contextual dropdown action menu for table rows and cards."
date: "2026-11-09"
---

# ActionMenu

Category: **Actions & Buttons**

Compact contextual dropdown action menu for table rows and cards.

---

## Basic Usage

Use `<ActionMenu />` directly within your `.kiw` page or layout templates:

```html
<ActionMenu class="my-custom-action-menu">
  <!-- Slot content or properties -->
</ActionMenu>
```

---

## Supported Props

- `align`: Configures the `align` attribute or behavior.
- `item`: Configures the `item` attribute or behavior.
- `trigger`: Configures the `trigger` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

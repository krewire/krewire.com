---
title: "List"
description: "Styled unordered and ordered list container."
date: "2026-11-09"
---

# List

Category: **Data Display & Content**

Styled unordered and ordered list container.

---

## Basic Usage

Use `<List />` directly within your `.kiw` page or layout templates:

```html
<List class="my-custom-list">
  <!-- Slot content or properties -->
</List>
```

---

## Supported Props

- `id`: Configures the `id` attribute or behavior.
- `ordered`: Configures the `ordered` attribute or behavior.
- `spacing`: Configures the `spacing` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

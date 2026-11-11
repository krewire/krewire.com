---
title: "Code"
description: "Inline monospace code snippet with high contrast styling."
date: "2026-11-09"
---

# Code

Category: **Data Display & Content**

Inline monospace code snippet with high contrast styling.

---

## Basic Usage

Use `<Code />` directly within your `.kiw` page or layout templates:

```html
<Code class="my-custom-code">
  <!-- Slot content or properties -->
</Code>
```

---

## Supported Props

- `id`: Configures the `id` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

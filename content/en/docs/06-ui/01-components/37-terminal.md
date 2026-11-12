---
title: "Terminal"
description: "Single-line interactive terminal preview with copy command button."
date: "2026-11-09"
---

# Terminal

Category: **Data Display & Content**

Single-line interactive terminal preview with copy command button.

---

## Basic Usage

Use `<Terminal />` directly within your `.kiw` page or layout templates:

```html
<Terminal class="my-custom-terminal">
  <!-- Slot content or properties -->
</Terminal>
```

---

## Supported Props

- `actions`: Configures the `actions` attribute or behavior.
- `command`: Configures the `command` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `prompt`: Configures the `prompt` attribute or behavior.
- `promptColor`: Configures the `promptColor` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

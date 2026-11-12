---
title: "Table"
description: "Responsive data table with striped, bordered, and hover highlight states."
date: "2026-11-09"
---

# Table

Category: **Data Display & Content**

Responsive data table with striped, bordered, and hover highlight states.

---

## Basic Usage

Use `<Table />` directly within your `.kiw` page or layout templates:

```html
<Table class="my-custom-table">
  <!-- Slot content or properties -->
</Table>
```

---

## Supported Props

- `bordered`: Configures the `bordered` attribute or behavior.
- `cell`: Configures the `cell` attribute or behavior.
- `h`: Configures the `h` attribute or behavior.
- `headers`: Configures the `headers` attribute or behavior.
- `hover`: Configures the `hover` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `rows`: Configures the `rows` attribute or behavior.
- `striped`: Configures the `striped` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

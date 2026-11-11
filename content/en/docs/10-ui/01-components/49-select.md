---
title: "Select"
description: "Native styled dropdown select input."
date: "2026-11-09"
---

# Select

Category: **Forms & Inputs**

Native styled dropdown select input.

---

## Basic Usage

Use `<Select />` directly within your `.kiw` page or layout templates:

```html
<Select class="my-custom-select">
  <!-- Slot content or properties -->
</Select>
```

---

## Supported Props

- `disabled`: Configures the `disabled` attribute or behavior.
- `error`: Configures the `error` attribute or behavior.
- `help`: Configures the `help` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `name`: Configures the `name` attribute or behavior.
- `required`: Configures the `required` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

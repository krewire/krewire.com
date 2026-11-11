---
title: "Radio"
description: "Single-choice radio button with pop-brutalist circle styling."
date: "2026-11-09"
---

# Radio

Category: **Forms & Inputs**

Single-choice radio button with pop-brutalist circle styling.

---

## Basic Usage

Use `<Radio />` directly within your `.kiw` page or layout templates:

```html
<Radio class="my-custom-radio">
  <!-- Slot content or properties -->
</Radio>
```

---

## Supported Props

- `checked`: Configures the `checked` attribute or behavior.
- `disabled`: Configures the `disabled` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `name`: Configures the `name` attribute or behavior.
- `required`: Configures the `required` attribute or behavior.
- `value`: Configures the `value` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

---
title: "Textarea"
description: "Multi-line text input field with auto-resizing support."
date: "2026-11-09"
---

# Textarea

Category: **Forms & Inputs**

Multi-line text input field with auto-resizing support.

---

## Basic Usage

Use `<Textarea />` directly within your `.kiw` page or layout templates:

```html
<Textarea class="my-custom-textarea">
  <!-- Slot content or properties -->
</Textarea>
```

---

## Supported Props

- `disabled`: Configures the `disabled` attribute or behavior.
- `error`: Configures the `error` attribute or behavior.
- `help`: Configures the `help` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `name`: Configures the `name` attribute or behavior.
- `placeholder`: Configures the `placeholder` attribute or behavior.
- `required`: Configures the `required` attribute or behavior.
- `rows`: Configures the `rows` attribute or behavior.
- `value`: Configures the `value` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

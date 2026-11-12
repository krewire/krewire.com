---
title: "Form"
description: "Accessible form wrapper with method spoofing and submission states."
date: "2026-11-09"
---

# Form

Category: **Forms & Inputs**

Accessible form wrapper with method spoofing and submission states.

---

## Basic Usage

Use `<Form />` directly within your `.kiw` page or layout templates:

```html
<Form class="my-custom-form">
  <!-- Slot content or properties -->
</Form>
```

---

## Supported Props

- `action`: Configures the `action` attribute or behavior.
- `enctype`: Configures the `enctype` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `method`: Configures the `method` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

---
title: "Slider"
description: "Range slider input with numeric bubble value display."
date: "2026-11-09"
---

# Slider

Category: **Forms & Inputs**

Range slider input with numeric bubble value display.

---

## Basic Usage

Use `<Slider />` directly within your `.kiw` page or layout templates:

```html
<Slider class="my-custom-slider">
  <!-- Slot content or properties -->
</Slider>
```

---

## Supported Props

- `disabled`: Configures the `disabled` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `max`: Configures the `max` attribute or behavior.
- `min`: Configures the `min` attribute or behavior.
- `name`: Configures the `name` attribute or behavior.
- `onInput`: Configures the `onInput` attribute or behavior.
- `showValue`: Configures the `showValue` attribute or behavior.
- `step`: Configures the `step` attribute or behavior.
- `value`: Configures the `value` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

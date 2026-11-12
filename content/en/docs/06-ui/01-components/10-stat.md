---
title: "Stat"
description: "Key performance indicator card with numeric metrics and trend indicators."
date: "2026-11-09"
---

# Stat

Category: **Data Display & Content**

Key performance indicator card with numeric metrics and trend indicators.

---

## Basic Usage

Use `<Stat />` directly within your `.kiw` page or layout templates:

```html
<Stat class="my-custom-stat">
  <!-- Slot content or properties -->
</Stat>
```

---

## Supported Props

- `change`: Configures the `change` attribute or behavior.
- `help`: Configures the `help` attribute or behavior.
- `icon`: Configures the `icon` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `subtitle`: Configures the `subtitle` attribute or behavior.
- `trend`: Configures the `trend` attribute or behavior.
- `unit`: Configures the `unit` attribute or behavior.
- `value`: Configures the `value` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

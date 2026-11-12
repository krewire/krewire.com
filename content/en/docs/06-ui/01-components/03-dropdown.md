---
title: "Dropdown"
description: "Zero-dependency dropdown menu built on native details and summary elements."
date: "2026-11-09"
---

# Dropdown

Category: **Actions & Buttons**

Zero-dependency dropdown menu built on native details and summary elements.

---

## Basic Usage

Use `<Dropdown />` directly within your `.kiw` page or layout templates:

```html
<Dropdown class="my-custom-dropdown">
  <!-- Slot content or properties -->
</Dropdown>
```

---

## Supported Props

- `align`: Configures the `align` attribute or behavior.
- `icon`: Optional leading icon rendered before the label in default trigger button.
- `item`: Configures the `item` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `trigger`: Configures the `trigger` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

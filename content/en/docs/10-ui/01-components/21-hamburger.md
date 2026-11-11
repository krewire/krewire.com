---
title: "Hamburger"
description: "Animated 3-line hamburger menu toggle button for mobile drawers."
date: "2026-11-09"
---

# Hamburger

Category: **Navigation & Structure**

Animated 3-line hamburger menu toggle button for mobile drawers.

---

## Basic Usage

Use `<Hamburger />` directly within your `.kiw` page or layout templates:

```html
<Hamburger class="my-custom-hamburger">
  <!-- Slot content or properties -->
</Hamburger>
```

---

## Supported Props

- `Class`: Configures the `Class` attribute or behavior.
- `Controls`: Configures the `Controls` attribute or behavior.
- `DataDrawerToggle`: Configures the `DataDrawerToggle` attribute or behavior.
- `Expanded`: Configures the `Expanded` attribute or behavior.
- `ID`: Configures the `ID` attribute or behavior.
- `Label`: Configures the `Label` attribute or behavior.
- `controls`: Configures the `controls` attribute or behavior.
- `dataDrawerToggle`: Configures the `dataDrawerToggle` attribute or behavior.
- `expanded`: Configures the `expanded` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

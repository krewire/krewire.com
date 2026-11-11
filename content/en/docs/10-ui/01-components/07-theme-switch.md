---
title: "ThemeSwitch"
description: "Dark and light mode toggle button with localStorage persistence."
date: "2026-11-09"
---

# ThemeSwitch

Category: **Actions & Buttons**

Dark and light mode toggle button with localStorage persistence.

---

## Basic Usage

Use `<ThemeSwitch />` directly within your `.kiw` page or layout templates:

```html
<ThemeSwitch class="my-custom-theme-switch">
  <!-- Slot content or properties -->
</ThemeSwitch>
```

---

## Supported Props

- None (accepts default slot content or children).

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

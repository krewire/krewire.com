---
title: "LangSwitch"
description: "Language selector switch for multi-lingual internationalized applications."
date: "2026-11-09"
---

# LangSwitch

Category: **Actions & Buttons**

Language selector switch for multi-lingual internationalized applications.

---

## Basic Usage

Use `<LangSwitch />` directly within your `.kiw` page or layout templates:

```html
<LangSwitch class="my-custom-lang-switch">
  <!-- Slot content or properties -->
</LangSwitch>
```

---

## Supported Props

- `code`: Configures the `code` attribute or behavior.
- `icon`: Configures the `icon` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

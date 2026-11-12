---
title: "Brand"
description: "Ecosystem brand mark with logo icon and version badge."
date: "2026-11-09"
---

# Brand

Category: **Navigation & Structure**

Ecosystem brand mark with logo icon and version badge.

---

## Basic Usage

Use `<Brand />` directly within your `.kiw` page or layout templates:

```html
<Brand class="my-custom-brand">
  <!-- Slot content or properties -->
</Brand>
```

---

## Supported Props

- `badge`: Configures the `badge` attribute or behavior.
- `badgeI18n`: Configures the `badgeI18n` attribute or behavior.
- `href`: Configures the `href` attribute or behavior.
- `icon`: Configures the `icon` attribute or behavior.
- `name`: Configures the `name` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.
- `version`: Configures the `version` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

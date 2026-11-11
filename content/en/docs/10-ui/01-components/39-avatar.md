---
title: "Avatar"
description: "User profile avatar displaying image with automatic fallback initials."
date: "2026-11-09"
---

# Avatar

Category: **Data Display & Content**

User profile avatar displaying image with automatic fallback initials.

---

## Basic Usage

Use `<Avatar />` directly within your `.kiw` page or layout templates:

```html
<Avatar class="my-custom-avatar">
  <!-- Slot content or properties -->
</Avatar>
```

---

## Supported Props

- `alt`: Configures the `alt` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `initials`: Configures the `initials` attribute or behavior.
- `shape`: Configures the `shape` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.
- `src`: Configures the `src` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

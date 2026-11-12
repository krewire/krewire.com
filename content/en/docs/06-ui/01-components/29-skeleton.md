---
title: "Skeleton"
description: "Animated shimmer placeholder for asynchronous loading states."
date: "2026-11-09"
---

# Skeleton

Category: **Feedback & Status**

Animated shimmer placeholder for asynchronous loading states.

---

## Basic Usage

Use `<Skeleton />` directly within your `.kiw` page or layout templates:

```html
<Skeleton class="my-custom-skeleton">
  <!-- Slot content or properties -->
</Skeleton>
```

---

## Supported Props

- `height`: Configures the `height` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.
- `width`: Configures the `width` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

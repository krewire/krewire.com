---
title: "Toast"
description: "Transient notification toast banner with auto-dismiss timer."
date: "2026-11-09"
---

# Toast

Category: **Feedback & Status**

Transient notification toast banner with auto-dismiss timer.

---

## Basic Usage

Use `<Toast />` directly within your `.kiw` page or layout templates:

```html
<Toast class="my-custom-toast">
  <!-- Slot content or properties -->
</Toast>
```

---

## Supported Props

- `dismissible`: Configures the `dismissible` attribute or behavior.
- `icon`: Configures the `icon` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `message`: Configures the `message` attribute or behavior.
- `position`: Configures the `position` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

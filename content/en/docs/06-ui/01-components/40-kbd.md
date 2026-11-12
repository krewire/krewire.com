---
title: "Kbd"
description: "Keyboard key visual indicator for shortcuts and command hotkeys."
date: "2026-11-09"
---

# Kbd

Category: **Data Display & Content**

Keyboard key visual indicator for shortcuts and command hotkeys.

---

## Basic Usage

Use `<Kbd />` directly within your `.kiw` page or layout templates:

```html
<Kbd class="my-custom-kbd">
  <!-- Slot content or properties -->
</Kbd>
```

---

## Supported Props

- `id`: Configures the `id` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

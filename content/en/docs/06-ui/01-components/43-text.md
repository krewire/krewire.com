---
title: "Text"
description: "Typography wrapper for body paragraphs, lead text, and muted captions."
date: "2026-11-09"
---

# Text

Category: **Data Display & Content**

Typography wrapper for body paragraphs, lead text, and muted captions.

---

## Basic Usage

Use `<Text />` directly within your `.kiw` page or layout templates:

```html
<Text class="my-custom-text">
  <!-- Slot content or properties -->
</Text>
```

---

## Supported Props

- `color`: Configures the `color` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.
- `weight`: Configures the `weight` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

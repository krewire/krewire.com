---
title: "Tag"
description: "Removable or clickable keyword tag with color variants."
date: "2026-11-09"
---

# Tag

Category: **Feedback & Status**

Removable or clickable keyword tag with color variants.

---

## Basic Usage

Use `<Tag />` directly within your `.kiw` page or layout templates:

```html
<Tag class="my-custom-tag">
  <!-- Slot content or properties -->
</Tag>
```

---

## Supported Props

- `icon`: Configures the `icon` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `removable`: Configures the `removable` attribute or behavior.
- `selected`: Configures the `selected` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

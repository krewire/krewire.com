---
title: "Callout"
description: "Bordered callout container for documentation tips, warnings, and notes."
date: "2026-11-09"
---

# Callout

Category: **Feedback & Status**

Bordered callout container for documentation tips, warnings, and notes.

---

## Basic Usage

Use `<Callout />` directly within your `.kiw` page or layout templates:

```html
<Callout class="my-custom-callout">
  <!-- Slot content or properties -->
</Callout>
```

---

## Supported Props

- `title`: Configures the `title` attribute or behavior.
- `type`: Configures the `type` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

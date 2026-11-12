---
title: "Container"
description: "Centered responsive page container maintaining standard max-width boundaries."
date: "2026-11-09"
---

# Container

Category: **Layout Primitives**

Centered responsive page container maintaining standard max-width boundaries.

---

## Basic Usage

Use `<Container />` directly within your `.kiw` page or layout templates:

```html
<Container class="my-custom-container">
  <!-- Slot content or properties -->
</Container>
```

---

## Supported Props

- `id`: Configures the `id` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

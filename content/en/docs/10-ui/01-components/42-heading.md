---
title: "Heading"
description: "Semantic heading typography wrapper supporting h1 through h6."
date: "2026-11-09"
---

# Heading

Category: **Data Display & Content**

Semantic heading typography wrapper supporting h1 through h6.

---

## Basic Usage

Use `<Heading />` directly within your `.kiw` page or layout templates:

```html
<Heading class="my-custom-heading">
  <!-- Slot content or properties -->
</Heading>
```

---

## Supported Props

- `id`: Configures the `id` attribute or behavior.
- `level`: Configures the `level` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

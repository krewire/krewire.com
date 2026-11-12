---
title: "Head"
description: "Document head injector for standard meta tags, OpenGraph, and favicon assets."
date: "2026-11-09"
---

# Head

Category: **Navigation & Structure**

Document head injector for standard meta tags, OpenGraph, and favicon assets.

---

## Basic Usage

Use `<Head />` directly within your `.kiw` page or layout templates:

```html
<Head class="my-custom-head">
  <!-- Slot content or properties -->
</Head>
```

---

## Supported Props

- `assetversion`: Configures the `assetversion` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

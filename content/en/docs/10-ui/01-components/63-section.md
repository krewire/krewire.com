---
title: "Section"
description: "Full-width semantic section wrapper with vertical padding tokens."
date: "2026-11-09"
---

# Section

Category: **Layout Primitives**

Full-width semantic section wrapper with vertical padding tokens.

---

## Basic Usage

Use `<Section />` directly within your `.kiw` page or layout templates:

```html
<Section class="my-custom-section">
  <!-- Slot content or properties -->
</Section>
```

---

## Supported Props

- `content`: Configures the `content` attribute or behavior.
- `desc`: Configures the `desc` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

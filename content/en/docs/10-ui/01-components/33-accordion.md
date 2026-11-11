---
title: "Accordion"
description: "Expandable and collapsible content panels with keyboard accessibility."
date: "2026-11-09"
---

# Accordion

Category: **Data Display & Content**

Expandable and collapsible content panels with keyboard accessibility.

---

## Basic Usage

Use `<Accordion />` directly within your `.kiw` page or layout templates:

```html
<Accordion class="my-custom-accordion">
  <!-- Slot content or properties -->
</Accordion>
```

---

## Supported Props

- `item`: Configures the `item` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

---
title: "Timeline"
description: "Chronological vertical activity and event timeline."
date: "2026-11-09"
---

# Timeline

Category: **Data Display & Content**

Chronological vertical activity and event timeline.

---

## Basic Usage

Use `<Timeline />` directly within your `.kiw` page or layout templates:

```html
<Timeline class="my-custom-timeline">
  <!-- Slot content or properties -->
</Timeline>
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

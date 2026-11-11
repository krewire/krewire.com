---
title: "Badge"
description: "Compact pill badge with color variants for tags, counts, and status indicators."
date: "2026-11-09"
---

# Badge

Category: **Feedback & Status**

Compact pill badge with color variants for tags, counts, and status indicators.

---

## Live Preview

<div data-component-preview="Badge" data-slug="badge"></div>

The frames above load the real `Badge.kiw` from Krewire UI, prebuilt by `kiw build` for every variant, in the light and dark themes side by side. Use the toggles to switch variants, then copy the component source or the usage snippet. **Open in playground** shows it next to the other core components.

---

## Basic Usage

Use `<Badge />` directly within your `.kiw` page or layout templates:

```html
<Badge class="my-custom-badge">
  <!-- Slot content or properties -->
</Badge>
```

---

## Supported Props

- `dot`: Configures the `dot` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

---
title: "Link"
description: "Accessible hyperlink with underline transitions and target security."
date: "2026-11-09"
---

# Link

Category: **Actions & Buttons**

Accessible hyperlink with underline transitions and target security.

---

## Basic Usage

Use `<Link />` directly within your `.kiw` page or layout templates:

```html
<Link class="my-custom-link">
  <!-- Slot content or properties -->
</Link>
```

---

## Supported Props

- `external`: Configures the `external` attribute or behavior.
- `href`: Configures the `href` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

---
title: "Banner"
description: "Prominent top-level announcement bar with dismiss trigger."
date: "2026-11-09"
---

# Banner

Category: **Feedback & Status**

Prominent top-level announcement bar with dismiss trigger.

---

## Basic Usage

Use `<Banner />` directly within your `.kiw` page or layout templates:

```html
<Banner class="my-custom-banner">
  <!-- Slot content or properties -->
</Banner>
```

---

## Supported Props

- `action`: Configures the `action` attribute or behavior.
- `badge`: Configures the `badge` attribute or behavior.
- `cta`: Configures the `cta` attribute or behavior.
- `dismissible`: Configures the `dismissible` attribute or behavior.
- `href`: Configures the `href` attribute or behavior.
- `icon`: Configures the `icon` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `message`: Configures the `message` attribute or behavior.
- `target`: Configures the `target` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

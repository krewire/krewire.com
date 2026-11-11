---
title: "Alert"
description: "Contextual feedback banner supporting info, success, warning, and error states."
date: "2026-11-09"
---

# Alert

Category: **Feedback & Status**

Contextual feedback banner supporting info, success, warning, and error states.

---

## Live Preview

<div data-component-preview="Alert" data-slug="alert"></div>

The frames above load the real `Alert.kiw` from Krewire UI, prebuilt by `kiw build` for every variant, in the light and dark themes side by side. Use the toggles to switch variants, then copy the component source or the usage snippet. **Open in playground** shows it next to the other core components.

---

## Basic Usage

Use `<Alert />` directly within your `.kiw` page or layout templates:

```html
<Alert class="my-custom-alert">
  <!-- Slot content or properties -->
</Alert>
```

---

## Supported Props

- `action`: Configures the `action` attribute or behavior.
- `actions`: Configures the `actions` attribute or behavior.
- `dismissible`: Configures the `dismissible` attribute or behavior.
- `icon`: Configures the `icon` attribute or behavior.
- `message`: Configures the `message` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

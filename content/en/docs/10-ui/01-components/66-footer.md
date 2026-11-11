---
title: "Footer"
description: "Comprehensive multi-column footer with brand, navigation links, and copyright."
date: "2026-11-09"
---

# Footer

Category: **Navigation & Structure**

Comprehensive multi-column footer with brand, navigation links, and copyright.

---

## Basic Usage

Use `<Footer />` directly within your `.kiw` page or layout templates:

```html
<Footer class="my-custom-footer">
  <!-- Slot content or properties -->
</Footer>
```

---

## Supported Props

- `bottom`: Configures the `bottom` attribute or behavior.
- `brand`: Configures the `brand` attribute or behavior.
- `brandBadge`: Configures the `brandBadge` attribute or behavior.
- `brandIcon`: Configures the `brandIcon` attribute or behavior.
- `brandName`: Configures the `brandName` attribute or behavior.
- `brandTagline`: Configures the `brandTagline` attribute or behavior.
- `columns`: Configures the `columns` attribute or behavior.
- `copyrightText`: Configures the `copyrightText` attribute or behavior.
- `gitHubURL`: Configures the `gitHubURL` attribute or behavior.
- `showEcosystemFooter`: Configures the `showEcosystemFooter` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

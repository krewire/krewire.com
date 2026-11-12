---
title: "Navbar"
description: "Main application header with brand logo, nav links, and drawer trigger."
date: "2026-11-09"
---

# Navbar

Category: **Navigation & Structure**

Main application header with brand logo, nav links, and drawer trigger.

---

## Basic Usage

Use `<Navbar />` directly within your `.kiw` page or layout templates:

```html
<Navbar class="my-custom-navbar">
  <!-- Slot content or properties -->
</Navbar>
```

---

## Supported Props

- `actions`: Configures the `actions` attribute or behavior.
- `brand`: Configures the `brand` attribute or behavior.
- `brandBadge`: Configures the `brandBadge` attribute or behavior.
- `brandBadgeI18n`: Configures the `brandBadgeI18n` attribute or behavior.
- `brandHref`: Configures the `brandHref` attribute or behavior.
- `brandIcon`: Configures the `brandIcon` attribute or behavior.
- `brandName`: Configures the `brandName` attribute or behavior.
- `drawer`: Configures the `drawer` attribute or behavior.
- `drawerBody`: Configures the `drawerBody` attribute or behavior.
- `gitHubURL`: Configures the `gitHubURL` attribute or behavior.
- `showLangToggle`: Configures the `showLangToggle` attribute or behavior.
- `showSidebarToggle`: Configures the `showSidebarToggle` attribute or behavior.
- `version`: Configures the `version` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

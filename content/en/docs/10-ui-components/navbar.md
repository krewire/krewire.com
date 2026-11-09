---
title: "Navbar"
description: "Main application header with brand logo, nav links, and drawer trigger."
date: "2026-11-09"
---

# Navbar

Category: **Navigation & Structure**  
Component File: `components/Navbar.kiw`  
Source: `github.com/krewire/ui/components/Navbar.kiw`

Main application header with brand logo, nav links, and drawer trigger.

---

## 1. Copy & Own Installation

Copy the `Navbar.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Navbar.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Navbar />` directly within your `.kiw` page or layout templates:

```html
<Navbar class="my-custom-navbar">
  <!-- Slot content or properties -->
</Navbar>
```

---

## 3. Supported Props & Attributes

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

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

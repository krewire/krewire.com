---
title: "Footer"
description: "Comprehensive multi-column footer with brand, navigation links, and copyright."
date: "2026-11-09"
---

# Footer

Category: **Navigation & Structure**  
Component File: `components/Footer.kiw`  
Source: `github.com/krewire/ui/components/Footer.kiw`

Comprehensive multi-column footer with brand, navigation links, and copyright.

---

## 1. Copy & Own Installation

Copy the `Footer.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Footer.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Footer />` directly within your `.kiw` page or layout templates:

```html
<Footer class="my-custom-footer">
  <!-- Slot content or properties -->
</Footer>
```

---

## 3. Supported Props & Attributes

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

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

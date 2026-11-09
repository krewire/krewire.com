---
title: "SideMenu"
description: "Vertical navigational list for administrative sidebars."
date: "2026-11-09"
---

# SideMenu

Category: **Navigation & Structure**  
Component File: `components/SideMenu.kiw`  
Source: `github.com/krewire/ui/components/SideMenu.kiw`

Vertical navigational list for administrative sidebars.

---

## 1. Copy & Own Installation

Copy the `SideMenu.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/SideMenu.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<SideMenu />` directly within your `.kiw` page or layout templates:

```html
<SideMenu class="my-custom-side-menu">
  <!-- Slot content or properties -->
</SideMenu>
```

---

## 3. Supported Props & Attributes

- `g`: Configures the `g` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

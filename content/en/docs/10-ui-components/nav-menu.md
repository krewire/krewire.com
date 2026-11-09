---
title: "NavMenu"
description: "Horizontal navigation menu with interactive dropdown menus."
date: "2026-11-09"
---

# NavMenu

Category: **Navigation & Structure**  
Component File: `components/NavMenu.kiw`  
Source: `github.com/krewire/ui/components/NavMenu.kiw`

Horizontal navigation menu with interactive dropdown menus.

---

## 1. Copy & Own Installation

Copy the `NavMenu.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/NavMenu.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<NavMenu />` directly within your `.kiw` page or layout templates:

```html
<NavMenu class="my-custom-nav-menu">
  <!-- Slot content or properties -->
</NavMenu>
```

---

## 3. Supported Props & Attributes

- `dropdown`: Configures the `dropdown` attribute or behavior.
- `links`: Configures the `links` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

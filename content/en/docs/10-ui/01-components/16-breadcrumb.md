---
title: "Breadcrumb"
description: "Hierarchical navigation trail indicating current page location."
date: "2026-11-09"
---

# Breadcrumb

Category: **Navigation & Structure**  
Component File: `components/Breadcrumb.kiw`  
Source: `github.com/krewire/ui/components/Breadcrumb.kiw`

Hierarchical navigation trail indicating current page location.

---

## 1. Copy & Own Installation

Copy the `Breadcrumb.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Breadcrumb.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Breadcrumb />` directly within your `.kiw` page or layout templates:

```html
<Breadcrumb class="my-custom-breadcrumb">
  <!-- Slot content or properties -->
</Breadcrumb>
```

---

## 3. Supported Props & Attributes

- `id`: Configures the `id` attribute or behavior.
- `item`: Configures the `item` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

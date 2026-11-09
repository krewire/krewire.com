---
title: "BreadcrumbItem"
description: "Individual segment within a breadcrumb navigation trail."
date: "2026-11-09"
---

# BreadcrumbItem

Category: **Navigation & Structure**  
Component File: `components/BreadcrumbItem.kiw`  
Source: `github.com/krewire/ui/components/BreadcrumbItem.kiw`

Individual segment within a breadcrumb navigation trail.

---

## 1. Copy & Own Installation

Copy the `BreadcrumbItem.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/BreadcrumbItem.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<BreadcrumbItem />` directly within your `.kiw` page or layout templates:

```html
<BreadcrumbItem class="my-custom-breadcrumb-item">
  <!-- Slot content or properties -->
</BreadcrumbItem>
```

---

## 3. Supported Props & Attributes

- `current`: Configures the `current` attribute or behavior.
- `href`: Configures the `href` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `target`: Configures the `target` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

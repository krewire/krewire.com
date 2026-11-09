---
title: "Divider"
description: "Horizontal or vertical separator line with optional text label."
date: "2026-11-09"
---

# Divider

Category: **Data Display & Content**  
Component File: `components/Divider.kiw`  
Source: `github.com/krewire/ui/components/Divider.kiw`

Horizontal or vertical separator line with optional text label.

---

## 1. Copy & Own Installation

Copy the `Divider.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Divider.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Divider />` directly within your `.kiw` page or layout templates:

```html
<Divider class="my-custom-divider">
  <!-- Slot content or properties -->
</Divider>
```

---

## 3. Supported Props & Attributes

- `label`: Configures the `label` attribute or behavior.
- `vertical`: Configures the `vertical` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

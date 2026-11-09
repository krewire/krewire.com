---
title: "Box"
description: "Versatile layout box supporting border, padding, and pop-brutalist shadows."
date: "2026-11-09"
---

# Box

Category: **Layout Primitives**  
Component File: `components/Box.kiw`  
Source: `github.com/krewire/ui/components/Box.kiw`

Versatile layout box supporting border, padding, and pop-brutalist shadows.

---

## 1. Copy & Own Installation

Copy the `Box.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Box.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Box />` directly within your `.kiw` page or layout templates:

```html
<Box class="my-custom-box">
  <!-- Slot content or properties -->
</Box>
```

---

## 3. Supported Props & Attributes

- `border`: Configures the `border` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `rounded`: Configures the `rounded` attribute or behavior.
- `shadow`: Configures the `shadow` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

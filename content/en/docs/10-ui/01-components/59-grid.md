---
title: "Grid"
description: "CSS Grid container with configurable columns and gap tokens."
date: "2026-11-09"
---

# Grid

Category: **Layout Primitives**  
Component File: `components/Grid.kiw`  
Source: `github.com/krewire/ui/components/Grid.kiw`

CSS Grid container with configurable columns and gap tokens.

---

## 1. Copy & Own Installation

Copy the `Grid.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Grid.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Grid />` directly within your `.kiw` page or layout templates:

```html
<Grid class="my-custom-grid">
  <!-- Slot content or properties -->
</Grid>
```

---

## 3. Supported Props & Attributes

- `cols`: Configures the `cols` attribute or behavior.
- `gap`: Configures the `gap` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

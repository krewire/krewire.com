---
title: "Spinner"
description: "Accessible animated SVG loader spinner for pending operations."
date: "2026-11-09"
---

# Spinner

Category: **Feedback & Status**  
Component File: `components/Spinner.kiw`  
Source: `github.com/krewire/ui/components/Spinner.kiw`

Accessible animated SVG loader spinner for pending operations.

---

## 1. Copy & Own Installation

Copy the `Spinner.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Spinner.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Spinner />` directly within your `.kiw` page or layout templates:

```html
<Spinner class="my-custom-spinner">
  <!-- Slot content or properties -->
</Spinner>
```

---

## 3. Supported Props & Attributes

- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

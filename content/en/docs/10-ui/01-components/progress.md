---
title: "Progress"
description: "Determinate and indeterminate progress bar indicator."
date: "2026-11-09"
---

# Progress

Category: **Feedback & Status**  
Component File: `components/Progress.kiw`  
Source: `github.com/krewire/ui/components/Progress.kiw`

Determinate and indeterminate progress bar indicator.

---

## 1. Copy & Own Installation

Copy the `Progress.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Progress.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Progress />` directly within your `.kiw` page or layout templates:

```html
<Progress class="my-custom-progress">
  <!-- Slot content or properties -->
</Progress>
```

---

## 3. Supported Props & Attributes

- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `showValue`: Configures the `showValue` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.
- `value`: Configures the `value` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

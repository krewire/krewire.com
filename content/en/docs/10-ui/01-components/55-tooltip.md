---
title: "Tooltip"
description: "Accessible hover and focus informative tooltip."
date: "2026-11-09"
---

# Tooltip

Category: **Overlays & Dialogs**  
Component File: `components/Tooltip.kiw`  
Source: `github.com/krewire/ui/components/Tooltip.kiw`

Accessible hover and focus informative tooltip.

---

## 1. Copy & Own Installation

Copy the `Tooltip.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Tooltip.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Tooltip />` directly within your `.kiw` page or layout templates:

```html
<Tooltip class="my-custom-tooltip">
  <!-- Slot content or properties -->
</Tooltip>
```

---

## 3. Supported Props & Attributes

- `id`: Configures the `id` attribute or behavior.
- `placement`: Configures the `placement` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

---
title: "ButtonGroup"
description: "Grouped button segment with unified borders and active states."
date: "2026-11-09"
---

# ButtonGroup

Category: **Actions & Buttons**  
Component File: `components/ButtonGroup.kiw`  
Source: `github.com/krewire/ui/components/ButtonGroup.kiw`

Grouped button segment with unified borders and active states.

---

## 1. Copy & Own Installation

Copy the `ButtonGroup.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/ButtonGroup.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<ButtonGroup />` directly within your `.kiw` page or layout templates:

```html
<ButtonGroup class="my-custom-button-group">
  <!-- Slot content or properties -->
</ButtonGroup>
```

---

## 3. Supported Props & Attributes

- `ariaLabel`: Configures the `ariaLabel` attribute or behavior.
- `item`: Configures the `item` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

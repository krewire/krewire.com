---
title: "Dropdown"
description: "Zero-dependency dropdown menu built on native details and summary elements."
date: "2026-11-09"
---

# Dropdown

Category: **Actions & Buttons**  
Component File: `components/Dropdown.kiw`  
Source: `github.com/krewire/ui/components/Dropdown.kiw`

Zero-dependency dropdown menu built on native details and summary elements.

---

## 1. Copy & Own Installation

Copy the `Dropdown.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Dropdown.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Dropdown />` directly within your `.kiw` page or layout templates:

```html
<Dropdown class="my-custom-dropdown">
  <!-- Slot content or properties -->
</Dropdown>
```

---

## 3. Supported Props & Attributes

- `align`: Configures the `align` attribute or behavior.
- `icon`: Optional leading icon rendered before the label in default trigger button.
- `item`: Configures the `item` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `trigger`: Configures the `trigger` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

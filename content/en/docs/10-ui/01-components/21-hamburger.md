---
title: "Hamburger"
description: "Animated 3-line hamburger menu toggle button for mobile drawers."
date: "2026-11-09"
---

# Hamburger

Category: **Navigation & Structure**  
Component File: `components/Hamburger.kiw`  
Source: `github.com/krewire/ui/components/Hamburger.kiw`

Animated 3-line hamburger menu toggle button for mobile drawers.

---

## 1. Copy & Own Installation

Copy the `Hamburger.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Hamburger.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Hamburger />` directly within your `.kiw` page or layout templates:

```html
<Hamburger class="my-custom-hamburger">
  <!-- Slot content or properties -->
</Hamburger>
```

---

## 3. Supported Props & Attributes

- `Class`: Configures the `Class` attribute or behavior.
- `Controls`: Configures the `Controls` attribute or behavior.
- `DataDrawerToggle`: Configures the `DataDrawerToggle` attribute or behavior.
- `Expanded`: Configures the `Expanded` attribute or behavior.
- `ID`: Configures the `ID` attribute or behavior.
- `Label`: Configures the `Label` attribute or behavior.
- `controls`: Configures the `controls` attribute or behavior.
- `dataDrawerToggle`: Configures the `dataDrawerToggle` attribute or behavior.
- `expanded`: Configures the `expanded` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

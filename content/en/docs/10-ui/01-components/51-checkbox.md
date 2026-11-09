---
title: "Checkbox"
description: "High-contrast accessible checkbox input with custom pop-border styling."
date: "2026-11-09"
---

# Checkbox

Category: **Forms & Inputs**  
Component File: `components/Checkbox.kiw`  
Source: `github.com/krewire/ui/components/Checkbox.kiw`

High-contrast accessible checkbox input with custom pop-border styling.

---

## 1. Copy & Own Installation

Copy the `Checkbox.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Checkbox.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Checkbox />` directly within your `.kiw` page or layout templates:

```html
<Checkbox class="my-custom-checkbox">
  <!-- Slot content or properties -->
</Checkbox>
```

---

## 3. Supported Props & Attributes

- `checked`: Configures the `checked` attribute or behavior.
- `disabled`: Configures the `disabled` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `name`: Configures the `name` attribute or behavior.
- `required`: Configures the `required` attribute or behavior.
- `value`: Configures the `value` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

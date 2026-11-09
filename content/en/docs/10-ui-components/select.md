---
title: "Select"
description: "Native styled dropdown select input."
date: "2026-11-09"
---

# Select

Category: **Forms & Inputs**  
Component File: `components/Select.kiw`  
Source: `github.com/krewire/ui/components/Select.kiw`

Native styled dropdown select input.

---

## 1. Copy & Own Installation

Copy the `Select.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Select.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Select />` directly within your `.kiw` page or layout templates:

```html
<Select class="my-custom-select">
  <!-- Slot content or properties -->
</Select>
```

---

## 3. Supported Props & Attributes

- `disabled`: Configures the `disabled` attribute or behavior.
- `error`: Configures the `error` attribute or behavior.
- `help`: Configures the `help` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `name`: Configures the `name` attribute or behavior.
- `required`: Configures the `required` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

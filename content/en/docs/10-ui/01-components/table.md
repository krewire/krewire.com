---
title: "Table"
description: "Responsive data table with striped, bordered, and hover highlight states."
date: "2026-11-09"
---

# Table

Category: **Data Display & Content**  
Component File: `components/Table.kiw`  
Source: `github.com/krewire/ui/components/Table.kiw`

Responsive data table with striped, bordered, and hover highlight states.

---

## 1. Copy & Own Installation

Copy the `Table.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Table.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Table />` directly within your `.kiw` page or layout templates:

```html
<Table class="my-custom-table">
  <!-- Slot content or properties -->
</Table>
```

---

## 3. Supported Props & Attributes

- `bordered`: Configures the `bordered` attribute or behavior.
- `cell`: Configures the `cell` attribute or behavior.
- `h`: Configures the `h` attribute or behavior.
- `headers`: Configures the `headers` attribute or behavior.
- `hover`: Configures the `hover` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `rows`: Configures the `rows` attribute or behavior.
- `striped`: Configures the `striped` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

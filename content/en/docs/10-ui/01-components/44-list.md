---
title: "List"
description: "Styled unordered and ordered list container."
date: "2026-11-09"
---

# List

Category: **Data Display & Content**  
Component File: `components/List.kiw`  
Source: `github.com/krewire/ui/components/List.kiw`

Styled unordered and ordered list container.

---

## 1. Copy & Own Installation

Copy the `List.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/List.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<List />` directly within your `.kiw` page or layout templates:

```html
<List class="my-custom-list">
  <!-- Slot content or properties -->
</List>
```

---

## 3. Supported Props & Attributes

- `id`: Configures the `id` attribute or behavior.
- `ordered`: Configures the `ordered` attribute or behavior.
- `spacing`: Configures the `spacing` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

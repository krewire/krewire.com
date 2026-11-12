---
title: "SearchInput"
description: "Search box with search icon and instant clear button."
date: "2026-11-09"
---

# SearchInput

Category: **Forms & Inputs**

Search box with search icon and instant clear button.

---

## Basic Usage

Use `<SearchInput />` directly within your `.kiw` page or layout templates:

```html
<SearchInput class="my-custom-search-input">
  <!-- Slot content or properties -->
</SearchInput>
```

---

## Supported Props

- `id`: Configures the `id` attribute or behavior.
- `name`: Configures the `name` attribute or behavior.
- `placeholder`: Configures the `placeholder` attribute or behavior.
- `shortcut`: Configures the `shortcut` attribute or behavior.
- `value`: Configures the `value` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

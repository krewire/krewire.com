---
title: "SearchInput"
description: "Search box with search icon and instant clear button."
date: "2026-11-09"
---

# SearchInput

Category: **Forms & Inputs**  
Component File: `components/SearchInput.kiw`  
Source: `github.com/krewire/ui/components/SearchInput.kiw`

Search box with search icon and instant clear button.

---

## 1. Copy & Own Installation

Copy the `SearchInput.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/SearchInput.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<SearchInput />` directly within your `.kiw` page or layout templates:

```html
<SearchInput class="my-custom-search-input">
  <!-- Slot content or properties -->
</SearchInput>
```

---

## 3. Supported Props & Attributes

- `id`: Configures the `id` attribute or behavior.
- `name`: Configures the `name` attribute or behavior.
- `placeholder`: Configures the `placeholder` attribute or behavior.
- `shortcut`: Configures the `shortcut` attribute or behavior.
- `value`: Configures the `value` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

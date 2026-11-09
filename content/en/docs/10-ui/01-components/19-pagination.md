---
title: "Pagination"
description: "Page navigation bar with previous, next, and numbered page links."
date: "2026-11-09"
---

# Pagination

Category: **Navigation & Structure**  
Component File: `components/Pagination.kiw`  
Source: `github.com/krewire/ui/components/Pagination.kiw`

Page navigation bar with previous, next, and numbered page links.

---

## 1. Copy & Own Installation

Copy the `Pagination.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Pagination.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Pagination />` directly within your `.kiw` page or layout templates:

```html
<Pagination class="my-custom-pagination">
  <!-- Slot content or properties -->
</Pagination>
```

---

## 3. Supported Props & Attributes

- `currentPage`: Configures the `currentPage` attribute or behavior.
- `hasNext`: Configures the `hasNext` attribute or behavior.
- `hasPrev`: Configures the `hasPrev` attribute or behavior.
- `nextUrl`: Configures the `nextUrl` attribute or behavior.
- `p`: Configures the `p` attribute or behavior.
- `pages`: Configures the `pages` attribute or behavior.
- `prevUrl`: Configures the `prevUrl` attribute or behavior.
- `totalPages`: Configures the `totalPages` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

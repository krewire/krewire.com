---
title: "Pagination"
description: "Page navigation bar with previous, next, and numbered page links."
date: "2026-11-09"
---

# Pagination

Category: **Navigation & Structure**

Page navigation bar with previous, next, and numbered page links.

---

## Basic Usage

Use `<Pagination />` directly within your `.kiw` page or layout templates:

```html
<Pagination class="my-custom-pagination">
  <!-- Slot content or properties -->
</Pagination>
```

---

## Supported Props

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

## Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

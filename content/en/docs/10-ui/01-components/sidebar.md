---
title: "Sidebar"
description: "Collapsible sidebar navigation for documentation and dashboard workspaces."
date: "2026-11-09"
---

# Sidebar

Category: **Navigation & Structure**  
Component File: `components/Sidebar.kiw`  
Source: `github.com/krewire/ui/components/Sidebar.kiw`

Collapsible sidebar navigation for documentation and dashboard workspaces.

---

## 1. Copy & Own Installation

Copy the `Sidebar.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Sidebar.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Sidebar />` directly within your `.kiw` page or layout templates:

```html
<Sidebar class="my-custom-sidebar">
  <!-- Slot content or properties -->
</Sidebar>
```

---

## 3. Supported Props & Attributes

- `head`: Configures the `head` attribute or behavior.
- `header`: Configures the `header` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `menu`: Configures the `menu` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

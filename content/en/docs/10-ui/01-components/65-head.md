---
title: "Head"
description: "Document head injector for standard meta tags, OpenGraph, and favicon assets."
date: "2026-11-09"
---

# Head

Category: **Navigation & Structure**  
Component File: `components/Head.kiw`  
Source: `github.com/krewire/ui/components/Head.kiw`

Document head injector for standard meta tags, OpenGraph, and favicon assets.

---

## 1. Copy & Own Installation

Copy the `Head.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Head.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Head />` directly within your `.kiw` page or layout templates:

```html
<Head class="my-custom-head">
  <!-- Slot content or properties -->
</Head>
```

---

## 3. Supported Props & Attributes

- `assetversion`: Configures the `assetversion` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

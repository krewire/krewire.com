---
title: "Container"
description: "Centered responsive page container maintaining standard max-width boundaries."
date: "2026-11-09"
---

# Container

Category: **Layout Primitives**  
Component File: `components/Container.kiw`  
Source: `github.com/krewire/ui/components/Container.kiw`

Centered responsive page container maintaining standard max-width boundaries.

---

## 1. Copy & Own Installation

Copy the `Container.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Container.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Container />` directly within your `.kiw` page or layout templates:

```html
<Container class="my-custom-container">
  <!-- Slot content or properties -->
</Container>
```

---

## 3. Supported Props & Attributes

- `id`: Configures the `id` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

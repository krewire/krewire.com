---
title: "Skeleton"
description: "Animated shimmer placeholder for asynchronous loading states."
date: "2026-11-09"
---

# Skeleton

Category: **Feedback & Status**  
Component File: `components/Skeleton.kiw`  
Source: `github.com/krewire/ui/components/Skeleton.kiw`

Animated shimmer placeholder for asynchronous loading states.

---

## 1. Copy & Own Installation

Copy the `Skeleton.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Skeleton.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Skeleton />` directly within your `.kiw` page or layout templates:

```html
<Skeleton class="my-custom-skeleton">
  <!-- Slot content or properties -->
</Skeleton>
```

---

## 3. Supported Props & Attributes

- `height`: Configures the `height` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.
- `width`: Configures the `width` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

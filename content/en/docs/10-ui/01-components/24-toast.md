---
title: "Toast"
description: "Transient notification toast banner with auto-dismiss timer."
date: "2026-11-09"
---

# Toast

Category: **Feedback & Status**  
Component File: `components/Toast.kiw`  
Source: `github.com/krewire/ui/components/Toast.kiw`

Transient notification toast banner with auto-dismiss timer.

---

## 1. Copy & Own Installation

Copy the `Toast.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Toast.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Toast />` directly within your `.kiw` page or layout templates:

```html
<Toast class="my-custom-toast">
  <!-- Slot content or properties -->
</Toast>
```

---

## 3. Supported Props & Attributes

- `dismissible`: Configures the `dismissible` attribute or behavior.
- `icon`: Configures the `icon` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `message`: Configures the `message` attribute or behavior.
- `position`: Configures the `position` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

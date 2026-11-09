---
title: "CodeWindow"
description: "macOS-style code preview window with traffic light buttons and copy action."
date: "2026-11-09"
---

# CodeWindow

Category: **Data Display & Content**  
Component File: `components/CodeWindow.kiw`  
Source: `github.com/krewire/ui/components/CodeWindow.kiw`

macOS-style code preview window with traffic light buttons and copy action.

---

## 1. Copy & Own Installation

Copy the `CodeWindow.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/CodeWindow.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<CodeWindow />` directly within your `.kiw` page or layout templates:

```html
<CodeWindow class="my-custom-code-window">
  <!-- Slot content or properties -->
</CodeWindow>
```

---

## 3. Supported Props & Attributes

- `actions`: Configures the `actions` attribute or behavior.
- `code`: Configures the `code` attribute or behavior.
- `context`: Configures the `context` attribute or behavior.
- `get`: Configures the `get` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `json`: Configures the `json` attribute or behavior.
- `map`: Configures the `map` attribute or behavior.
- `new`: Configures the `new` attribute or behavior.
- `run`: Configures the `run` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

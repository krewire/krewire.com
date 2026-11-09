---
title: "Form"
description: "Accessible form wrapper with method spoofing and submission states."
date: "2026-11-09"
---

# Form

Category: **Forms & Inputs**  
Component File: `components/Form.kiw`  
Source: `github.com/krewire/ui/components/Form.kiw`

Accessible form wrapper with method spoofing and submission states.

---

## 1. Copy & Own Installation

Copy the `Form.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Form.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Form />` directly within your `.kiw` page or layout templates:

```html
<Form class="my-custom-form">
  <!-- Slot content or properties -->
</Form>
```

---

## 3. Supported Props & Attributes

- `action`: Configures the `action` attribute or behavior.
- `enctype`: Configures the `enctype` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `method`: Configures the `method` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

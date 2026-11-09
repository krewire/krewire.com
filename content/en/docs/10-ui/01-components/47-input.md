---
title: "Input"
description: "Standard text input with floating label, error states, and helper text."
date: "2026-11-09"
---

# Input

Category: **Forms & Inputs**  
Component File: `components/Input.kiw`  
Source: `github.com/krewire/ui/components/Input.kiw`

Standard text input with floating label, error states, and helper text.

---

## Live Preview

<div data-component-preview="Input" data-slug="input"></div>

The frames above load the real `Input.kiw` from Krewire UI, prebuilt by `kiw build` for every variant, in the light and dark themes side by side. Use the toggles to switch variants, then copy the component source or the usage snippet. **Open in playground** shows it next to the other core components.

---

## 1. Copy & Own Installation

Copy the `Input.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Input.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Input />` directly within your `.kiw` page or layout templates:

```html
<Input class="my-custom-input">
  <!-- Slot content or properties -->
</Input>
```

---

## 3. Supported Props & Attributes

- `disabled`: Configures the `disabled` attribute or behavior.
- `error`: Configures the `error` attribute or behavior.
- `help`: Configures the `help` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `name`: Configures the `name` attribute or behavior.
- `placeholder`: Configures the `placeholder` attribute or behavior.
- `required`: Configures the `required` attribute or behavior.
- `type`: Configures the `type` attribute or behavior.
- `value`: Configures the `value` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

---
title: "Timeline"
description: "Chronological vertical activity and event timeline."
date: "2026-11-09"
---

# Timeline

Category: **Data Display & Content**  
Component File: `components/Timeline.kiw`  
Source: `github.com/krewire/ui/components/Timeline.kiw`

Chronological vertical activity and event timeline.

---

## 1. Copy & Own Installation

Copy the `Timeline.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Timeline.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Timeline />` directly within your `.kiw` page or layout templates:

```html
<Timeline class="my-custom-timeline">
  <!-- Slot content or properties -->
</Timeline>
```

---

## 3. Supported Props & Attributes

- `item`: Configures the `item` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.

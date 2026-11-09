---
title: "UI Components Catalog"
description: "Comprehensive reference catalog for all 66 Krewire UI pop-brutalist Single-File Components (.kiw)."
date: "2026-11-09"
---

# Krewire UI Components Catalog

Import path: `github.com/krewire/ui`  
Design System Stylesheet: `/styles/krewire-ui.css`

Welcome to the **Krewire UI (`kui`) Components Catalog**.

Krewire UI provides **66 accessible, zero-dependency, pop-brutalist Single-File Components (`.kiw`)** built natively for the Krewire Go ecosystem.

All components follow the **shadcn/ui "Copy & Own"** philosophy: you copy `.kiw` files directly into your project's `./components/` directory, giving you complete ownership without opaque runtime dependencies or upstream breaking changes.

---

## Component Categories

### Navigation & Structure

- [**`<Navbar />`**](/docs/ui/components/navbar) — Main application header with brand logo, nav links, and drawer trigger.
- [**`<NavMenu />`**](/docs/ui/components/nav-menu) — Horizontal navigation menu with interactive dropdown menus.
- [**`<Sidebar />`**](/docs/ui/components/sidebar) — Collapsible sidebar navigation for documentation and dashboard workspaces.
- [**`<SideMenu />`**](/docs/ui/components/side-menu) — Vertical navigational list for administrative sidebars.
- [**`<Breadcrumb />`**](/docs/ui/components/breadcrumb) — Hierarchical navigation trail indicating current page location.
- [**`<BreadcrumbItem />`**](/docs/ui/components/breadcrumb-item) — Individual segment within a breadcrumb navigation trail.
- [**`<Tabs />`**](/docs/ui/components/tabs) — Accessible tabbed content switcher supporting persistent active states.
- [**`<Pagination />`**](/docs/ui/components/pagination) — Page navigation bar with previous, next, and numbered page links.
- [**`<Drawer />`**](/docs/ui/components/drawer) — Smooth sliding drawer panel for mobile navigation and overlays.
- [**`<Hamburger />`**](/docs/ui/components/hamburger) — Animated 3-line hamburger menu toggle button for mobile drawers.
- [**`<Brand />`**](/docs/ui/components/brand) — Ecosystem brand mark with logo icon and version badge.
- [**`<Head />`**](/docs/ui/components/head) — Document head injector for standard meta tags, OpenGraph, and favicon assets.
- [**`<Header />`**](/docs/ui/components/header) — Top sectional page banner header with title and subtitle.
- [**`<Footer />`**](/docs/ui/components/footer) — Comprehensive multi-column footer with brand, navigation links, and copyright.

### Actions & Buttons

- [**`<Button />`**](/docs/ui/components/button) — Core interactive button supporting primary, secondary, accent, outline, and ghost variants.
- [**`<ButtonGroup />`**](/docs/ui/components/button-group) — Grouped button segment with unified borders and active states.
- [**`<Dropdown />`**](/docs/ui/components/dropdown) — Zero-dependency dropdown menu built on native details and summary elements.
- [**`<ActionMenu />`**](/docs/ui/components/action-menu) — Compact contextual dropdown action menu for table rows and cards.
- [**`<Link />`**](/docs/ui/components/link) — Accessible hyperlink with underline transitions and target security.
- [**`<Toggle />`**](/docs/ui/components/toggle) — Binary toggle switch for instant setting changes.
- [**`<ThemeSwitch />`**](/docs/ui/components/theme-switch) — Dark and light mode toggle button with localStorage persistence.
- [**`<LangSwitch />`**](/docs/ui/components/lang-switch) — Language selector switch for multi-lingual internationalized applications.

### Feedback & Status

- [**`<Alert />`**](/docs/ui/components/alert) — Contextual feedback banner supporting info, success, warning, and error states.
- [**`<Banner />`**](/docs/ui/components/banner) — Prominent top-level announcement bar with dismiss trigger.
- [**`<Toast />`**](/docs/ui/components/toast) — Transient notification toast banner with auto-dismiss timer.
- [**`<Badge />`**](/docs/ui/components/badge) — Compact pill badge with color variants for tags, counts, and status indicators.
- [**`<Tag />`**](/docs/ui/components/tag) — Removable or clickable keyword tag with color variants.
- [**`<Progress />`**](/docs/ui/components/progress) — Determinate and indeterminate progress bar indicator.
- [**`<Spinner />`**](/docs/ui/components/spinner) — Accessible animated SVG loader spinner for pending operations.
- [**`<Skeleton />`**](/docs/ui/components/skeleton) — Animated shimmer placeholder for asynchronous loading states.
- [**`<Empty />`**](/docs/ui/components/empty) — Empty-state placeholder with icon, message, and call-to-action button.
- [**`<Callout />`**](/docs/ui/components/callout) — Bordered callout container for documentation tips, warnings, and notes.

### Data Display & Content

- [**`<Card />`**](/docs/ui/components/card) — Elevated card surface with header, media, body, and footer slots.
- [**`<Stat />`**](/docs/ui/components/stat) — Key performance indicator card with numeric metrics and trend indicators.
- [**`<Table />`**](/docs/ui/components/table) — Responsive data table with striped, bordered, and hover highlight states.
- [**`<Accordion />`**](/docs/ui/components/accordion) — Expandable and collapsible content panels with keyboard accessibility.
- [**`<Collapse />`**](/docs/ui/components/collapse) — Collapsible single-panel container for toggling visibility.
- [**`<Timeline />`**](/docs/ui/components/timeline) — Chronological vertical activity and event timeline.
- [**`<CodeWindow />`**](/docs/ui/components/code-window) — macOS-style code preview window with traffic light buttons and copy action.
- [**`<Terminal />`**](/docs/ui/components/terminal) — Single-line interactive terminal preview with copy command button.
- [**`<Code />`**](/docs/ui/components/code) — Inline monospace code snippet with high contrast styling.
- [**`<Avatar />`**](/docs/ui/components/avatar) — User profile avatar displaying image with automatic fallback initials.
- [**`<Kbd />`**](/docs/ui/components/kbd) — Keyboard key visual indicator for shortcuts and command hotkeys.
- [**`<Divider />`**](/docs/ui/components/divider) — Horizontal or vertical separator line with optional text label.
- [**`<Heading />`**](/docs/ui/components/heading) — Semantic heading typography wrapper supporting h1 through h6.
- [**`<Text />`**](/docs/ui/components/text) — Typography wrapper for body paragraphs, lead text, and muted captions.
- [**`<List />`**](/docs/ui/components/list) — Styled unordered and ordered list container.
- [**`<Ecosystem />`**](/docs/ui/components/ecosystem) — Interactive grid component displaying Krewire's 10-repository architecture.

### Forms & Inputs

- [**`<Form />`**](/docs/ui/components/form) — Accessible form wrapper with method spoofing and submission states.
- [**`<Input />`**](/docs/ui/components/input) — Standard text input with floating label, error states, and helper text.
- [**`<SearchInput />`**](/docs/ui/components/search-input) — Search box with search icon and instant clear button.
- [**`<Select />`**](/docs/ui/components/select) — Native styled dropdown select input.
- [**`<Textarea />`**](/docs/ui/components/textarea) — Multi-line text input field with auto-resizing support.
- [**`<Checkbox />`**](/docs/ui/components/checkbox) — High-contrast accessible checkbox input with custom pop-border styling.
- [**`<Radio />`**](/docs/ui/components/radio) — Single-choice radio button with pop-brutalist circle styling.
- [**`<Slider />`**](/docs/ui/components/slider) — Range slider input with numeric bubble value display.

### Overlays & Dialogs

- [**`<Modal />`**](/docs/ui/components/modal) — Accessible modal dialog built on the native HTML dialog element.
- [**`<Popover />`**](/docs/ui/components/popover) — Floating contextual popover container anchored to target elements.
- [**`<Tooltip />`**](/docs/ui/components/tooltip) — Accessible hover and focus informative tooltip.

### Layout Primitives

- [**`<Container />`**](/docs/ui/components/container) — Centered responsive page container maintaining standard max-width boundaries.
- [**`<Box />`**](/docs/ui/components/box) — Versatile layout box supporting border, padding, and pop-brutalist shadows.
- [**`<Flex />`**](/docs/ui/components/flex) — Flexible layout container with configurable direction, align, and justify.
- [**`<Grid />`**](/docs/ui/components/grid) — CSS Grid container with configurable columns and gap tokens.
- [**`<HStack />`**](/docs/ui/components/h-stack) — Horizontal flex container with uniform spacing between children.
- [**`<VStack />`**](/docs/ui/components/v-stack) — Vertical flex container with uniform spacing between children.
- [**`<Stack />`**](/docs/ui/components/stack) — Directional layout container with configurable gap spacing.
- [**`<Section />`**](/docs/ui/components/section) — Full-width semantic section wrapper with vertical padding tokens.

---

## How to Install Any Component

To add any component to your project, use the standard copy command:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/<ComponentName>.kiw ./components/
```

And link the canonical stylesheet in your layout `<head>`:

```html
<link rel="stylesheet" href="/styles/krewire-ui.css">
```

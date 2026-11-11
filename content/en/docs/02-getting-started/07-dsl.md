---
title: "DSL (.kiw)"
description: "In-depth guide to the .kiw component templating language: YAML frontmatter, Go html/template expressions, scoped CSS, client scripts, and embedded Markdown."
date: "2026-09-29"
---

# DSL (.kiw)

The **`.kiw` Component DSL** (`pkg/kiw`) is Krewire's unified, file-based component templating language. It brings modern Single-File Component (SFC) ergonomics—reminiscent of Astro, Svelte, or Vue—directly to the Go ecosystem without requiring Node.js, `npm`, webpack, or JavaScript build fatigue.

Every `.kiw` file encapsulates its metadata, markup, scoped styles, and client interactions in a single, readable file that compiles down to high-performance Go template trees and static assets in milliseconds.

---

## 1. Anatomy of a `.kiw` File

A canonical `.kiw` component consists of four distinct, optional sections:

```html
---
# 1. YAML Frontmatter
title: "User Profile"
layout: Base
roles: ["admin", "editor"]
---

<!-- 2. Template Body (Go html/template) -->
<article class="profile-card">
  <h2>{{.Title}}</h2>
  <p>Status: <span class="badge">Active</span></p>

  <!-- 3. Embedded Markdown (Optional) -->
  <markdown>
  ### Bio
  Fullstack developer passionate about **Go** and systems engineering.
  </markdown>
</article>

<!-- 4. Scoped CSS (Automatically Scoped) -->
<style>
.profile-card {
  border: var(--pop-border, 2px solid #0B1F3B);
  border-radius: 12px;
  padding: 1.5rem;
  background: var(--bg-2, #f0ede5);
}
h2 {
  color: var(--primary, #39D353);
  margin-top: 0;
}
.badge {
  font-weight: bold;
  text-transform: uppercase;
}
</style>

<!-- 5. Client Script (Automatically Extracted) -->
<script hydrate="idle">
console.log("Profile card mounted and active.");
</script>
```

---

## 2. YAML Frontmatter

The top of any page or layout file can declare frontmatter delimited by triple dashes (`---`). Frontmatter is parsed into typed Go data structures via `gopkg.in/yaml.v3`:

```yaml
---
title: "Documentation Guide"
description: "Learn how to build documentation portals."
layout: Base
date: "2026-09-29"
draft: false
tags:
  - golang
  - architecture
author:
  name: "Krewire Team"
  email: "team@krewire.com"
---
```

### Special Frontmatter Keys:

| Key | Purpose | Used In |
| :--- | :--- | :--- |
| `title` | Page title injected into the layout's `<title>` tag and OpenGraph headers. | `pages/` |
| `layout` | Name of the layout wrapper inside `layouts/` (e.g. `layout: Base` wraps with `layouts/Base.kiw`). | `pages/` |
| `description` | Meta description for search engines and social cards. | `pages/` |
| `draft` | When set to `true`, the page is excluded from production builds. | `pages/` |
| `date` | ISO 8601 publication or modification timestamp. | `pages/`, `content/` |

All custom keys declared in frontmatter are exposed directly inside template expressions via the root context: `{{.Title}}`, `{{.Tags}}`, `{{.Author.Name}}`.

---

## 3. Template Body & Go Expressions

The body of a `.kiw` file uses Go's standard `html/template` syntax. It provides context-aware HTML escaping to prevent Cross-Site Scripting (XSS) by default.

### Common Expressions:

#### Data Interpolation
```html
<h1>{{.Title}}</h1>
<p>Author: {{.Author.Name}}</p>
```

#### Conditional Rendering
```html
{{if .IsAdmin}}
  <a href="/admin/dashboard" class="btn btn-danger">Admin Panel</a>
{{else if .IsSubscriber}}
  <span class="badge">Subscriber</span>
{{else}}
  <a href="/signup" class="btn">Join Now</a>
{{end}}
```

#### Looping over Slices and Maps
```html
<ul class="tag-list">
  {{range .Tags}}
    <li class="tag-item">{{.}}</li>
  {{else}}
    <li class="empty">No tags available</li>
  {{end}}
</ul>
```

#### Class Composition Helper (`class`)
To collapse verbose conditional modifier blocks inside class attributes without bypassing contextual HTML escaping:

```html
<!-- Composes base token + conditional modifiers + trailing class props -->
<a href="{href}" class="{{class "btn" .Variant "primary" .Size "md" .Class}}">
  {body}
</a>
```

The `class` helper automatically handles prefix resolution (e.g. `btn-primary`, `btn-md`), falls back to defaults when props are empty or omitted, collapses whitespace, and returns a plain string so standard `html/template` attribute-context escaping is strictly preserved.

#### Layout Content Injection
In layout components (`layouts/*.kiw`), the page's rendered body is injected using the `{{.Content}}` directive:

```html
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>{{.Title}}</title>
</head>
<body>
  <header><Header /></header>
  <main id="main">
    {{.Content}}
  </main>
  <footer><Footer /></footer>
</body>
</html>
```

No `<link>` or `<script>` tag is needed for site CSS/JS: every file in
`public/assets/` ending in `.css` or `.js` is injected automatically —
stylesheets into `<head>`, scripts into `<head>` — with a `?v=<version>`
cache-busting query derived from the project version. Writing a tag by hand
still works; the pipeline never adds a second copy of an asset you linked
yourself. See `auto_assets:` in [`krewire.yaml`](configuration) to turn
injection off or exclude specific files.

---

## 4. Reusable Component Composition

Components located in the `components/` directory (or installed into `./.krewire/pkg/`, or embedded from Krewire UI) can be instantiated dynamically inside any page, layout, or component using custom component tags:

```html
<!-- Tag-based invocation with attributes as props -->
<Button Variant="primary" Href="/docs/getting-started">Get Started</Button>
<Alert Type="warning" Message="System update in progress" />
```

The equivalent Go template helper is `{{component "ComponentName" .}}`; both forms produce the exact same component tree.

### 4.1 Compound Components (Dotted Notation)

Components can be organized into subdirectories and invoked using dotted compound notation:

```html
<Card>
  <Card.Header title="Deployment Status" />
  <Card.Body>
    <p>All microservices are running nominally.</p>
  </Card.Body>
  <Card.Footer>
    <Button Variant="ghost">Dismiss</Button>
  </Card.Footer>
</Card>
```

Scaffold compound components directly with the CLI:
```bash
kiw make component Card.Header     # generates components/Card/Header.kiw
kiw make component Card.Body       # generates components/Card/Body.kiw
kiw make component Card.Footer     # generates components/Card/Footer.kiw
```

### 4.2 Namespaced & Vendor Components

Installed UI packages from `./.krewire/pkg/` or embedded design systems can be invoked with colon namespace syntax:

```html
<ui:Button variant="primary">Launch</ui:Button>
<acme:MetricCard label="Active Requests" value=1420 />
```

### 4.3 Slot Outlets & Fallback Content

Child content passed inside a paired tag populates the default slot (`{body}`). Components can declare named slot outlets with default fallback content that renders automatically when no content is provided by the caller:

```html
<!-- components/Card/Header.kiw -->
<div class="card-header">
  <slot name="title">
    <h3>Default Header Title</h3>
  </slot>
  <slot name="actions"></slot>
</div>
```

Callers supply slot content using `<template slot="...">` or `<slot:...>`:

```html
<Card.Header>
  <template slot="title">Custom Cluster Metrics</template>
  <template slot="actions">
    <Button size="sm">Refresh</Button>
  </template>
</Card.Header>
```

### 4.4 Inline Snippets & Local Macros

Reusable template snippets can be authored and rendered in the same `.kiw` file without creating external files:

```html
{#snippet userBadge(name, role)}
  <span class="badge {role}">{name}</span>
{/snippet}

<div class="team-roster">
  {@render userBadge("Alice", "admin")}
  {@render userBadge("Bob", "contributor")}
</div>
```

Snippets desugar cleanly into scoped Go template definitions and invocations.

### 4.5 Declarative Component Imports

Components can declare imports in `<script>` blocks or frontmatter:

```html
<script>
  import Header from './Header.kiw';
  import { Badge } from '@krewire/ui';
</script>

<Header />
<Badge variant="success">Online</Badge>
```

### 4.6 Standalone Component Packages (`.kpkg`)

Reusable component libraries can be distributed and installed as standalone packages with zero npm dependencies:
- **Manifest (`kiw.pkg.json`)**: Declares package metadata, exported component paths, and stylesheets.
- **Archive Format (`.kpkg`)**: Gzipped tarball with SHA-256 integrity verification.
- **Install Directory (`./.krewire/pkg/`)**: Local project packages repository.

```bash
kiw pkg init @acme/ui              # Scaffolds package directory with kiw.pkg.json
kiw pkg pack ./acme-ui             # Packs into acme-ui-0.1.0.kpkg archive
kiw pkg add ./acme-ui-0.1.0.kpkg   # Installs into ./.krewire/pkg/acme-ui
kiw pkg list                       # Lists installed packages and exported components
```

Stylesheets declared in `kiw.pkg.json` (`styles: "./styles/theme.css"`) are automatically compiled and injected into the site asset pipeline.

### Example: Component with Scoped Props

Define `components/Alert.kiw`:

```html
<div class="alert alert-{{.Type}}">
  <span class="icon">ℹ</span>
  <div class="message">{{.Message}}</div>
</div>

<style>
.alert {
  display: flex;
  gap: 0.75rem;
  padding: 1rem;
  border-radius: 8px;
}
.alert-info { background: #e0f2fe; color: #0369a1; }
.alert-warning { background: #fef3c7; color: #92400e; }
</style>
```

Instantiate `Alert` inside `pages/index.kiw`:

```html
<Alert Type="warning" Message="Maintenance scheduled tonight at 23:00 UTC" />
```

### Interactive Islands & Client Hydration

Krewire implements an **Island Architecture** (`pkg/runtime/mount` and `pkg/runtime/boot`) allowing selective client-side hydration of interactive components while keeping the rest of the page 100% zero-JavaScript static HTML.

You declare hydration behavior using `client:*` directives directly on component tags:

```html
<!-- Hydrate immediately on page load -->
<Counter client:load initial=0 />

<!-- Hydrate once the browser main thread is idle (requestIdleCallback) -->
<SearchDialog client:idle placeholder="Search docs..." />

<!-- Hydrate only when scrolled into the viewport (IntersectionObserver) -->
<CommentThread client:visible threadId="kw-99" />
```

**How It Works:**
1. **Server Rendering (SSR/SSG)**: During build, Krewire renders the component to HTML and wraps it in a lightweight marker container (`data-kiw-mount="Counter" data-kiw-hydrate="load" data-kiw-props='{"initial":0}'`).
2. **Instant Preview**: The browser paints the full content immediately with zero layout shift (CLS: 0).
3. **Progressive Hydration**: The Krewire WebAssembly runtime (`pkg/runtime`) scans for mount markers, seeds the VDOM fiber with the existing DOM, and attaches reactive events without re-rendering server text.

---

## 5. Automatic Scoped CSS

Styles defined inside `<style>` blocks are **scoped to the component automatically** at build time. You never have to worry about CSS class name collisions.

### How Scoping Works:

1. When parsing a component (e.g. `components/Card.kiw`), Krewire injects a unique attribute into the component's root elements:
   ```html
   <div class="card" data-kiw-component="Card">...</div>
   ```
2. The CSS compiler (`pkg/web/ssg`) rewrites your selectors to bind strictly to that component's scope:
   ```css
   /* Your authored CSS */
   .card { padding: 20px; }
   h3 { font-size: 1.25rem; }

   /* Compiled CSS output in assets/style.css */
   [data-kiw-component="Card"] .card, [data-kiw-component="Card"].card { padding: 20px; }
   [data-kiw-component="Card"] h3, [data-kiw-component="Card"]h3 { font-size: 1.25rem; }
   ```
3. All component styles are bundled and minified into a single `assets/style.css` file, eliminating render-blocking CSS requests.

---

## 6. Client Scripts & Hydration Directives

Scripts inside `<script>` tags are extracted from the HTML and compiled into static JavaScript asset files. This guarantees clean Separation of Concerns and strict Content Security Policy (CSP) compliance without inline script execution risks.

```html
<script hydrate="visible">
  const button = document.querySelector(".btn-interactive");
  button.addEventListener("click", () => {
    console.log("Button clicked!");
  });
</script>
```

### Hydration Tiers (`hydrate` attribute):

| Hydration Directive | Execution Timing | Ideal Use Case |
| :--- | :--- | :--- |
| `hydrate="load"` *(default)* | Executes immediately when the DOM loads (`DOMContentLoaded`). | Navigation menus, critical UI toggles, theme switchers. |
| `hydrate="idle"` | Executes when the main browser thread is idle (`requestIdleCallback`). | Analytics, non-critical telemetry, background widgets. |
| `hydrate="visible"` | Executes only when the element enters the viewport (`IntersectionObserver`). | Image carousels, lazy-loaded interactive charts, comments. |

---

## 7. Embedded Markdown

For content-heavy sections, you can write native Markdown directly within your `.kiw` component using the `<markdown>` tag:

```html
<section class="changelog">
  <h2>Release Notes</h2>
  <markdown>
  ### Version 0.1.0 (Initial Release)
  - **Libraries:** Core modular library architecture.
  - **Kiw CLI:** Scaffolding, dev server, and build pipeline.
  - **Zero Bloat:** 100% Go standard library, zero npm dependencies.
  </markdown>
</section>
```

The Krewire compiler automatically converts the Markdown block to clean, sanitized HTML at build time using `pkg/kiw` and `pkg/markdown`.

---

## 8. Built-in Helpers & Internationalization

Krewire provides built-in template helper functions:

| Helper | Syntax | Description |
| :--- | :--- | :--- |
| Component tags | `<Name />` | Renders a child component using JSX-like syntax. |
| `component` | `{{component "Name" .}}` | Compatible template-helper form for rendering a child component with context or dict props. |
| `t` / `translate` | `{{t "home.welcome"}}` | Translates a key using the active language bundle (`lang/*.json`). |
| `tLocale` | `{{tLocale "fr" "home.welcome"}}` | Translates a key for a specific locale. |
| `locales` | `{{locales}}` | Returns a list of all configured language codes. |

---

## Next Steps

Now that you have mastered the `.kiw` DSL and Krewire's component system:

- Explore the [**1.3 Krewire Workloads →**](/overview/krewire-workloads) matrix to see how `.kiw` powers static sites, books, and web applications.
- Return to the [**2. Getting Started Index →**](/getting-started).

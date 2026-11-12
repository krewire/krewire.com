---
title: "Data Flow: SSG vs Runtime"
description: "How data reaches .kiw views in Krewire: build-time evaluation for site and book, request-time rendering for app, Go handlers injecting structs and query results, and client scripts talking to JSON endpoints."
date: "2026-11-09"
---

# Data Flow: SSG vs Runtime

<nav class="track-crumbs" aria-label="Architecture Deep-Dive track">
<span class="track-crumbs-label">Deep-Dive track</span>
<ol>
<li><a href="/architecture">Overview</a></li>
<li><a href="/architecture/layered-pyramid">1 · Pyramid</a></li>
<li><a href="/architecture/ecosystem-dag">2 · Ecosystem DAG</a></li>
<li><a href="/architecture/scope-levels">3 · Scope levels</a></li>
<li aria-current="step"><span>4 · Data flow</span></li>
</ol>
</nav>

A `.kiw` file never fetches data itself. It is a template that receives a **data value** and turns it into HTML. What differs between workloads is *when* that value is computed and *who* computes it:

- In a **`site` or `book`** the data is evaluated **once, at build time**, by `kiw build` (`pkg/web/ssg`, `pkg/book`). The output is plain files; no server runs your code.
- In an **`app`** the data is computed **per request** by Go functions you register on a `web.App` (`pkg/web`, assembled with `pkg/app`). The same renderer turns it into HTML on every request.
- In both, a **client script** can talk to JSON endpoints after the page loads — in an `app` those endpoints are your own Go handlers.

---

## 1. The Request-Response Lifecycle

```text
BUILD TIME — site, book (runs once, on your machine or in CI)

  krewire.yaml ─┐
  content/*.md ─┤
  lang/*.json  ─┼──► kiw build ──► page data = yaml values + frontmatter
  pages/*.kiw  ─┤    (ssg.LoadFromDir)        + Collections + i18n
  components/  ─┤                                   │
  layouts/     ─┘                                   ▼
                                     ssg render: page .kiw + layout + kui
                                                    │
                                                    ▼
                                  .krewire/build/*.html ──► any static host


REQUEST TIME — app (runs on every request, inside your binary)

  Browser ── GET / ──► web.App router
                          │  middleware: Recover → AccessLog → CSRF → …
                          ▼
                  PageSpec.DataFunc(r) ──► your Go code ──► SQL, services
                          │  returns a struct / map / rows
                          ▼
                  ssg render: Root .kiw view + layout + kui   (same engine)
                          │  EmitProps → data-props JSON on the root element
                          ▼
  Browser ◄── HTML ───────┘
     │
     └── Scripts: fetch /api/… ──► Router().Get / Post ──► web.JSON(…) ──► Browser
```

The renderer in the middle is the same in both columns. `web.App` builds an `ssg.Site` from its pages, components, and layouts, so a page with fixed data renders byte-for-byte the same whether it is served live or exported.

| | Build-time SSG (`site`, `book`) | Request-time SSR (`app`) |
| :--- | :--- | :--- |
| Who computes the data | `kiw build` | Your Go code on every request (`PageSpec.DataFunc`) or once at startup (`PageSpec.Data`) |
| Data sources | `krewire.yaml`, page frontmatter, `content/<collection>/*.md`, `lang/` dictionaries | Anything Go can reach: databases (`pkg/orm`), services, the request (query, cookies, session, CSRF token) |
| Output | Static `.html` files in `.krewire/build/` | An HTML response per request, from one Go binary |
| Server needed | None — any static host | The app binary |
| Per-user content | Only client-side, after load | Yes — computed per request |
| Freshness | Until the next build | Every request |
| Interactivity | Client scripts calling external or separate APIs | Client scripts calling the same binary's JSON routes |

---

## 2. Build Time: How a Site Gets Its Data

`kiw build` (via `ssg.LoadFromDir`) merges several sources into each page's data:

| Source | Available in templates as |
| :--- | :--- |
| `krewire.yaml` → `project.version` | `.Version` (`{version}`) |
| Page frontmatter (`title`, any custom key) | `.Title`, `.Description`, and each key **with its original spelling**, e.g. `{{.tagline}}` |
| `content/<name>/*.md` collections | `.Collections.<name>` — a list of `{Title, Date, Permalink, Data, Content, Draft, Slug, Description}` |
| `lang/<locale>/*.json` | the `t` / `translate` template functions |
| `pages/<dir>/[slug].kiw` | one page per collection entry, with `.Title`, `.Content`, and `.Page` |

A listing page plus one page per post — with `content/posts/hello.md` carrying `title`, `slug`, and `date` frontmatter:

```html
---
title: Posts
layout: Base
tagline: "Notes from the build pipeline"
---
<h1>{{.Title}}</h1>
<p>{{.tagline}} · site version {version}</p>
<ul>
  {{range .Collections.posts}}
  <li><a href="/posts/{{.Slug}}">{{.Title}}</a> — {{.Date.Format "2 Jan 2006"}}</li>
  {{end}}
</ul>
```

```html
---
layout: Base
---
<article>
  <h1>{{.Title}}</h1>
  <div>{{.Content}}</div>
</article>
```

Saved as `pages/posts.kiw` and `pages/posts/[slug].kiw`, `kiw build --target site` writes `posts.html`, `posts/hello.html`, and `posts/second.html`:

```html
<p>Notes from the build pipeline · site version 0.1.0</p>
<ul>
  <li><a href="/posts/second">Second post</a> — 5 Nov 2026</li>
  <li><a href="/posts/hello">Hello, Krewire</a> — 1 Nov 2026</li>
</ul>
```

Three details that save debugging time:

- The `{name}` shorthand capitalizes the first letter (`{tagline}` → `.Tagline`). Data loaded from YAML keeps its own spelling, so use the explicit `{{.tagline}}` and `{{range .Collections.posts}}` forms for frontmatter keys and collection names.
- A site with a `content/` folder is also picked up by the book pipeline when you run plain `kiw build`. Use `kiw build --target site` when `content/` only holds collections.
- Nothing on this side runs at request time. To show data that changes between builds, either rebuild (for example on a schedule in CI) or fetch it from the browser (section 5).

---

## 3. Request Time: Go Handlers Inject Data

In an `app`, each page is a `web.PageSpec`. The page's `Root` names the `.kiw` view (registered as a component), and the data comes from one of two fields:

| Field | Runs | Exported by `App.Export` | Use for |
| :--- | :--- | :--- | :--- |
| `Data any` | never — fixed at startup | **yes**, to static HTML | marketing pages, docs, anything identical for every visitor |
| `DataFunc func(*http.Request) (any, error)` | on every `GET` | no (dynamic pages are skipped) | per-user data, database queries, anything that depends on the request |

The following program — a complete, compiled example — shows both, plus a JSON API used by a client script in section 5. Its two views live in `views/`:

```html
<section>
  <h1>About v{version}</h1>
  <ul>
    {#each features as f}<li>{f}</li>{/each}
  </ul>
</section>
```

```html
<section id="notes">
  <h1>Notes <small>({count})</small></h1>
  <ul id="note-list">
    {#each notes as n}<li>{n.Text}</li>{/each}
  </ul>
  <form id="note-form">
    <input name="text" placeholder="New note" required>
    <button type="submit">Add</button>
  </form>
</section>
```

And the Go side:

```go
// Command notes demonstrates the three ways data reaches a page in a
// Krewire app: fixed data (exportable), per-request data, and JSON fetched
// by a client script.
package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/krewire/krewire/pkg/kiw"
	"github.com/krewire/krewire/pkg/validation"
	"github.com/krewire/krewire/pkg/web"
	"github.com/krewire/krewire/pkg/web/ssg"
)

//go:embed views/*.kiw static
var files embed.FS

// Note is the row type shared by the page, the API, and the script.
type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Store stands in for a database (see the fullstack tutorial for pkg/orm).
type Store struct {
	mu    sync.Mutex
	notes []Note
}

func (s *Store) List() []Note {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Note(nil), s.notes...)
}

func (s *Store) Add(text string) Note {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := Note{ID: len(s.notes) + 1, Text: text}
	s.notes = append(s.notes, n)
	return n
}

// AboutPage is fixed at startup, so the page can be exported to HTML.
type AboutPage struct {
	Version  string
	Features []string
}

// NotesPage is computed per request.
type NotesPage struct {
	Notes []Note
	Count int
	Token string
}

const shell = `<!doctype html><html><head><title>{{.Title}}</title></head>
<body><main>{{.Content}}</main></body></html>`

// NewApp wires the pages, the API, and the static script.
func NewApp(st *Store) (*web.App, error) {
	a := web.NewApp()
	a.Use(web.CSRF())
	a.Layout(ssg.Layout{Name: "Shell", Body: shell})

	views, _ := fs.Glob(files, "views/*.kiw")
	for _, p := range views {
		src, _ := files.ReadFile(p)
		mod, err := kiw.ParseKiw(string(src))
		if err != nil {
			return nil, err
		}
		name := strings.TrimSuffix(strings.TrimPrefix(p, "views/"), ".kiw")
		a.Component(ssg.Component{Name: name, Body: mod.Body})
	}

	// 1. Fixed data: rendered the same for everyone, exportable.
	a.Page(web.PageSpec{
		Path: "/about", Title: "About", Layout: "Shell", Root: "About",
		Data: AboutPage{Version: "1.0.0", Features: []string{"SSR", "JSON API", "Static export"}},
	})

	// 2. Per-request data: a handler-style function runs on every GET.
	a.Page(web.PageSpec{
		Path: "/", Title: "Notes", Layout: "Shell", Root: "Notes",
		DataFunc: func(r *http.Request) (any, error) {
			list := st.List()
			return NotesPage{Notes: list, Count: len(list), Token: web.CSRFFrom(r.Context())}, nil
		},
		EmitProps: true,                        // data-props='{"Notes":[…],…}' on the root element
		Scripts:   []string{"/static/notes.js"}, // client script that talks to the API
	})

	// 3. JSON endpoints for the client script.
	a.Router().Get("/api/notes", func(w http.ResponseWriter, r *http.Request, _ web.Params) {
		web.JSON(w, http.StatusOK, st.List())
	})
	a.Router().Post("/api/notes", func(w http.ResponseWriter, r *http.Request, _ web.Params) {
		var in struct {
			Text string `json:"text" validate:"required,max=200"`
		}
		if err := web.ReadJSON(r, &in); err != nil {
			web.Error(w, err)
			return
		}
		if err := validation.Struct(&in); err != nil {
			web.Error(w, web.BadRequest(err.Error()))
			return
		}
		web.JSON(w, http.StatusCreated, st.Add(in.Text))
	})

	static, _ := fs.Sub(files, "static")
	a.Static("/static", static)
	return a, nil
}

func main() {
	a, err := NewApp(&Store{})
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 1 && os.Args[1] == "export" {
		written, err := a.Export("dist")
		if err != nil {
			panic(err)
		}
		println(strings.Join(written, "\n"))
		return
	}
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	if err := a.Run(addr); err != nil {
		panic(err)
	}
}
```

What happens on `GET /`:

1. The router runs the middleware chain. `web.CSRF()` issues an `XSRF-TOKEN` cookie and puts the token in the request context.
2. `DataFunc` builds a `NotesPage` struct from the store (in a real app, a `pkg/orm` query — see the [fullstack tutorial](/getting-started/fullstack-app-tutorial)) and reads the token with `web.CSRFFrom(r.Context())`.
3. The renderer executes `Notes.kiw` with that struct as `.`: `{count}` is `.Count`, `{#each notes as n}` ranges over `.Notes`.
4. Because `EmitProps` is set, the struct is also serialized into a `data-props` attribute on the view's root element, and `/static/notes.js` is injected before `</body>`.

`go run . export` calls `App.Export("dist")`: it writes `dist/about.html` (fixed data) and skips `/` (it has a `DataFunc`).

### Structs, maps, and query results

Any Go value works as page data:

- **Structs** (recommended): fields are read by name — `{count}` reads `.Count`, `{n.Text}` reads the field `Text` of the loop variable.
- **Maps**: `map[string]any{"Count": 3}` behaves like a struct with those keys. Capitalize keys so the `{count}` shorthand finds them.
- **Query results**: pass the slice you scanned, e.g. `[]models.Task` from `db.Model(&models.Task{}).FindAll(&out)`, and loop over it with `{#each}`.

### The one rule for components: data goes in through props

A Krewire UI component (or your own) receives data **as attributes**: `<Badge text={summary} />`, `<Input value={draft} />`, `<Alert message={error} />`. The attribute expression is evaluated in the page's data context, so it sees `.Summary`, `.Draft`, `.Error`.

The **body** of a component — the content between `<Card>` and `</Card>` — is rendered separately, before it is handed to the component, and without the page's data. Page fields used inside a body render empty. Keep bodies static, or render the dynamic part outside the component and pass values as props. The [tutorial's `TaskBoard.kiw`](/getting-started/fullstack-app-tutorial) follows this pattern: the form and the list are plain HTML around self-closing components.

### Two things a live app wires up itself (today)

- **`.kiw` views:** an `--app` scaffold registers components written as Go strings. To write views as `.kiw` files, parse them with `kiw.ParseKiw` and register `ssg.Component{Name, Body, Style}` — as `main.go` above and the tutorial's `views.go` do. Register the Krewire UI catalog from `ui.ComponentsFS` the same way.
- **Component CSS:** pages link `/assets/style.css`, but a live `web.App` does not yet serve the collected `<style>` blocks of `.kiw` components. Concatenate them and register them with `a.Asset("assets/style.css", css)` (see the tutorial).

---

## 4. Choosing Between Them

| Question | If yes |
| :--- | :--- |
| Is the content the same for every visitor and changes only when you publish? | `site` / `book`, or an `app` page with `Data` that you `Export`. |
| Does it depend on the user, the session, or a database row that changes often? | `app` page with `DataFunc`. |
| Is it a small interactive part of an otherwise static page? | Static page + client script calling a JSON endpoint. |
| Must it work with JavaScript disabled? | Server-render it (`DataFunc`) and use HTML forms (post/redirect/get). |

The [progressive pipeline](/overview/why-krewire) is designed so you can move a page from the left column to the right one without rewriting its `.kiw` view: only where the data comes from changes.

---

## 5. Client-Side Interaction: Scripts, `fetch`, and Events

A client script enhances a server-rendered page; it does not replace it. In the example above, `PageSpec.Scripts` injects `static/notes.js`:

```text
// Progressive enhancement: the list is already server-rendered; this script
// adds notes through the JSON API without a full page reload.
(function () {
  var root = document.querySelector('[data-props]');
  var props = root ? JSON.parse(root.getAttribute('data-props')) : {};
  var list = document.getElementById('note-list');
  var form = document.getElementById('note-form');

  function render(notes) {
    list.textContent = '';
    notes.forEach(function (n) {
      var li = document.createElement('li');
      li.textContent = n.text; // textContent, never innerHTML, for user data
      list.appendChild(li);
    });
  }

  form.addEventListener('submit', function (e) {
    e.preventDefault();
    fetch('/api/notes', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', 'X-CSRF-Token': props.Token },
      body: JSON.stringify({ text: form.elements.text.value })
    })
      .then(function (r) { if (!r.ok) throw new Error(r.status); return fetch('/api/notes'); })
      .then(function (r) { return r.json(); })
      .then(function (notes) { render(notes); form.reset(); })
      .catch(function (err) { console.error('add note failed', err); });
  });
})();
```

How the pieces connect:

| Concern | Mechanism |
| :--- | :--- |
| Initial state | `EmitProps: true` → `data-props='{"Notes":…,"Count":…,"Token":"…"}'` on the root element; the script reads it with `JSON.parse` instead of making a first request. |
| Reading data | `fetch('/api/notes')` → a `Router().Get` handler that returns `web.JSON(w, 200, value)`. |
| Writing data | `fetch` with `method: 'POST'` and a JSON body → a `Router().Post` handler using `web.ReadJSON` and `validation.Struct`. |
| CSRF on `fetch` | Send the token in the `X-CSRF-Token` header (from `data-props`, or from the `XSRF-TOKEN` cookie, which is readable by scripts by default). Without it the middleware answers `403 {"code":"forbidden","message":"csrf token mismatch"}`. |
| Errors | `web.Error` writes `{"code":…,"message":…}` with the right status (`400` for validation, `403` for CSRF), so the script can branch on `r.ok`. |
| Re-rendering | Update the DOM with `textContent` (never `innerHTML` for user data), or reload the page to get a fresh server render. |
| Events | Plain DOM events (`submit`, `click`) in an ordinary script file; in an `app`, list its URL in `PageSpec.Scripts` and serve it with `App.Static`. |

This flow was run in a headless browser against the example: submitting the form sent `POST /api/notes` (`201`), followed by `GET /api/notes` (`200`), and the list updated without a reload.

For a `site`, the same script pattern applies, with the JSON coming from an API that lives elsewhere (another `app`, a `service`, or a third-party endpoint) — the static build itself has no server routes.

---

**Previous:** [11.3 Scope Levels](/architecture/scope-levels) · **Back to:** [Architecture Deep-Dive](/architecture)

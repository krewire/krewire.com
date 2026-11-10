// Package components defines the application's SSR components (FRK-STR-006).
// Pages reference these by name via PageSpec.Root.
package components

import "github.com/krewire/krewire/pkg/web/ssg"

// heroBody is the default landing component. It expects a map with the keys
// badge, title, subtitle, and an optional cta_left map of label/href.
const heroBody = `<section class="hero">
  {{if .badge}}<span class="hero-badge">{{.badge}}</span>{{end}}
  <h1>{{.title}}</h1>
  {{if .subtitle}}<p class="hero-subtitle">{{.subtitle}}</p>{{end}}
  {{if .cta_left}}<a class="hero-cta" href="{{.cta_left.href}}">{{.cta_left.label}}</a>{{end}}
</section>`

// Hero returns the default page component used by the scaffolded routes.
func Hero() ssg.Component {
	return ssg.Component{
		Name:  "Hero",
		Body:  heroBody,
		Style: ".hero{display:grid;gap:.75rem;padding:3rem 0}.hero h1{margin:0;font-size:2.25rem}.hero-subtitle{opacity:.75}.hero-badge{font-size:.75rem;letter-spacing:.08em;text-transform:uppercase;opacity:.6}.hero-cta{margin-top:1rem;padding:.6rem 1.1rem;border-radius:.5rem;background:var(--color-primary,#39D353);color:#000;text-decoration:none;font-weight:600}",
	}
}

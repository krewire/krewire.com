package mcp

import (
	"fmt"
	"strings"
	"sync"
)

// ComponentMeta represents a Pop-Brutalist .kiw UI component definition.
type ComponentMeta struct {
	Name        string   `json:"name"`
	Category    string   `json:"category"` // "actions", "layout", "feedback", "forms", "navigation", "display"
	Description string   `json:"description"`
	Props       []Prop   `json:"props,omitempty"`
	Variants    []string `json:"variants,omitempty"`
	Example     string   `json:"example"`
}

// Prop represents a component property.
type Prop struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Default     string `json:"default,omitempty"`
	Description string `json:"description"`
}

// ComponentCatalog provides the authoritative catalog of Krewire UI (.kiw) components.
type ComponentCatalog struct {
	mu         sync.RWMutex
	components map[string]ComponentMeta
}

// NewComponentCatalog initializes the Pop-Brutalist UI component registry.
func NewComponentCatalog() *ComponentCatalog {
	c := &ComponentCatalog{
		components: make(map[string]ComponentMeta),
	}
	c.seed()
	return c
}

// List returns components filtered by optional category.
func (c *ComponentCatalog) List(category string) []ComponentMeta {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var out []ComponentMeta
	catFilter := strings.ToLower(strings.TrimSpace(category))
	for _, comp := range c.components {
		if catFilter == "" || strings.ToLower(comp.Category) == catFilter {
			out = append(out, comp)
		}
	}
	return out
}

// Get returns component details by name.
func (c *ComponentCatalog) Get(name string) (*ComponentMeta, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	key := strings.ToLower(strings.TrimSpace(name))
	for k, comp := range c.components {
		if strings.ToLower(k) == key || strings.ToLower(comp.Name) == key {
			return &comp, nil
		}
	}
	return nil, fmt.Errorf("component %q not found in Krewire UI catalog", name)
}

func (c *ComponentCatalog) seed() {
	components := []ComponentMeta{
		{
			Name:        "Button",
			Category:    "actions",
			Description: "Pop-brutalist tactile button with bold border and hard drop shadow.",
			Variants:    []string{"primary", "secondary", "outline", "danger", "ghost"},
			Props: []Prop{
				{Name: "Variant", Type: "string", Default: "primary", Description: "Visual style variant"},
				{Name: "Size", Type: "string", Default: "md", Description: "Size: sm, md, lg"},
				{Name: "Href", Type: "string", Default: "", Description: "Renders as anchor tag if set"},
				{Name: "Disabled", Type: "bool", Default: "false", Description: "Disables interaction"},
			},
			Example: `<Button Variant="primary" Size="lg" Href="/docs">Get Started →</Button>`,
		},
		{
			Name:        "Card",
			Category:    "layout",
			Description: "Hard-edged container with high-contrast border and deep brutalist shadow.",
			Variants:    []string{"default", "raised", "outline", "accent"},
			Props: []Prop{
				{Name: "Title", Type: "string", Default: "", Description: "Card heading title"},
				{Name: "Variant", Type: "string", Default: "default", Description: "Style preset"},
				{Name: "Class", Type: "string", Default: "", Description: "Additional utility classes"},
			},
			Example: `<Card Title="Ecosystem Battery" Variant="raised">
  <p>Modular Go monorepo orchestrating 10 repositories.</p>
</Card>`,
		},
		{
			Name:        "Badge",
			Category:    "feedback",
			Description: "Compact status and tag pill with pop-brutalist border and dot indicator.",
			Variants:    []string{"primary", "success", "warning", "danger", "neutral"},
			Props: []Prop{
				{Name: "Variant", Type: "string", Default: "primary", Description: "Color theme"},
				{Name: "Dot", Type: "bool", Default: "false", Description: "Whether to render status pulse dot"},
			},
			Example: `<Badge Variant="success" Dot=true>OPERATIONAL</Badge>`,
		},
		{
			Name:        "Modal",
			Category:    "feedback",
			Description: "Accessible dialogue modal with backdrop overlay and brutalist frame.",
			Variants:    []string{"default", "danger"},
			Props: []Prop{
				{Name: "Title", Type: "string", Default: "", Description: "Dialog header"},
				{Name: "Open", Type: "bool", Default: "false", Description: "Visibility state"},
			},
			Example: `<Modal Title="Confirm Deployment" Open=true>
  <p>Proceed with rolling restart?</p>
</Modal>`,
		},
		{
			Name:        "Input",
			Category:    "forms",
			Description: "High-contrast text input field with tactile focus border.",
			Variants:    []string{"default", "error"},
			Props: []Prop{
				{Name: "Type", Type: "string", Default: "text", Description: "Input type (text, password, email)"},
				{Name: "Placeholder", Type: "string", Default: "", Description: "Hint placeholder"},
				{Name: "Disabled", Type: "bool", Default: "false", Description: "Disabled state"},
			},
			Example: `<Input Type="email" Placeholder="dev@krewire.com" />`,
		},
		{
			Name:        "Navbar",
			Category:    "navigation",
			Description: "Responsive top navigation bar with brand slot, nav links, and mobile menu.",
			Variants:    []string{"default", "bordered"},
			Props: []Prop{
				{Name: "Brand", Type: "string", Default: "Krewire", Description: "Brand title"},
				{Name: "Logo", Type: "string", Default: "", Description: "Logo asset path"},
			},
			Example: `<Navbar Brand="Krewire" Logo="/assets/images/logo.png" />`,
		},
		{
			Name:        "Terminal",
			Category:    "display",
			Description: "Brutalist CLI terminal emulator window with syntax prompt and copy button.",
			Variants:    []string{"dark", "light"},
			Props: []Prop{
				{Name: "Command", Type: "string", Default: "", Description: "Terminal command line"},
				{Name: "PromptColor", Type: "string", Default: "primary", Description: "Prompt indicator color"},
			},
			Example: `<Terminal Command="kiw new blog --app && cd blog && kiw dev" />`,
		},
		{
			Name:        "Alert",
			Category:    "feedback",
			Description: "Notification banner with brutalist border and status icons.",
			Variants:    []string{"info", "success", "warning", "danger"},
			Props: []Prop{
				{Name: "Variant", Type: "string", Default: "info", Description: "Alert severity"},
				{Name: "Title", Type: "string", Default: "", Description: "Alert headline"},
			},
			Example: `<Alert Variant="warning" Title="Database Migration Required">
  Run 'krew migrate up' before starting worker.
</Alert>`,
		},
		{
			Name:        "Tabs",
			Category:    "navigation",
			Description: "Segmented tab bar for switching content views and language toggles.",
			Variants:    []string{"default", "pills"},
			Props: []Prop{
				{Name: "Active", Type: "string", Default: "", Description: "Active tab identifier"},
			},
			Example: `<Tabs Active="go">
  <Tab Title="Go" Id="go">go run .</Tab>
  <Tab Title="CLI" Id="cli">kiw run</Tab>
</Tabs>`,
		},
	}

	for _, comp := range components {
		c.components[comp.Name] = comp
	}
}

package mcp

// RepoInfo documents a repository in the Krewire digital SDLC ecosystem.
type RepoInfo struct {
	Name        string   `json:"name"`
	Layer       int      `json:"layer"` // 0 to 5
	Role        string   `json:"role"`
	Description string   `json:"description"`
	Packages    []string `json:"packages,omitempty"`
}

// ArchitectureRule represents a core engineering invariant.
type ArchitectureRule struct {
	ID          string `json:"id"`
	Pillar      string `json:"pillar"` // "S1: Secure", "S2: Sustainable", "S3: Scalable"
	Title       string `json:"title"`
	Description string `json:"description"`
	Enforcement string `json:"enforcement"`
}

// GetEcosystem returns the authoritative breakdown of the 10 Krewire repositories.
func GetEcosystem() []RepoInfo {
	return []RepoInfo{
		{
			Name:        "krewire",
			Layer:       2,
			Role:        "Foundational Go Monorepo & AI Engine",
			Description: "Core domain packages (web, ui, kiw, orm, cache, resilience, sec, storage) and autonomous ReAct agent engine (pkg/ai).",
			Packages:    []string{"pkg/model", "pkg/lifecycle", "pkg/errs", "pkg/version", "pkg/web", "pkg/kiw", "pkg/orm", "pkg/cache", "pkg/sec", "pkg/ai"},
		},
		{
			Name:        "krew",
			Layer:       3,
			Role:        "Universal Developer Toolchain & CLI Driver",
			Description: "Single CLI driver (cmd/krew) for project lifecycle, live reload, builds, scaffolding, database migrations, and MCP devtools.",
			Packages:    []string{"cmd/krew", "internal/cli", "internal/commands", "internal/scaffold", "internal/mcp"},
		},
		{
			Name:        "tui",
			Layer:       2,
			Role:        "Pure-Go Terminal User Interface Primitives",
			Description: "Lightweight, zero-external-dependency TUI framework for terminal applications, widgets, and live command outputs.",
			Packages:    []string{"term", "widgets"},
		},
		{
			Name:        "ui",
			Layer:       2,
			Role:        "Pop-Brutalist Design System (.kiw)",
			Description: "Pop-Brutalist .kiw web component catalog (shadcn/ui distribution model, pure Go and Tailwind CSS).",
			Packages:    []string{"components", "styles"},
		},
		{
			Name:        "cloud",
			Layer:       4,
			Role:        "Enterprise Hybrid PaaS & Control Plane",
			Description: "Production PaaS runner, supervisor, process pools, edge gateway, and multi-tenant control plane.",
			Packages:    []string{"cmd/krewire-cloud-server", "pkg/gateway", "pkg/supervisor"},
		},
		{
			Name:        "auth",
			Layer:       4,
			Role:        "Enterprise Identity & Access Management (IAM)",
			Description: "Identity store, PBKDF2/TOTP, OpenID Connect (OIDC/PKCE), SAML 2.0, and hierarchical RBAC platform.",
			Packages:    []string{"cmd/krewire-auth-server", "pkg/oidc", "pkg/saml", "pkg/rbac"},
		},
		{
			Name:        "dash",
			Layer:       5,
			Role:        "AI-Native Virtual Workspace & Sandboxes",
			Description: "Isolated execution sandboxes, PTY terminal streaming, and collaborative dev presence.",
			Packages:    []string{"cmd/krewire-dash-server", "pkg/sandbox", "pkg/terminal"},
		},
		{
			Name:        "internal",
			Layer:       5,
			Role:        "Architecture Hub & Governance",
			Description: "Private documentation hub containing Architectural Decision Records (ADRs), roadmaps, and contributor guides.",
		},
		{
			Name:        "krewire.com",
			Layer:       5,
			Role:        "Official Portal, Docs & Public MCP Server",
			Description: "Official website, documentation portal, and public Model Context Protocol server (https://krewire.com).",
		},
		{
			Name:        "krewire.github.io",
			Layer:       5,
			Role:        "Open-Source Community & Sponsorship Portal",
			Description: "Public open-source sponsorship and community engagement portal (open.krewire.com).",
		},
	}
}

// GetArchitectureRules returns the immutable 6-Layer Pyramid rules and 3S filters.
func GetArchitectureRules() []ArchitectureRule {
	return []ArchitectureRule{
		{
			ID:          "RULE-PYRAMID-DOWNWARD",
			Pillar:      "S3: Scalable",
			Title:       "Strict Downward-Only Dependencies",
			Description: "Dependencies flow only downward (Layer N -> Layer < N). Higher layers must never be imported by lower layers.",
			Enforcement: "Enforced by compile checks and layering guards.",
		},
		{
			ID:          "RULE-PURE-GO",
			Pillar:      "S2: Sustainable",
			Title:       "Pure-Go Invariant (CGO_ENABLED=0)",
			Description: "All Go code must compile cleanly without CGO or external C dependencies for lean, single static binary deployment.",
			Enforcement: "CGO_ENABLED=0 go build ./... in all CI/CD pipelines.",
		},
		{
			ID:          "RULE-ZERO-PEER-COUPLING",
			Pillar:      "S3: Scalable",
			Title:       "Zero Peer Coupling (Lateral Independence)",
			Description: "Sibling platforms at Layer 4 (cloud, auth) and Layer 5 (dash) must never import each other directly. Use API gateways or token federation.",
			Enforcement: "Strict repository boundaries and go.work separation.",
		},
		{
			ID:          "RULE-SAFE-PATH",
			Pillar:      "S1: Secure",
			Title:       "SafePath Traversal Prevention",
			Description: "All filesystem operations reading or writing user/workspace input must be sanitized and bounded via sec.SafePath.",
			Enforcement: "Mandatory sec.SafePath call before any os.WriteFile or os.ReadFile.",
		},
		{
			ID:          "RULE-STDLIB-FIRST",
			Pillar:      "S1: Secure",
			Title:       "Stdlib-First Boundary",
			Description: "Always prefer standard library packages (net/http, crypto, os, sync) over third-party dependencies.",
			Enforcement: "Zero third-party runtime dependencies in core framework.",
		},
		{
			ID:          "RULE-REUSABLE-PROMOTION",
			Pillar:      "S2: Sustainable",
			Title:       "Reusable Module Placement Rule",
			Description: "Any reusable package or common domain logic needed by more than one platform MUST be placed in krewire/pkg/*.",
			Enforcement: "Downstream platforms (auth, cloud, dash) never duplicate common infrastructure.",
		},
	}
}

// LayeringEvaluation represents the result of validating a proposed import or peer connection.
type LayeringEvaluation struct {
	Allowed     bool   `json:"allowed"`
	FromRepo    string `json:"from_repo"`
	FromLayer   int    `json:"from_layer"`
	ToRepo      string `json:"to_repo"`
	ToLayer     int    `json:"to_layer"`
	Explanation string `json:"explanation"`
	Rule        string `json:"rule"`
}

var repoLayerMap = map[string]int{
	"stdlib":            0,
	"pkg/model":         1,
	"pkg/lifecycle":     1,
	"pkg/errs":          1,
	"pkg/version":       1,
	"krewire":           2,
	"pkg/web":           2,
	"pkg/orm":           2,
	"pkg/cache":         2,
	"pkg/sec":           2,
	"pkg/ai":            2,
	"tui":               2,
	"ui":                2,
	"krew":              3,
	"cloud":             4,
	"auth":              4,
	"dash":              5,
	"internal":          5,
	"krewire.com":       5,
	"krewire.github.io": 5,
}

// CheckLayering validates whether an import from one repo/package to another satisfies the 6-Layer Pyramid.
func CheckLayering(from, to string) LayeringEvaluation {
	fromLayer, hasFrom := repoLayerMap[from]
	if !hasFrom {
		fromLayer = 2 // default to domain library
	}
	toLayer, hasTo := repoLayerMap[to]
	if !hasTo {
		toLayer = 2
	}

	if from == to {
		return LayeringEvaluation{
			Allowed:     true,
			FromRepo:    from,
			FromLayer:   fromLayer,
			ToRepo:      to,
			ToLayer:     toLayer,
			Explanation: "Same module/repository internal import is permitted.",
			Rule:        "INTERNAL-PACKAGE-IMPORT",
		}
	}

	if fromLayer < toLayer {
		return LayeringEvaluation{
			Allowed:     false,
			FromRepo:    from,
			FromLayer:   fromLayer,
			ToRepo:      to,
			ToLayer:     toLayer,
			Explanation: "Upward dependency violation: Layer " + string(rune('0'+fromLayer)) + " (" + from + ") cannot import higher Layer " + string(rune('0'+toLayer)) + " (" + to + ").",
			Rule:        "RULE-PYRAMID-DOWNWARD",
		}
	}

	if fromLayer == toLayer && (fromLayer == 4 || fromLayer == 5) {
		return LayeringEvaluation{
			Allowed:     false,
			FromRepo:    from,
			FromLayer:   fromLayer,
			ToRepo:      to,
			ToLayer:     toLayer,
			Explanation: "Zero peer coupling violation: Sibling platforms at Layer " + string(rune('0'+fromLayer)) + " (" + from + " and " + to + ") must not import each other. Use API Gateways or token federation.",
			Rule:        "RULE-ZERO-PEER-COUPLING",
		}
	}

	return LayeringEvaluation{
		Allowed:     true,
		FromRepo:    from,
		FromLayer:   fromLayer,
		ToRepo:      to,
		ToLayer:     toLayer,
		Explanation: "Valid downward dependency: Layer " + string(rune('0'+fromLayer)) + " imports Layer " + string(rune('0'+toLayer)) + ".",
		Rule:        "RULE-PYRAMID-DOWNWARD",
	}
}

// CLICommandInfo documents a developer CLI command in krew.
type CLICommandInfo struct {
	Command     string `json:"command"`
	Description string `json:"description"`
	Usage       string `json:"usage"`
	Layer       int    `json:"layer"`
}

// GetCLICommands returns the canonical krew CLI developer driver commands.
func GetCLICommands() []CLICommandInfo {
	return []CLICommandInfo{
		{Command: "krew init", Description: "Scaffold a new Krewire project with prewired batteries for site, book, app, cli, worker, or service.", Usage: "krew init [name] [--site|--book|--app|--cli]", Layer: 3},
		{Command: "krew dev", Description: "Start live-reload development server with hot module replacement.", Usage: "krew dev [--port 8080]", Layer: 3},
		{Command: "krew build", Description: "Compile single static pure-Go binary with embedded assets (CGO_ENABLED=0).", Usage: "krew build [--release]", Layer: 3},
		{Command: "krew make migration", Description: "Generate a new Laravel-style code-first database migration under database/migrations/.", Usage: "krew make migration <name>", Layer: 3},
		{Command: "krew migrate", Description: "Execute pending database migrations with distributed advisory locking.", Usage: "krew migrate [--rollback|--status]", Layer: 3},
		{Command: "krew make component", Description: "Generate a Pop-Brutalist .kiw component in the components catalog.", Usage: "krew make component <Name>", Layer: 3},
		{Command: "krew deploy", Description: "Deploy project to Krewire Cloud PaaS or self-hosted runner.", Usage: "krew deploy [--cloud|--prod]", Layer: 3},
		{Command: "krew compat", Description: "Audit cross-repository version compatibility constraints across the workspace.", Usage: "krew compat", Layer: 3},
		{Command: "krew vet", Description: "Run formatting, static analysis, layering invariant checks, and security audit.", Usage: "krew vet [--sec]", Layer: 3},
		{Command: "krew ai", Description: "Launch autonomous developer assistant or ReAct agent in terminal.", Usage: "krew ai [prompt]", Layer: 3},
	}
}

// GenerateMigrationTemplate returns a Laravel-style Go code-first migration file for pkg/orm/migrate.
func GenerateMigrationTemplate(tableName, migrationName string) string {
	if migrationName == "" {
		migrationName = "create_" + tableName + "_table"
	}
	return `package migrations

import (
	"context"

	"github.com/krewire/krewire/pkg/orm/migrate"
	"github.com/krewire/krewire/pkg/orm/schema"
)

func init() {
	migrate.Register(&migrate.GoMigration{
		Version: "20261123000000",
		Name:    "` + migrationName + `",
		Up: func(ctx context.Context, s *schema.Builder) error {
			return s.Create("` + tableName + `", func(t *schema.Blueprint) {
				t.ID()
				t.String("name", 255)
				t.Bool("is_active").Default(true)
				t.Timestamps()
			})
		},
		Down: func(ctx context.Context, s *schema.Builder) error {
			return s.DropIfExists("` + tableName + `")
		},
	})
}
`
}

// ScaffoldKrewireYAML generates a valid krewire.yaml declarative configuration.
func ScaffoldKrewireYAML(template, name string) string {
	if template == "" {
		template = "app"
	}
	if name == "" {
		name = "my-" + template
	}
	return `version: "1.0"
name: "` + name + `"
stability: "stable"

build:
  target: "bin/` + name + `"
  cgo: false
  flags: ["-s", "-w"]

server:
  host: "0.0.0.0"
  port: 8080
  read_header_timeout: "5s"
  max_body_bytes: "2MB"

security:
  csp: "default-src 'self'"
  hsts: true
  csrf: true
`
}

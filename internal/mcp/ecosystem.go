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
			Description: "Pop-Brutalist .kiw web component catalog (shadcn/ui distribution model, zero npm runtime dependencies).",
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

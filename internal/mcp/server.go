package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ProtocolVersion is the MCP specification version supported.
const ProtocolVersion = "2024-11-05"

// Request represents a JSON-RPC 2.0 request or notification.
type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      any             `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

// Response represents a JSON-RPC 2.0 response.
type Response struct {
	JSONRPC string `json:"jsonrpc"`
	ID      any    `json:"id"`
	Result  any    `json:"result,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

// Error represents a JSON-RPC 2.0 error object.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Standard JSON-RPC error codes.
const (
	CodeParseError     = -32700
	CodeInvalidRequest = -32600
	CodeMethodNotFound = -32601
	CodeInvalidParams  = -32602
	CodeInternalError  = -32603
)

// Server coordinates public MCP protocol interactions for Krewire documentation and ecosystem.
type Server struct {
	docs       *DocIndexer
	components *ComponentCatalog
	startTime  time.Time
}

// NewServer initializes the Krewire public MCP Server.
func NewServer(docs *DocIndexer) *Server {
	return &Server{
		docs:       docs,
		components: NewComponentCatalog(),
		startTime:  time.Now(),
	}
}

// HandleRequest processes an incoming JSON-RPC request and returns a response,
// or nil if the request was a notification.
func (s *Server) HandleRequest(ctx context.Context, req Request) *Response {
	if req.ID == nil && strings.HasPrefix(req.Method, "notifications/") {
		return nil // Notifications receive no response
	}

	res, err := s.dispatch(ctx, req.Method, req.Params)
	if err != nil {
		return &Response{
			JSONRPC: "2.0",
			ID:      req.ID,
			Error:   err,
		}
	}

	return &Response{
		JSONRPC: "2.0",
		ID:      req.ID,
		Result:  res,
	}
}

func (s *Server) dispatch(ctx context.Context, method string, paramsRaw json.RawMessage) (any, *Error) {
	switch method {
	case "initialize":
		return map[string]any{
			"protocolVersion": ProtocolVersion,
			"capabilities": map[string]any{
				"tools":     map[string]any{},
				"resources": map[string]any{},
				"prompts":   map[string]any{},
			},
			"serverInfo": map[string]any{
				"name":    "krewire",
				"version": "0.2.4",
			},
		}, nil

	case "ping":
		return map[string]any{}, nil

	case "resources/list":
		return map[string]any{
			"resources": []map[string]any{
				{
					"uri":         "krewire://docs/toc",
					"name":        "Documentation Table of Contents",
					"description": "Hierarchical index of all Krewire documentation chapters and guides",
					"mimeType":    "application/json",
				},
				{
					"uri":         "krewire://ecosystem",
					"name":        "Krewire Ecosystem Matrix",
					"description": "Overview of the 10 Krewire repositories, layers, and package responsibilities",
					"mimeType":    "application/json",
				},
				{
					"uri":         "krewire://components",
					"name":        "Pop-Brutalist Component Catalog",
					"description": "Complete list of available .kiw UI components, props, and variants",
					"mimeType":    "application/json",
				},
				{
					"uri":         "krewire://rules",
					"name":        "Architecture Invariants & 3S Filter",
					"description": "The 6-Layer Pyramid rules, Three Pillars (S1-S3), and coding guidelines",
					"mimeType":    "application/json",
				},
			},
		}, nil

	case "resources/read":
		var p struct {
			URI string `json:"uri"`
		}
		if err := json.Unmarshal(paramsRaw, &p); err != nil {
			return nil, &Error{Code: CodeInvalidParams, Message: "invalid resource uri parameter"}
		}
		text, err := s.readResource(p.URI)
		if err != nil {
			return nil, &Error{Code: CodeInvalidParams, Message: err.Error()}
		}
		return map[string]any{
			"contents": []map[string]any{
				{
					"uri":      p.URI,
					"mimeType": "application/json",
					"text":     text,
				},
			},
		}, nil

	case "prompts/list":
		return map[string]any{
			"prompts": []map[string]any{
				{
					"name":        "krewire-architect",
					"description": "System instructions for designing applications on Krewire's 6-Layer Pyramid Architecture.",
				},
				{
					"name":        "kiw-component-builder",
					"description": "Guidance on building accessible Pop-Brutalist .kiw web components.",
				},
				{
					"name":        "krewire-migrator",
					"description": "Instructions for writing Laravel-style database migrations in pure Go.",
				},
			},
		}, nil

	case "prompts/get":
		var p struct {
			Name string `json:"name"`
		}
		if err := json.Unmarshal(paramsRaw, &p); err != nil {
			return nil, &Error{Code: CodeInvalidParams, Message: "invalid prompt parameters"}
		}
		return s.getPrompt(p.Name)

	case "tools/list":
		return map[string]any{
			"tools": []map[string]any{
				{
					"name":        "search_docs",
					"description": "Search official Krewire documentation, guides, architecture notes, and concepts.",
					"inputSchema": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"query": map[string]any{
								"type":        "string",
								"description": "Search query keywords (e.g. 'orm postgres', 'kiw component', 'auth oidc')",
							},
							"locale": map[string]any{
								"type":        "string",
								"description": "Language code: 'en' (English) or 'id' (Indonesian), default: 'en'",
							},
							"limit": map[string]any{
								"type":        "integer",
								"description": "Maximum number of results to return (default: 5)",
							},
						},
						"required": []string{"query"},
					},
				},
				{
					"name":        "get_doc",
					"description": "Retrieve the full markdown content of an official Krewire documentation chapter.",
					"inputSchema": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"path": map[string]any{
								"type":        "string",
								"description": "Documentation path or slug (e.g. '02-getting-started/01-quickstart')",
							},
							"locale": map[string]any{
								"type":        "string",
								"description": "Language code: 'en' or 'id', default: 'en'",
							},
						},
						"required": []string{"path"},
					},
				},
				{
					"name":        "list_docs",
					"description": "List all official Krewire documentation chapters and table of contents.",
					"inputSchema": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"locale": map[string]any{
								"type":        "string",
								"description": "Language code: 'en' or 'id', default: 'en'",
							},
						},
					},
				},
				{
					"name":        "list_components",
					"description": "List all Pop-Brutalist .kiw UI components available in the Krewire design system.",
					"inputSchema": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"category": map[string]any{
								"type":        "string",
								"description": "Optional category filter: actions, layout, feedback, forms, navigation, display",
							},
						},
					},
				},
				{
					"name":        "get_component",
					"description": "Get detailed properties, style variants, and code examples for a .kiw UI component.",
					"inputSchema": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"name": map[string]any{
								"type":        "string",
								"description": "Component name (e.g. Button, Card, Modal, Input, Badge)",
							},
						},
						"required": []string{"name"},
					},
				},
				{
					"name":        "get_ecosystem",
					"description": "Get the authoritative 10-repository architectural breakdown, layer classification (Layer 0–5), and package layout.",
					"inputSchema": map[string]any{
						"type":       "object",
						"properties": map[string]any{},
					},
				},
				{
					"name":        "get_architecture_rules",
					"description": "Get the immutable architectural invariants (6-Layer Pyramid, Three Pillars S1-S3, safePath guard, stdlib-first, pure-Go invariant) required for AI agents building in Krewire.",
					"inputSchema": map[string]any{
						"type":       "object",
						"properties": map[string]any{},
					},
				},
				{
					"name":        "check_layering_rule",
					"description": "Validate whether an import from a source repository/package to a target satisfies the 6-Layer Pyramid Architecture rules.",
					"inputSchema": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"from": map[string]any{
								"type":        "string",
								"description": "Source repository or package (e.g. 'cloud', 'krew', 'pkg/web')",
							},
							"to": map[string]any{
								"type":        "string",
								"description": "Target repository or package (e.g. 'dash', 'auth', 'pkg/model')",
							},
						},
						"required": []string{"from", "to"},
					},
				},
				{
					"name":        "generate_migration",
					"description": "Generate a Go code-first Laravel-style database migration file template for pkg/orm/migrate.",
					"inputSchema": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"table_name": map[string]any{
								"type":        "string",
								"description": "Target table name (e.g. 'users', 'posts', 'api_tokens')",
							},
							"migration_name": map[string]any{
								"type":        "string",
								"description": "Migration name slug (e.g. 'create_users_table')",
							},
						},
						"required": []string{"table_name"},
					},
				},
				{
					"name":        "scaffold_project_config",
					"description": "Scaffold a canonical krewire.yaml declarative configuration.",
					"inputSchema": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"template": map[string]any{
								"type":        "string",
								"description": "Starter template: app, cli, site, book, worker, service, gateway, or infra",
							},
							"name": map[string]any{
								"type":        "string",
								"description": "Project identifier name",
							},
						},
					},
				},
				{
					"name":        "get_cli_commands",
					"description": "Get the canonical developer CLI command dictionary and usage options for krew.",
					"inputSchema": map[string]any{
						"type":       "object",
						"properties": map[string]any{},
					},
				},
			},
		}, nil

	case "tools/call":
		var p struct {
			Name      string          `json:"name"`
			Arguments json.RawMessage `json:"arguments"`
		}
		if err := json.Unmarshal(paramsRaw, &p); err != nil {
			return nil, &Error{Code: CodeInvalidParams, Message: "invalid tool call parameters"}
		}
		resultText, isErr := s.callTool(ctx, p.Name, p.Arguments)
		return map[string]any{
			"content": []map[string]any{
				{
					"type": "text",
					"text": resultText,
				},
			},
			"isError": isErr,
		}, nil

	default:
		return nil, &Error{Code: CodeMethodNotFound, Message: fmt.Sprintf("method not found: %s", method)}
	}
}

func (s *Server) readResource(uri string) (string, error) {
	switch uri {
	case "krewire://docs/toc":
		docs := s.docs.List("en")
		b, _ := json.MarshalIndent(docs, "", "  ")
		return string(b), nil

	case "krewire://ecosystem":
		eco := GetEcosystem()
		b, _ := json.MarshalIndent(eco, "", "  ")
		return string(b), nil

	case "krewire://components":
		comps := s.components.List("")
		b, _ := json.MarshalIndent(comps, "", "  ")
		return string(b), nil

	case "krewire://rules":
		rules := GetArchitectureRules()
		b, _ := json.MarshalIndent(rules, "", "  ")
		return string(b), nil

	default:
		return "", fmt.Errorf("unknown resource uri: %s", uri)
	}
}

func (s *Server) getPrompt(name string) (any, *Error) {
	switch name {
	case "krewire-architect":
		return map[string]any{
			"description": "System context for architecting applications within Krewire",
			"messages": []map[string]any{
				{
					"role": "user",
					"content": map[string]any{
						"type": "text",
						"text": "You are building an application in the Krewire digital SDLC ecosystem. Enforce the 6-Layer Pyramid Architecture strictly: Layer 0 (stdlib) -> Layer 1 (model/lifecycle) -> Layer 2 (domain libraries: orm, web, ui) -> Layer 3 (krew devtool) -> Layer 4 (cloud/auth) -> Layer 5 (dash/portal). All code must be pure Go (CGO_ENABLED=0) and follow the Three Pillars: Secure (stdlib-first, safePath), Sustainable (lean footprint, single binary, pure Go), and Scalable (reversible provider swaps).",
					},
				},
			},
		}, nil

	case "kiw-component-builder":
		return map[string]any{
			"description": "Guidance on authoring Pop-Brutalist .kiw components",
			"messages": []map[string]any{
				{
					"role": "user",
					"content": map[string]any{
						"type": "text",
						"text": "Write a Pop-Brutalist .kiw component using Krewire UI standards. Components use Go template syntax with high-contrast borders (border-2 border-border), hard drop shadows, and scoped CSS styles. Props are read lower-camel casing. No external npm dependencies.",
					},
				},
			},
		}, nil

	case "krewire-migrator":
		return map[string]any{
			"description": "Guidance on writing Laravel-style database migrations in Go",
			"messages": []map[string]any{
				{
					"role": "user",
					"content": map[string]any{
						"type": "text",
						"text": "Write a programmatic Go database migration using 's.Create(\"table\", func(table *schema.Blueprint) { ... })' and 's.Table(\"table\", func(table *schema.Blueprint) { ... })' mirroring Laravel's Schema Facade. Migration files reside in database/migrations/<timestamp>_<name>.go with 'package migrations' and an init() function registering migrate.GoMigration.",
					},
				},
			},
		}, nil

	default:
		return nil, &Error{Code: CodeInvalidParams, Message: fmt.Sprintf("prompt %q not found", name)}
	}
}

func (s *Server) callTool(ctx context.Context, name string, argsRaw json.RawMessage) (string, bool) {
	switch name {
	case "search_docs":
		var p struct {
			Query  string `json:"query"`
			Locale string `json:"locale"`
			Limit  int    `json:"limit"`
		}
		if err := json.Unmarshal(argsRaw, &p); err != nil {
			return fmt.Sprintf("invalid arguments: %v", err), true
		}
		results := s.docs.Search(p.Query, p.Locale, p.Limit)
		b, _ := json.MarshalIndent(results, "", "  ")
		return string(b), false

	case "get_doc":
		var p struct {
			Path   string `json:"path"`
			Locale string `json:"locale"`
		}
		if err := json.Unmarshal(argsRaw, &p); err != nil {
			return fmt.Sprintf("invalid arguments: %v", err), true
		}
		page, err := s.docs.Get(p.Path, p.Locale)
		if err != nil {
			return fmt.Sprintf("error retrieving doc: %v", err), true
		}
		b, _ := json.MarshalIndent(page, "", "  ")
		return string(b), false

	case "list_docs":
		var p struct {
			Locale string `json:"locale"`
		}
		if len(argsRaw) > 0 {
			_ = json.Unmarshal(argsRaw, &p)
		}
		docs := s.docs.List(p.Locale)
		b, _ := json.MarshalIndent(docs, "", "  ")
		return string(b), false

	case "list_components":
		var p struct {
			Category string `json:"category"`
		}
		if len(argsRaw) > 0 {
			_ = json.Unmarshal(argsRaw, &p)
		}
		comps := s.components.List(p.Category)
		b, _ := json.MarshalIndent(comps, "", "  ")
		return string(b), false

	case "get_component":
		var p struct {
			Name string `json:"name"`
		}
		if err := json.Unmarshal(argsRaw, &p); err != nil {
			return fmt.Sprintf("invalid arguments: %v", err), true
		}
		comp, err := s.components.Get(p.Name)
		if err != nil {
			return fmt.Sprintf("error retrieving component: %v", err), true
		}
		b, _ := json.MarshalIndent(comp, "", "  ")
		return string(b), false

	case "get_ecosystem":
		eco := GetEcosystem()
		b, _ := json.MarshalIndent(eco, "", "  ")
		return string(b), false

	case "get_architecture_rules":
		rules := GetArchitectureRules()
		b, _ := json.MarshalIndent(rules, "", "  ")
		return string(b), false

	case "check_layering_rule":
		var p struct {
			From string `json:"from"`
			To   string `json:"to"`
		}
		if err := json.Unmarshal(argsRaw, &p); err != nil {
			return fmt.Sprintf("invalid arguments: %v", err), true
		}
		eval := CheckLayering(p.From, p.To)
		b, _ := json.MarshalIndent(eval, "", "  ")
		return string(b), false

	case "generate_migration":
		var p struct {
			TableName     string `json:"table_name"`
			MigrationName string `json:"migration_name"`
		}
		if err := json.Unmarshal(argsRaw, &p); err != nil {
			return fmt.Sprintf("invalid arguments: %v", err), true
		}
		code := GenerateMigrationTemplate(p.TableName, p.MigrationName)
		return code, false

	case "scaffold_project_config":
		var p struct {
			Template string `json:"template"`
			Kind     string `json:"kind"`
			Name     string `json:"name"`
		}
		if err := json.Unmarshal(argsRaw, &p); err != nil {
			return fmt.Sprintf("invalid arguments: %v", err), true
		}
		tmpl := p.Template
		if tmpl == "" {
			tmpl = p.Kind
		}
		yamlDoc := ScaffoldKrewireYAML(tmpl, p.Name)
		return yamlDoc, false

	case "get_cli_commands":
		cmds := GetCLICommands()
		b, _ := json.MarshalIndent(cmds, "", "  ")
		return string(b), false

	default:
		return fmt.Sprintf("unknown tool: %s", name), true
	}
}

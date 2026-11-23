package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"
)

func createTestIndexer() *DocIndexer {
	mockFS := fstest.MapFS{
		"content/en/docs/01-overview/01-why.md": &fstest.MapFile{
			Data: []byte(`---
title: "Why Krewire"
description: "Philosophy and motivation behind the ecosystem"
tags:
  - philosophy
  - foundation
---
# Why Krewire
Krewire is an end-to-end digital SDLC ecosystem in pure Go.
`),
		},
		"content/en/docs/02-getting-started/01-quickstart.md": &fstest.MapFile{
			Data: []byte(`---
title: "Quickstart"
description: "Get started with Krewire in 5 minutes"
tags:
  - quickstart
  - cli
---
# Quickstart
Install the CLI tool using curl: curl -fsSL https://krewire.com/scripts/install.sh | sh
`),
		},
	}

	idx, _ := NewDocIndexer(mockFS)
	return idx
}

func TestServer_Initialize(t *testing.T) {
	srv := NewServer(createTestIndexer())
	req := Request{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "initialize",
	}

	resp := srv.HandleRequest(context.Background(), req)
	if resp == nil || resp.Error != nil {
		t.Fatalf("unexpected error response: %+v", resp)
	}

	resMap, ok := resp.Result.(map[string]any)
	if !ok {
		t.Fatalf("expected map result, got %T", resp.Result)
	}
	if resMap["protocolVersion"] != ProtocolVersion {
		t.Errorf("expected protocol version %s, got %v", ProtocolVersion, resMap["protocolVersion"])
	}
}

func TestServer_ToolsListAndCall(t *testing.T) {
	srv := NewServer(createTestIndexer())

	// Test tools/list
	reqList := Request{
		JSONRPC: "2.0",
		ID:      2,
		Method:  "tools/list",
	}
	respList := srv.HandleRequest(context.Background(), reqList)
	if respList == nil || respList.Error != nil {
		t.Fatalf("tools/list failed: %+v", respList)
	}

	// Test search_docs
	searchArgs, _ := json.Marshal(map[string]any{
		"query":  "quickstart",
		"locale": "en",
	})
	reqSearch := Request{
		JSONRPC: "2.0",
		ID:      3,
		Method:  "tools/call",
		Params: json.RawMessage(fmtJSON(map[string]any{
			"name":      "search_docs",
			"arguments": searchArgs,
		})),
	}
	respSearch := srv.HandleRequest(context.Background(), reqSearch)
	if respSearch == nil || respSearch.Error != nil {
		t.Fatalf("search_docs failed: %+v", respSearch)
	}

	// Test get_component
	compArgs, _ := json.Marshal(map[string]any{
		"name": "Button",
	})
	reqComp := Request{
		JSONRPC: "2.0",
		ID:      4,
		Method:  "tools/call",
		Params: json.RawMessage(fmtJSON(map[string]any{
			"name":      "get_component",
			"arguments": compArgs,
		})),
	}
	respComp := srv.HandleRequest(context.Background(), reqComp)
	if respComp == nil || respComp.Error != nil {
		t.Fatalf("get_component failed: %+v", respComp)
	}

	// Test get_ecosystem
	reqEco := Request{
		JSONRPC: "2.0",
		ID:      5,
		Method:  "tools/call",
		Params:  json.RawMessage(`{"name": "get_ecosystem", "arguments": {}}`),
	}
	respEco := srv.HandleRequest(context.Background(), reqEco)
	if respEco == nil || respEco.Error != nil {
		t.Fatalf("get_ecosystem failed: %+v", respEco)
	}

	// Test get_architecture_rules
	reqRules := Request{
		JSONRPC: "2.0",
		ID:      6,
		Method:  "tools/call",
		Params:  json.RawMessage(`{"name": "get_architecture_rules", "arguments": {}}`),
	}
	respRules := srv.HandleRequest(context.Background(), reqRules)
	if respRules == nil || respRules.Error != nil {
		t.Fatalf("get_architecture_rules failed: %+v", respRules)
	}
}

func TestServer_Resources(t *testing.T) {
	srv := NewServer(createTestIndexer())

	reqList := Request{
		JSONRPC: "2.0",
		ID:      10,
		Method:  "resources/list",
	}
	respList := srv.HandleRequest(context.Background(), reqList)
	if respList == nil || respList.Error != nil {
		t.Fatalf("resources/list failed: %+v", respList)
	}

	reqRead := Request{
		JSONRPC: "2.0",
		ID:      11,
		Method:  "resources/read",
		Params:  json.RawMessage(`{"uri": "krewire://ecosystem"}`),
	}
	respRead := srv.HandleRequest(context.Background(), reqRead)
	if respRead == nil || respRead.Error != nil {
		t.Fatalf("resources/read failed: %+v", respRead)
	}
}

func TestHTTPTransport_MessageStateless(t *testing.T) {
	srv := NewServer(createTestIndexer())
	transport := NewHTTPTransport(srv)
	mux := http.NewServeMux()
	transport.Mount(mux, "/mcp")

	body := []byte(`{"jsonrpc":"2.0","id":100,"method":"ping"}`)
	req := httptest.NewRequest(http.MethodPost, "/mcp/message", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d: %s", w.Code, w.Body.String())
	}

	var resp Response
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resp.ID != float64(100) {
		t.Errorf("expected ID 100, got %v", resp.ID)
	}
}

func TestHTTPTransport_RootPOST(t *testing.T) {
	srv := NewServer(createTestIndexer())
	transport := NewHTTPTransport(srv)
	mux := http.NewServeMux()
	transport.Mount(mux, "/mcp")

	reqBody := `{"jsonrpc":"2.0","id":101,"method":"initialize","params":{}}`
	req := httptest.NewRequest(http.MethodPost, "/mcp", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d: %s", w.Code, w.Body.String())
	}

	var resp Response
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resp.ID != float64(101) {
		t.Errorf("expected ID 101, got %v", resp.ID)
	}
}

func TestHTTPTransport_Health(t *testing.T) {
	srv := NewServer(createTestIndexer())
	transport := NewHTTPTransport(srv)
	mux := http.NewServeMux()
	transport.Mount(mux, "/mcp")

	req := httptest.NewRequest(http.MethodGet, "/mcp/health", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), `"server":"krewire"`) {
		t.Errorf("expected server name krewire in health response: %s", w.Body.String())
	}
}

func TestServer_EcosystemTools(t *testing.T) {
	srv := NewServer(createTestIndexer())
	ctx := context.Background()

	// 1. check_layering_rule
	argsLayer, _ := json.Marshal(map[string]any{"from": "cloud", "to": "dash"})
	respLayer := srv.HandleRequest(ctx, Request{
		JSONRPC: "2.0",
		ID:      10,
		Method:  "tools/call",
		Params:  json.RawMessage(fmtJSON(map[string]any{"name": "check_layering_rule", "arguments": argsLayer})),
	})
	if respLayer == nil || respLayer.Error != nil {
		t.Fatalf("check_layering_rule failed: %+v", respLayer)
	}

	// 2. generate_migration
	argsMig, _ := json.Marshal(map[string]any{"table_name": "orders", "migration_name": "create_orders_table"})
	respMig := srv.HandleRequest(ctx, Request{
		JSONRPC: "2.0",
		ID:      11,
		Method:  "tools/call",
		Params:  json.RawMessage(fmtJSON(map[string]any{"name": "generate_migration", "arguments": argsMig})),
	})
	if respMig == nil || respMig.Error != nil {
		t.Fatalf("generate_migration failed: %+v", respMig)
	}

	// 3. scaffold_project_config
	argsScaffold, _ := json.Marshal(map[string]any{"kind": "service", "name": "billing-svc"})
	respScaffold := srv.HandleRequest(ctx, Request{
		JSONRPC: "2.0",
		ID:      12,
		Method:  "tools/call",
		Params:  json.RawMessage(fmtJSON(map[string]any{"name": "scaffold_project_config", "arguments": argsScaffold})),
	})
	if respScaffold == nil || respScaffold.Error != nil {
		t.Fatalf("scaffold_project_config failed: %+v", respScaffold)
	}

	// 4. get_cli_commands
	respCLI := srv.HandleRequest(ctx, Request{
		JSONRPC: "2.0",
		ID:      13,
		Method:  "tools/call",
		Params:  json.RawMessage(fmtJSON(map[string]any{"name": "get_cli_commands", "arguments": map[string]any{}})),
	})
	if respCLI == nil || respCLI.Error != nil {
		t.Fatalf("get_cli_commands failed: %+v", respCLI)
	}
}

func TestStdioTransport(t *testing.T) {
	srv := NewServer(createTestIndexer())
	in := strings.NewReader("{\"jsonrpc\":\"2.0\",\"id\":999,\"method\":\"ping\"}\n")
	var out bytes.Buffer

	transport := NewStdioTransport(srv, in, &out)
	if err := transport.Serve(context.Background()); err != nil {
		t.Fatalf("stdio serve failed: %v", err)
	}

	if !strings.Contains(out.String(), `"id":999`) {
		t.Errorf("expected response with id 999, got: %s", out.String())
	}
}

func fmtJSON(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestDocIndexer_GetDoc_AcceptanceCriteria(t *testing.T) {
	// Create indexer with bilingual documents and numeric prefixes
	mockFS := fstest.MapFS{
		"content/en/docs/00-index.md": &fstest.MapFile{
			Data: []byte(`---
title: "Documentation Overview"
description: "Master index of Krewire documentation"
---
# Documentation
Welcome to Krewire documentation.
`),
		},
		"content/en/docs/02-getting-started/00-index.md": &fstest.MapFile{
			Data: []byte(`---
title: "Getting Started"
description: "Beginner track for Krewire"
---
# Getting Started
Learn Krewire step by step.
`),
		},
		"content/en/docs/02-getting-started/08-mcp-server.md": &fstest.MapFile{
			Data: []byte(`---
title: "Model Context Protocol (MCP) Server"
description: "Connect AI coding agents directly to Krewire"
tags:
  - mcp
  - agents
---
# MCP Server
Krewire provides native MCP support.
`),
		},
		"content/id/docs/00-index.md": &fstest.MapFile{
			Data: []byte(`---
title: "Dokumentasi Krewire"
description: "Indeks utama dokumentasi Krewire"
---
# Dokumentasi
Selamat datang di dokumentasi Krewire.
`),
		},
		"content/id/docs/02-getting-started/08-mcp-server.md": &fstest.MapFile{
			Data: []byte(`---
title: "Server Model Context Protocol (MCP)"
description: "Hubungkan agen koding AI langsung ke Krewire"
tags:
  - mcp
  - agen
---
# Server MCP
Krewire menyediakan dukungan MCP native.
`),
		},
	}

	idx, err := NewDocIndexer(mockFS)
	if err != nil {
		t.Fatalf("failed to create indexer: %v", err)
	}

	srv := NewServer(idx)
	ctx := context.Background()

	// Criteria 1 & 2 & 3: Every doc returned by list_docs can be fetched by path and clean_path with non-empty fields
	for _, locale := range []string{"en", "id"} {
		docs := idx.List(locale)
		if len(docs) == 0 {
			t.Fatalf("expected non-empty docs list for locale %s", locale)
		}

		for _, d := range docs {
			// Test by Path
			docByPath, err := idx.Get(d.Path, locale)
			if err != nil {
				t.Errorf("[%s] failed to get doc by Path %q: %v", locale, d.Path, err)
			} else {
				validateDocPage(t, docByPath, locale)
			}

			// Test by CleanPath
			if d.CleanPath != "" {
				docByClean, err := idx.Get(d.CleanPath, locale)
				if err != nil {
					t.Errorf("[%s] failed to get doc by CleanPath %q: %v", locale, d.CleanPath, err)
				} else {
					validateDocPage(t, docByClean, locale)
				}
			}
		}
	}

	// Test path format variants for "02-getting-started/08-mcp-server"
	variants := []string{
		"02-getting-started/08-mcp-server",
		"getting-started/mcp-server",
		"/02-getting-started/08-mcp-server",
		"/getting-started/mcp-server",
		"02-getting-started/08-mcp-server.md",
		"getting-started/mcp-server.md",
		"getting-started/mcp-server.html",
		"docs/getting-started/mcp-server",
		"/docs/getting-started/mcp-server",
		"https://krewire.com/docs/getting-started/mcp-server",
		"http://localhost:8080/docs/getting-started/mcp-server",
		"  getting-started/mcp-server  ",
	}

	for _, v := range variants {
		doc, err := idx.Get(v, "en")
		if err != nil {
			t.Errorf("failed to retrieve variant %q: %v", v, err)
		} else if doc.Path != "02-getting-started/08-mcp-server" {
			t.Errorf("variant %q resolved to %q, expected %q", v, doc.Path, "02-getting-started/08-mcp-server")
		}
	}

	// Test Indonesian retrieval with locale in path
	idVariants := []string{
		"id/docs/getting-started/mcp-server",
		"id/getting-started/mcp-server",
		"/id/docs/getting-started/mcp-server",
		"https://krewire.com/id/docs/getting-started/mcp-server",
	}
	for _, v := range idVariants {
		doc, err := idx.Get(v, "")
		if err != nil {
			t.Errorf("failed to retrieve ID variant %q: %v", v, err)
		} else if doc.Locale != "id" {
			t.Errorf("ID variant %q returned locale %q, expected 'id'", v, doc.Locale)
		}
	}

	// Criteria 4: Unknown path returns isError: true with actionable error message
	badReq := Request{
		JSONRPC: "2.0",
		ID:      50,
		Method:  "tools/call",
		Params: json.RawMessage(`{
			"name": "get_doc",
			"arguments": {"path": "non-existent-ch-99"}
		}`),
	}
	respBad := srv.HandleRequest(ctx, badReq)
	if respBad == nil || respBad.Result == nil {
		t.Fatalf("expected non-nil response for bad path")
	}
	badMap, ok := respBad.Result.(map[string]any)
	if !ok {
		t.Fatalf("expected result map, got %T", respBad.Result)
	}
	if badMap["isError"] != true {
		t.Errorf("expected isError: true for unknown doc, got %v", badMap["isError"])
	}
	contentSlice := badMap["content"].([]map[string]any)
	errText := contentSlice[0]["text"].(string)
	if !strings.Contains(errText, "documentation page not found") || !strings.Contains(errText, "list_docs") {
		t.Errorf("expected actionable error message containing list_docs, got: %s", errText)
	}

	// Criteria 5: Regression test over HTTP transport
	transport := NewHTTPTransport(srv)
	mux := http.NewServeMux()
	transport.Mount(mux, "/mcp")

	httpReqBody := `{"jsonrpc":"2.0","id":60,"method":"tools/call","params":{"name":"get_doc","arguments":{"path":"getting-started/mcp-server"}}}`
	httpReq := httptest.NewRequest(http.MethodPost, "/mcp/message", strings.NewReader(httpReqBody))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, httpReq)
	if w.Code != http.StatusOK {
		t.Fatalf("HTTP get_doc returned %d: %s", w.Code, w.Body.String())
	}
	var httpResp Response
	if err := json.Unmarshal(w.Body.Bytes(), &httpResp); err != nil {
		t.Fatalf("failed to decode HTTP response: %v", err)
	}
	resMap := httpResp.Result.(map[string]any)
	if resMap["isError"] == true {
		t.Fatalf("unexpected isError true on valid get_doc")
	}
	txt := resMap["content"].([]any)[0].(map[string]any)["text"].(string)
	var returnedPage DocPage
	if err := json.Unmarshal([]byte(txt), &returnedPage); err != nil {
		t.Fatalf("returned text is not valid DocPage JSON: %v", err)
	}
	validateDocPage(t, &returnedPage, "en")
}

func validateDocPage(t *testing.T, doc *DocPage, expectedLocale string) {
	t.Helper()
	if doc == nil {
		t.Fatal("doc is nil")
	}
	if doc.Path == "" {
		t.Errorf("empty doc.Path")
	}
	if doc.CleanPath == "" {
		t.Errorf("empty doc.CleanPath for %s", doc.Path)
	}
	if doc.Locale != expectedLocale {
		t.Errorf("expected locale %q, got %q", expectedLocale, doc.Locale)
	}
	if doc.Title == "" {
		t.Errorf("empty doc.Title for %s", doc.Path)
	}
	if doc.Content == "" {
		t.Errorf("empty doc.Content for %s", doc.Path)
	}
	if doc.URL == "" {
		t.Errorf("empty doc.URL for %s", doc.Path)
	}
}

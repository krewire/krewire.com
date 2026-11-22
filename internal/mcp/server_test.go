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
	if !strings.Contains(w.Body.String(), "krewire-docs-mcp") {
		t.Errorf("expected krewire-docs-mcp in health response: %s", w.Body.String())
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

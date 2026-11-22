package mcp

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// HTTPTransport coordinates SSE and HTTP message endpoints for the public MCP server.
type HTTPTransport struct {
	srv      *Server
	mu       sync.RWMutex
	sessions map[string]chan *Response
}

// NewHTTPTransport creates an HTTP transport handler for the MCP Server.
func NewHTTPTransport(srv *Server) *HTTPTransport {
	return &HTTPTransport{
		srv:      srv,
		sessions: make(map[string]chan *Response),
	}
}

// Mount attaches MCP routes (/sse, /message, /health) under prefix to the given mux.
func (h *HTTPTransport) Mount(mux *http.ServeMux, prefix string) {
	prefix = stringsTrimSlash(prefix)
	ssePath := prefix + "/sse"
	msgPath := prefix + "/message"
	healthPath := prefix + "/health"
	rootPath := prefix + "/"

	mux.HandleFunc(ssePath, h.handleSSE)
	mux.HandleFunc(msgPath, h.handleMessage)
	mux.HandleFunc(healthPath, h.handleHealth)
	mux.HandleFunc(rootPath, h.handleRoot)
}

func stringsTrimSlash(s string) string {
	for len(s) > 0 && s[len(s)-1] == '/' {
		s = s[:len(s)-1]
	}
	return s
}

func (h *HTTPTransport) registerSession(id string) chan *Response {
	ch := make(chan *Response, 32)
	h.mu.Lock()
	h.sessions[id] = ch
	h.mu.Unlock()
	return ch
}

func (h *HTTPTransport) unregisterSession(id string) {
	h.mu.Lock()
	if ch, ok := h.sessions[id]; ok {
		close(ch)
		delete(h.sessions, id)
	}
	h.mu.Unlock()
}

func (h *HTTPTransport) getSession(id string) (chan *Response, bool) {
	h.mu.RLock()
	ch, ok := h.sessions[id]
	h.mu.RUnlock()
	return ch, ok
}

func (h *HTTPTransport) handleSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	sessionID := generateSessionID()
	respCh := h.registerSession(sessionID)
	defer h.unregisterSession(sessionID)

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(http.StatusOK)

	// Send endpoint event for MCP clients
	endpointURI := fmt.Sprintf("/mcp/message?sessionId=%s", sessionID)
	fmt.Fprintf(w, "event: endpoint\ndata: %s\n\n", endpointURI)
	flusher.Flush()

	ctx := r.Context()
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// SSE keepalive ping
			fmt.Fprintf(w, ": ping\n\n")
			flusher.Flush()
		case resp, ok := <-respCh:
			if !ok {
				return
			}
			b, err := json.Marshal(resp)
			if err != nil {
				continue
			}
			fmt.Fprintf(w, "event: message\ndata: %s\n\n", string(b))
			flusher.Flush()
		}
	}
}

func (h *HTTPTransport) handleMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20)) // 1MB limit
	if err != nil {
		http.Error(w, "read error", http.StatusBadRequest)
		return
	}

	var req Request
	if err := json.Unmarshal(body, &req); err != nil {
		resp := &Response{
			JSONRPC: "2.0",
			Error:   &Error{Code: CodeParseError, Message: "parse error: invalid JSON"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	sessionID := r.URL.Query().Get("sessionId")
	ctx := r.Context()
	resp := h.srv.HandleRequest(ctx, req)

	// If session channel exists, dispatch asynchronously through SSE
	if sessionID != "" {
		if ch, ok := h.getSession(sessionID); ok {
			if resp != nil {
				select {
				case ch <- resp:
				default:
				}
			}
			w.WriteHeader(http.StatusAccepted)
			return
		}
	}

	// Stateless JSON-RPC fallback: return response directly in HTTP response
	if resp != nil {
		writeJSON(w, http.StatusOK, resp)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPTransport) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	uptime := time.Since(h.srv.startTime).Round(time.Second)
	status := map[string]any{
		"status":       "ok",
		"server":       "krewire-docs-mcp",
		"version":      "0.2.4",
		"uptime":       uptime.String(),
		"docs_indexed": h.srv.docs.Count(),
		"transport":    "sse+http",
		"endpoints": map[string]string{
			"sse":     "/mcp/sse",
			"message": "/mcp/message",
			"portal":  "/mcp",
		},
	}
	_ = json.NewEncoder(w).Encode(status)
}

func (h *HTTPTransport) handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/mcp" && r.URL.Path != "/mcp/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	meta := map[string]any{
		"name":        "krewire-docs-mcp",
		"description": "Official Model Context Protocol (MCP) server for Krewire documentation and ecosystem",
		"version":     "0.2.4",
		"protocol":    ProtocolVersion,
		"docs_url":    "https://krewire.com/mcp",
		"endpoints": map[string]string{
			"sse":     "https://krewire.com/mcp/sse",
			"message": "https://krewire.com/mcp/message",
			"health":  "https://krewire.com/mcp/health",
		},
	}
	_ = json.NewEncoder(w).Encode(meta)
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func generateSessionID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

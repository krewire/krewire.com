package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/krewire/krewire.com/mcp/mcp"
)

const version = "0.2.4"

func main() {
	var (
		portFlag    = flag.String("port", "8088", "HTTP port for SSE and JSON-RPC endpoints")
		hostFlag    = flag.String("host", "0.0.0.0", "Host address to bind HTTP server")
		docsDirFlag = flag.String("docs", "", "Path to content directory (default: auto-detected 'content')")
		stdioFlag   = flag.Bool("stdio", false, "Run in stdio mode for local IDEs (Cursor, Claude, Antigravity)")
		versionFlag = flag.Bool("version", false, "Print version and exit")
	)
	flag.Parse()

	if *versionFlag {
		fmt.Printf("krewire-mcp-server version %s (protocol %s)\n", version, mcp.ProtocolVersion)
		os.Exit(0)
	}

	docsDir := resolveDocsDir(*docsDirFlag)
	var docsIndexer *mcp.DocIndexer
	if docsDir != "" {
		idx, err := mcp.NewDocIndexer(os.DirFS(docsDir))
		if err != nil {
			log.Printf("warning: indexing docs from %s: %v", docsDir, err)
		} else {
			docsIndexer = idx
			log.Printf("indexed %d documentation documents from %s", docsIndexer.Count(), docsDir)
		}
	} else {
		log.Printf("warning: content directory not found, starting with empty doc index")
		docsIndexer, _ = mcp.NewDocIndexer(nil)
	}

	srv := mcp.NewServer(docsIndexer)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
	}()

	if *stdioFlag {
		stdioTransport := mcp.NewStdioTransport(srv, os.Stdin, os.Stdout)
		if err := stdioTransport.Serve(ctx); err != nil {
			log.Fatalf("stdio transport: %v", err)
		}
		return
	}

	// HTTP/SSE Transport
	mux := http.NewServeMux()
	httpTransport := mcp.NewHTTPTransport(srv)
	httpTransport.Mount(mux, "/mcp")

	addr := net.JoinHostPort(*hostFlag, *portFlag)
	httpServer := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, sCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer sCancel()
		_ = httpServer.Shutdown(shutdownCtx)
	}()

	log.Printf("Krewire MCP Server v%s listening on http://%s", version, addr)
	log.Printf("  SSE endpoint:    http://%s/mcp/sse", addr)
	log.Printf("  Message endpoint: http://%s/mcp/message", addr)
	log.Printf("  Health endpoint:  http://%s/mcp/health", addr)

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("http server failed: %v", err)
	}
}

func resolveDocsDir(custom string) string {
	if custom != "" {
		if fi, err := os.Stat(custom); err == nil && fi.IsDir() {
			return custom
		}
	}

	// Probe common locations relative to current working directory or binary
	candidates := []string{
		"content",
		filepath.Join("krewire.com", "content"),
		filepath.Join("..", "content"),
		"/usr/share/krewire/content",
		"/usr/share/nginx/html/content",
	}

	for _, cand := range candidates {
		if fi, err := os.Stat(cand); err == nil && fi.IsDir() {
			return cand
		}
	}
	return ""
}

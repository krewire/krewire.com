package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// StdioTransport handles line-delimited JSON-RPC over stdin/stdout for local IDEs.
type StdioTransport struct {
	srv *Server
	in  io.Reader
	out io.Writer
}

// NewStdioTransport creates a stdio transport for the given MCP server.
func NewStdioTransport(srv *Server, in io.Reader, out io.Writer) *StdioTransport {
	if in == nil {
		in = os.Stdin
	}
	if out == nil {
		out = os.Stdout
	}
	return &StdioTransport{srv: srv, in: in, out: out}
}

// Serve runs the stdio message loop until EOF or context cancellation.
func (t *StdioTransport) Serve(ctx context.Context) error {
	scanner := bufio.NewScanner(t.in)
	// Support up to 4MB JSON-RPC payloads
	buf := make([]byte, 64*1024)
	scanner.Buffer(buf, 4*1024*1024)

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		var req Request
		if err := json.Unmarshal(line, &req); err != nil {
			resp := &Response{
				JSONRPC: "2.0",
				Error:   &Error{Code: CodeParseError, Message: "parse error: invalid JSON"},
			}
			t.writeResponse(resp)
			continue
		}

		resp := t.srv.HandleRequest(ctx, req)
		if resp != nil {
			t.writeResponse(resp)
		}
	}

	return scanner.Err()
}

func (t *StdioTransport) writeResponse(resp *Response) {
	b, err := json.Marshal(resp)
	if err != nil {
		return
	}
	_, _ = fmt.Fprintln(t.out, string(b))
}

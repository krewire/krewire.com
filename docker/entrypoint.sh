#!/bin/sh
set -e

# Start Krewire MCP Server daemon in background
echo "Starting Krewire MCP Server daemon on 127.0.0.1:8088..."
/usr/local/bin/krewire-mcp-server --port 8088 --host 127.0.0.1 --docs /usr/share/nginx/html/content &

# Exec Nginx in foreground as PID 1
exec nginx -g "daemon off;"

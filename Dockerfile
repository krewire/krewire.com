# ── Stage 1: Build MCP Server binary ─────────────────────────────────────────
FROM golang:1.24-alpine AS builder

WORKDIR /src
COPY go.mod ./
COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /bin/krewire-mcp-server ./cmd/krewire-mcp-server

# ── Stage 2: Serve Nginx + MCP ───────────────────────────────────────────────
# Pin to digest to prevent unexpected image mutation on re-pull.
# To update: run `docker pull nginx:alpine` then update the digest below.
FROM nginx:alpine@sha256:df221db836e1754089190208cee7eeda94f233197056426eda74a43ab1abeac2

LABEL org.opencontainers.image.title="krewire-app" \
      org.opencontainers.image.description="Krewire documentation site and public MCP server" \
      org.opencontainers.image.source="https://github.com/krewire/krewire.com" \
      org.opencontainers.image.version="v0.2.4" \
      org.opencontainers.image.licenses="MIT"

# Remove default nginx config to avoid any default server leakage
RUN rm /etc/nginx/conf.d/default.conf

# Inject hardened nginx configuration and entrypoint
COPY docker/nginx.conf /etc/nginx/conf.d/krewire.conf
COPY docker/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

# Copy MCP server binary from builder
COPY --from=builder /bin/krewire-mcp-server /usr/local/bin/krewire-mcp-server

# Copy compiled static site artifacts and content tree
COPY .krewire/build /usr/share/nginx/html
COPY content /usr/share/nginx/html/content

EXPOSE 80

# Liveness check: nginx process is running and can serve root page & mcp health
HEALTHCHECK --interval=5s --timeout=3s --start-period=3s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://127.0.0.1/mcp/health || wget --quiet --tries=1 --spider http://127.0.0.1/ || exit 1

CMD ["/usr/local/bin/entrypoint.sh"]

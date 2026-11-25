# Production Deployment (`krewire.com`)

`krewire.com` runs directly on this host as a containerized static site paired with the embedded `krewire-mcp-server` daemon, fronted by Cloudflare Edge and a host Nginx reverse proxy.

---

## Network Topology

```text
Public User (Browser / MCP Client)
       │  HTTPS (Port 443)
       ▼
Cloudflare Edge Proxy (DNS Proxied · Orange Cloud · Flexible SSL)
       │  HTTP (Port 80) — Zero SSL Cert on Origin
       ▼
Host Nginx Reverse Proxy (Port 80)
   • /etc/nginx/sites-available/krewire.com
   • /etc/nginx/conf.d/cloudflare.conf (Restores real client IP & X-Forwarded-Proto)
       │  HTTP Proxy Pass (127.0.0.1:8090)
       ▼
Docker Container: krewire-app (127.0.0.1:8090:80)
   ├── Alpine Nginx (PID 1, Port 80) — Serves .krewire/build/ & proxies /mcp/*
   └── krewire-mcp-server (127.0.0.1:8088) — Serves /mcp/sse, /mcp/message, /mcp/health
```

---

## Invariant Deployment Rules

1. **Zero Origin SSL Certificate**:
   - Do **not** install ACME, Certbot, or local TLS certificates on the origin server for `krewire.com`.
   - TLS termination is handled entirely at the Cloudflare Edge (Flexible SSL).
   - Host Nginx listens on **port 80 only**.
2. **Container Identity & Port Binding**:
   - Defined in [`docker-compose.yml`](../docker-compose.yml): the container name must remain `krewire-app`, bound strictly to loopback `127.0.0.1:8090:80`.
3. **Container Internals ([`Dockerfile`](../Dockerfile), [`docker/entrypoint.sh`](../docker/entrypoint.sh), [`docker/nginx.conf`](../docker/nginx.conf))**:
   - Stage 1 (`golang:1.24-alpine`) compiles `krewire-mcp-server` with `CGO_ENABLED=0` and `-ldflags="-s -w"`.
   - Stage 2 (`nginx:alpine` pinned by digest) copies `.krewire/build` to `/usr/share/nginx/html` and `content/` to `/usr/share/nginx/html/content`.
   - `entrypoint.sh` starts `krewire-mcp-server --port 8088 --host 127.0.0.1 --docs /usr/share/nginx/html/content` in the background, then `exec`s `nginx -g "daemon off;"` as PID 1.
   - Inside the container, `docker/nginx.conf` sets security headers (`X-Content-Type-Options`, `X-Frame-Options`, `Content-Security-Policy`, `Strict-Transport-Security`, `Referrer-Policy`, `Permissions-Policy`), disables caching on HTML and `.sh` installer scripts, and proxies `^/mcp(/.*)?$` to `http://127.0.0.1:8088` with `proxy_buffering off` for SSE.

---

## Deployment Cycle

Always rebuild the static site with `krew build` before rebuilding the Docker image:

```bash
cd /home/krewire/workspace/krewire/krewire.com
krew build
docker compose build
docker compose up -d
```

---

## Health Verification

After deploying, verify the container status, host Nginx routing, container HTTP response, and MCP health endpoint:

```bash
docker ps --filter "name=krewire-app"
curl -sI -H "Host: krewire.com" http://127.0.0.1/
curl -sI http://127.0.0.1:8090/
curl -s http://127.0.0.1:8090/mcp/health
```

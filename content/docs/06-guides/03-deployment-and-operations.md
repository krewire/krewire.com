---
title: "Deployment & Operations"
description: "Production operations playbook: multi-stage Dockerfiles, Nginx reverse proxies, systemd unit files, and zero-downtime deployments."
date: "2026-10-08"
---

# Deployment & Operations

Krewire applications compile into single static binaries with embedded assets, making production operations simpler, faster, and more reliable than traditional multi-language web runtimes.

---

## 1. Multi-Stage Dockerfile

Because Krewire monoliths and services have zero runtime CGo or Node.js dependencies, production container images can be based on lightweight Alpine or scratch images:

```dockerfile
# Stage 1: Build binary
FROM golang:1.27-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/server ./cmd/server/main.go

# Stage 2: Minimal runtime image
FROM alpine:3.20

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/server /app/server

EXPOSE 8080
USER nobody:nobody

ENTRYPOINT ["/app/server"]
```

The resulting container image is typically **under 25MB** with zero operating system vulnerabilities.

---

## 2. Nginx Reverse Proxy Configuration

When hosting multiple applications behind Nginx on a Linux VPS:

```nginx
server {
    listen 80;
    server_name example.com;
    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl http2;
    server_name example.com;

    ssl_certificate /etc/letsencrypt/live/example.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/example.com/privkey.pem;

    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

---

## 3. Systemd Service Unit

To run a Krewire binary directly as a managed system daemon:

```ini
# /etc/systemd/system/krewire-app.service
[Unit]
Description=Krewire Production Web Application
After=network.target

[Service]
Type=simple
User=www-data
Group=www-data
WorkingDirectory=/var/www/app
ExecStart=/var/www/app/bin/server
Restart=always
RestartSec=5
Environment=KIW_ENV=production
Environment=KIW_PORT=8080

# Hardening
ProtectSystem=full
NoNewPrivileges=true

[Install]
WantedBy=multi-user.target
```

Enable and start the service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now krewire-app
```

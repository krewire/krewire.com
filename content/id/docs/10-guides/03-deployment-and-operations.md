---
title: "Deployment & Operasional"
description: "Playbook operasional produksi: Dockerfile multi-stage, reverse proxy Nginx, berkas unit systemd, dan deployment tanpa downtime (zero-downtime)."
date: "2026-10-08"
---

# Deployment & Operasional

Aplikasi Krewire dikompilasi menjadi binary statis tunggal dengan aset yang disematkan (embedded assets), menjadikan operasional produksi lebih sederhana, lebih cepat, dan lebih andal dibandingkan runtime web multi-bahasa konvensional.

---

## 1. Dockerfile Multi-Stage

Karena monolit dan service Krewire tidak memiliki dependensi runtime CGo maupun Node.js, image kontainer produksi dapat berbasis image Alpine atau scratch yang sangat ringan:

```dockerfile
# Stage 1: Build binary
FROM golang:1.27-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/server ./cmd/server/main.go

# Stage 2: Image runtime minimal
FROM alpine:3.20

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/server /app/server

EXPOSE 8080
USER nobody:nobody

ENTRYPOINT ["/app/server"]
```

Image kontainer yang dihasilkan biasanya berukuran **di bawah 25MB** dengan nol kerentanan sistem operasi.

---

## 2. Konfigurasi Reverse Proxy Nginx

Saat menampung beberapa aplikasi di balik Nginx pada Linux VPS:

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

## 3. Unit Service Systemd

Untuk menjalankan binary Krewire secara langsung sebagai daemon sistem terkelola:

```ini
# /etc/systemd/system/krewire-app.service
[Unit]
Description=Aplikasi Web Produksi Krewire
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
Environment=APP_ADDR=:8080

# Hardening
ProtectSystem=full
NoNewPrivileges=true

[Install]
WantedBy=multi-user.target
```

Aktifkan dan jalankan service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now krewire-app
```

---

## 4. Hosting Situs Statis & Portal Dokumentasi

Untuk workload `site` dan `book` yang di-build menggunakan `kiw build`, artefak statis dikompilasi ke dalam direktori `.krewire/build`.

### Virtual Host Nginx untuk Aset Statis

```nginx
server {
    server_name example.com;

    root /var/www/example.com;
    index index.html;

    location / {
        try_files $uri $uri/ $uri.html /index.html;
    }

    # Caching aset statis
    location ~* \.(css|js|png|jpg|jpeg|gif|ico|svg|woff|woff2|ttf|eot)$ {
        expires 7d;
        add_header Cache-Control "public, immutable";
    }

    listen 443 ssl http2;
    listen [::]:443 ssl http2;
    ssl_certificate /etc/letsencrypt/live/example.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/example.com/privkey.pem;
    include /etc/letsencrypt/options-ssl-nginx.conf;
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;
}

server {
    listen 80;
    listen [::]:80;
    server_name example.com;
    return 301 https://$host$request_uri;
}
```

### TLS Otomatis via Certbot

Menerbitkan sertifikat TLS otomatis menggunakan Certbot:

```bash
sudo certbot --nginx -d example.com --non-interactive --agree-tos --redirect -m admin@example.com
```

### Deployment CI/CD dengan GitHub Actions & Rsync

```yaml
- name: Deploy ke VPS
  run: |
    rsync -az --delete \
      -e "ssh -i ~/.ssh/deploy_key -o StrictHostKeyChecking=no" \
      .krewire/build/ \
      user@vps-host:/var/www/example.com/
```

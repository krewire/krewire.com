# Krewire product landing pages

Static, zero-JavaScript landing pages for the two Krewire products, deployed as
their own nginx containers and routed by the host edge.

- **Krewire Dash** → `https://dash.krewire.com/`
- **Krewire Cloud** → `https://cloud.krewire.com/`

Each host keeps its existing console: Dash at `/console`, Cloud at `/console`
(which redirects to `/login`). Only the exact root path is served by the landing
container, so API, WebSocket, and console routes continue to reach the daemons.

## Layout

```text
landings/
├── Dockerfile                    # one image per product, selected with --build-arg SITE
├── compose.yaml                  # landing-dash :8094, landing-cloud :8095 (loopback only)
├── shared/landing.css            # the single design-system owner
├── dash/index.html               # Dash copy
├── cloud/index.html              # Cloud copy
└── nginx/
    ├── container.conf            # nginx config inside each landing image
    ├── dash.krewire.com.conf     # host edge config (canonical source)
    └── cloud.krewire.com.conf    # host edge config (canonical source)
```

Design tokens, typography, cards, and layout live only in `shared/landing.css`;
a page switches accent through `<html data-accent="dash|cloud">`. Do not copy
styles into the HTML files.

## Deploy

```bash
# 1. Build and start the two landing containers (loopback only).
docker compose -f krewire.com/landings/compose.yaml up -d --build

# 2. Publish the edge configuration and reload nginx.
sudo install -m 0644 krewire.com/landings/nginx/dash.krewire.com.conf \
  /etc/nginx/sites-available/dash.krewire.com
sudo install -m 0644 krewire.com/landings/nginx/cloud.krewire.com.conf \
  /etc/nginx/sites-available/cloud.krewire.com
sudo ln -sfn /etc/nginx/sites-available/dash.krewire.com /etc/nginx/sites-enabled/dash.krewire.com
sudo ln -sfn /etc/nginx/sites-available/cloud.krewire.com /etc/nginx/sites-enabled/cloud.krewire.com
sudo nginx -t && sudo systemctl reload nginx
```

The containers bind to `127.0.0.1` only; nginx terminates TLS with the existing
`auth.krewire.com` certificate, whose SANs already include `dash.krewire.com`
and `cloud.krewire.com`.

## Verify

```bash
curl -fsS http://127.0.0.1:8094/_health        # ok
curl -fsS http://127.0.0.1:8095/_health        # ok

curl -fsS --resolve dash.krewire.com:443:127.0.0.1  https://dash.krewire.com/ | grep -o '<title>[^<]*'
curl -fsS --resolve cloud.krewire.com:443:127.0.0.1 https://cloud.krewire.com/ | grep -o '<title>[^<]*'

# Console routes must still reach the daemons.
curl -fsS -o /dev/null -w '%{http_code}\n' --resolve dash.krewire.com:443:127.0.0.1 https://dash.krewire.com/console
curl -fsS -o /dev/null -w '%{http_code}\n' --resolve cloud.krewire.com:443:127.0.0.1 https://cloud.krewire.com/console
```

## Maintenance

- Copy changes: edit `dash/index.html` or `cloud/index.html`.
- Design changes: edit `shared/landing.css`; both products pick it up on rebuild.
- Rebuild after any change: `docker compose -f krewire.com/landings/compose.yaml up -d --build`.
- The pages are accessible without JavaScript, honour `prefers-color-scheme`,
  and are constrained by a `script-src 'none'` CSP.

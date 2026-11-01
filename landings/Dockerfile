# syntax=docker/dockerfile:1
# ── Krewire product landing pages ─────────────────────────────────────────────
# Build one image per product by selecting the site directory:
#   docker build --build-arg SITE=dash  -t krewire-landing-dash  .
#   docker build --build-arg SITE=cloud -t krewire-landing-cloud .
# Both images ship the shared design system plus exactly one site.
FROM nginx:alpine@sha256:df221db836e1754089190208cee7eeda94f233197056426eda74a43ab1abeac2

ARG SITE=dash

LABEL org.opencontainers.image.title="krewire-landing-${SITE}" \
      org.opencontainers.image.description="Krewire ${SITE} product landing page" \
      org.opencontainers.image.source="https://github.com/krewire/krewire.com" \
      org.opencontainers.image.version="v0.1.0" \
      org.opencontainers.image.licenses="MIT"

RUN rm /etc/nginx/conf.d/default.conf

COPY nginx/container.conf /etc/nginx/conf.d/landing.conf
COPY shared/ /usr/share/nginx/html/shared/
COPY ${SITE}/ /usr/share/nginx/html/

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://127.0.0.1/_health || exit 1

CMD ["nginx", "-g", "daemon off;"]

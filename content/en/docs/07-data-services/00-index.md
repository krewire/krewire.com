---
title: "Data, Storage & Services"
description: "Reference catalog for Krewire's data persistence and backend services: multi-driver ORM, object storage, virtual filesystem, caching, mailer, and Indonesian payments."
date: "2026-10-25"
---

# Data, Storage & Services

Modern digital services require robust, modular storage and integration patterns. Krewire decouples data backends from business logic using driver-agnostic abstractions:

- **Pluggable Multi-Driver Architecture:** Switch between embedded, local, and distributed engines without changing application code.
- **Atomic Reliability:** Virtual filesystems and atomic writers ensure operations survive unexpected process termination or power loss.
- **Localization First:** Native support for Indonesian banking infrastructure (BCA, Mandiri, BRI, BNI) and national QRIS payments.

---

## What is in this Chapter?

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">7.1</span>
      <h3 class="chapter-card-title">
        <a href="/data-services/database-and-orm">Database & Multi-Driver ORM (pkg/orm) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Fluent query builder and multi-driver database layer supporting SQLite, MySQL, PostgreSQL, MongoDB, and Elasticsearch with automated migrations.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">7.2</span>
      <h3 class="chapter-card-title">
        <a href="/data-services/storage">Multi-Driver Object Storage (pkg/storage) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Unified object storage interface with swappable drivers for in-memory testing, local disk volumes, and S3-compatible cloud storage (MinIO, AWS, Cloudflare R2).
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">7.3</span>
      <h3 class="chapter-card-title">
        <a href="/data-services/filesystem">Virtual Filesystem & Atomic Files (pkg/fs, pkg/file) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Isolated, testable virtual filesystem abstractions (io/fs, memfs) and crash-safe atomic file writing primitives.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">7.4</span>
      <h3 class="chapter-card-title">
        <a href="/data-services/caching">Multi-Driver Caching (pkg/cache, pkg/redis) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      High-speed in-memory LRU cache and distributed Redis cache client with TTL expiration, atomic invalidation, and tag-based grouping.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">7.5</span>
      <h3 class="chapter-card-title">
        <a href="/data-services/mail">Transactional Mailer (pkg/mail) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Multi-driver transactional email dispatcher with built-in template rendering, attachment support, and integrations for SMTP, Resend, SendGrid, and Postmark.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">7.6</span>
      <h3 class="chapter-card-title">
        <a href="/data-services/commerce-and-payments">Commerce & Indonesian Payments (pkg/commerce) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      E-commerce ordering primitives, shopping cart workflows, and native Indonesian payment gateways: Midtrans, QRIS, Virtual Accounts, and e-wallets.
    </p>
  </div>

</div>

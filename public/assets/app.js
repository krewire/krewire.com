// =============================================================================
// Krewire — Core Application & Interactive Script (app.js)
// Zero external dependencies — pure fast vanilla JS
// =============================================================================

// ── 1. Theme Management (runs immediately in <head> to prevent FOUC) ─────────
(function () {
  'use strict';
  try {
    var stored = localStorage.getItem('krewire-theme') || 'auto';
    var mode = stored === 'auto'
      ? (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light')
      : stored;
    document.documentElement.dataset.theme = mode;
    document.documentElement.classList.toggle('dark', mode === 'dark');
  } catch (e) {}

  window.krewireTheme = {
    toggle: function () {
      var cur = document.documentElement.dataset.theme || (document.documentElement.classList.contains('dark') ? 'dark' : 'light');
      var nxt = cur === 'dark' ? 'light' : 'dark';
      document.documentElement.dataset.theme = nxt;
      document.documentElement.classList.toggle('dark', nxt === 'dark');
      try {
        localStorage.setItem('krewire-theme', nxt);
      } catch (e) {}
    }
  };
})();

// ── 2. Copy Command Helper ───────────────────────────────────────────────────
function copyCmd(btn, text) {
  if (!btn) return;
  var orig = btn.innerText;
  function showCopied() {
    btn.innerText = '✓ Copied!';
    setTimeout(function () { btn.innerText = orig; }, 2000);
  }

  if (navigator.clipboard && window.isSecureContext) {
    navigator.clipboard.writeText(text).then(showCopied).catch(function () {
      fallbackCopy(text, showCopied);
    });
  } else {
    fallbackCopy(text, showCopied);
  }
}

function fallbackCopy(text, cb) {
  try {
    var ta = document.createElement('textarea');
    ta.value = text;
    ta.style.position = 'fixed';
    ta.style.left = '-9999px';
    ta.style.top = '-9999px';
    document.body.appendChild(ta);
    ta.focus();
    ta.select();
    document.execCommand('copy');
    document.body.removeChild(ta);
    if (cb) cb();
  } catch (err) {}
}

window.copyCmd = copyCmd;

// ── 3. Forge Drawer & Mobile Menu Controller ────────────────────────────────
(function () {
  'use strict';

  function getDrawerElements(id) {
    var drawer = document.getElementById(id);
    var backdrop = document.getElementById(id + '-backdrop') || document.querySelector('[data-drawer-backdrop="' + id + '"]');
    var togglers = document.querySelectorAll('[data-drawer-toggle="' + id + '"], [aria-controls="' + id + '"]');
    return { drawer: drawer, backdrop: backdrop, togglers: togglers };
  }

  window.forgeDrawer = {
    open: function (id) {
      var els = getDrawerElements(id);
      if (!els.drawer) return;
      els.drawer.classList.add('drawer-open');
      if (els.backdrop) els.backdrop.classList.add('backdrop-open');
      els.togglers.forEach(function (btn) {
        btn.setAttribute('aria-expanded', 'true');
        btn.classList.add('is-active');
      });
      document.body.classList.add('drawer-active');
    },
    close: function (id) {
      var els = getDrawerElements(id);
      if (!els.drawer) return;
      els.drawer.classList.remove('drawer-open');
      if (els.backdrop) els.backdrop.classList.remove('backdrop-open');
      els.togglers.forEach(function (btn) {
        btn.setAttribute('aria-expanded', 'false');
        btn.classList.remove('is-active');
      });
      document.body.classList.remove('drawer-active');
    },
    toggle: function (id) {
      var els = getDrawerElements(id);
      if (!els.drawer) return;
      if (els.drawer.classList.contains('drawer-open')) {
        this.close(id);
      } else {
        this.open(id);
      }
    }
  };

  document.addEventListener('DOMContentLoaded', function () {
    // Event delegation for drawers
    document.addEventListener('click', function (e) {
      var toggleBtn = e.target.closest('[data-drawer-toggle]');
      if (toggleBtn) {
        var id = toggleBtn.getAttribute('data-drawer-toggle');
        if (id) {
          e.preventDefault();
          window.forgeDrawer.toggle(id);
          return;
        }
      }

      var closeBtn = e.target.closest('[data-drawer-close]');
      if (closeBtn) {
        var id = closeBtn.getAttribute('data-drawer-close');
        if (id) {
          e.preventDefault();
          window.forgeDrawer.close(id);
          return;
        }
      }

      var backdrop = e.target.closest('[data-drawer-backdrop]');
      if (backdrop) {
        var id = backdrop.getAttribute('data-drawer-backdrop');
        if (id) {
          e.preventDefault();
          window.forgeDrawer.close(id);
          return;
        }
      }

      // Close drawer when clicking any link inside drawer body
      var drawerLink = e.target.closest('.drawer a');
      if (drawerLink) {
        var drawer = drawerLink.closest('.drawer');
        if (drawer && drawer.id) {
          window.forgeDrawer.close(drawer.id);
        }
      }
    });

    // Close on Escape key
    document.addEventListener('keydown', function (e) {
      if (e.key === 'Escape') {
        document.querySelectorAll('.drawer.drawer-open').forEach(function (d) {
          window.forgeDrawer.close(d.id);
        });
      }
    });

    // Theme switches
    document.querySelectorAll('[data-theme-toggle]').forEach(function (btn) {
      btn.addEventListener('click', function () {
        if (window.krewireTheme && window.krewireTheme.toggle) {
          window.krewireTheme.toggle();
        }
      });
    });
  });
})();

// ── 4. i18n System (English & Indonesian) ────────────────────────────────────
(function () {
  'use strict';

  var translations = {
    "en": {
      "nav": {
        "get_started": "Get Started",
        "workloads": "8 Workloads",
        "dsl": "Component DSL",
        "docs": "Documentation",
        "ecosystem": "Ecosystem"
      },
      "hero": {
        "eyebrow": "◈ End-to-End Digital SDLC Ecosystem · v0.1.0 · Open Source",
        "title_html": "Modular Go Libraries.<br class=\"hidden sm:inline\">Every Workload.",
        "lead_html": "Krewire is an end-to-end <b class=\"text-fg dark:text-fg-dark\">digital SDLC ecosystem</b> — from the first line of the spec to production operations — built on three pillars: <b class=\"text-fg dark:text-fg-dark\">secure</b>, <b class=\"text-fg dark:text-fg-dark\">sustainable</b>, and <b class=\"text-fg dark:text-fg-dark\">scalable</b>. From single-binary web monoliths to interactive TUIs, high-speed static sites, book pipelines, background workers, microservices, cloud infra, and Go WebAssembly runtimes — build everything in Go with one CLI <code class=\"code-inline\">kiw</code> and zero JS fatigue.",
        "cta_start": "Get Started — 5 Min →",
        "cta_workloads": "Explore 8 Workloads ↓",
        "meta": "Go 1.27.1+ · Spec to Production · File-based routing · Scoped CSS · 100% Zero JS by default"
      },
      "trust": {
        "label": "Unified Go Toolchain",
        "stdlib": "100% Go Standard Library",
        "npm": "0 npm / 0 Bundlers",
        "workloads": "8 Workloads in 1 Binary",
        "coldstart": "< 1ms Cold Starts",
        "config": "Single Config krewire.yaml"
      },
      "pillars": {
        "badge": "Three Pillars",
        "title": "Secure. Sustainable. Scalable.",
        "lead": "Krewire is an end-to-end digital SDLC ecosystem — from the first line of the specification to running production. Every capability must pass all three pillars before it ships, and each claim traces to a spec, a test, or a gate.",
        "secure_title": "Secure",
        "secure_desc": "Security is a default, not a feature. A stdlib-first dependency graph, secure-by-default HTTP primitives, OWASP/CWE-aligned controls, secrets referenced — never stored — and WASM sandboxing keep the attack surface small and auditable.",
        "sustainable_title": "Sustainable",
        "sustainable_desc": "Software a single builder can still read, run, and afford a decade from now. Single static binaries with embedded assets, opt-in batteries that cost nothing when unused, boring proven parts, and no license fees — ever.",
        "scalable_title": "Scalable",
        "scalable_desc": "Growth is additive, not a rewrite. The progressive pipeline takes a product from static site to monolith to workers, services, and infra — each stage an opt-in, reversible battery, from a $5 VPS to hyperscale."
      },
      "workloads": {
        "badge": "Unified Workload Engine",
        "title": "Modular Libraries. Eight Workloads.",
        "lead": "Stop assembling disparate frameworks for web, CLI, background queues, and infra. Krewire provides unified ergonomics, shared configuration, and zero context switching.",
        "app_desc": "Fullstack web monolith — compile your web interface, HTTP routing, and API handlers into a single static Go binary.",
        "cli_desc": "Developer command-line tools & rich terminal interfaces powered by the lightweight TUI library with zero dependencies.",
        "site_desc": "High-performance static marketing sites, landing pages, and blogs built with the .kiw DSL and scoped CSS styling.",
        "book_desc": "Technical documentation, knowledge bases, and API references compiled with the standalone mdbind engine.",
        "worker_desc": "Asynchronous background queues, scheduled recurring tasks, and durable workers with graceful shutdowns.",
        "service_desc": "Headless microservices, JSON REST APIs, and gRPC endpoints engineered for high concurrency and low latency.",
        "infra_desc": "Automated bare-metal and VPS provisioning, TLS certificate issuance, and zero-downtime blue-green deployments.",
        "runtime_desc": "Reactive WebAssembly runtime compiling Go to client-side WASM with virtual DOM diffing and instant hot hydration."
      },
      "steps": {
        "badge": "Developer Experience",
        "title": "From Zero to Production in 4 Commands",
        "lead": "No boilerplate scaffolding, no package manager wrestling, no hidden magic.",
        "s1_desc": "Install the standalone kiw CLI tool via a single-line shell script or Go install.",
        "s2_desc": "Generate a production-ready project template for any workload in under one second.",
        "s3_desc": "Instant hot-reload local development server with file-based routing and scoped CSS updates.",
        "s4_desc": "Compile into a single static binary or static asset bundle and ship to GitHub Pages or Docker."
      },
      "why": {
        "badge": "Architecture & Philosophy",
        "title": "Go-First. Progressive. Zero Bloat.",
        "lead": "Built for engineers who value compilation speed, predictable deployments, and maintainable software architecture.",
        "w1_title": "Pure Go Toolchain",
        "w1_desc": "No Node.js, no npm dependencies, and no bundler configuration. Your entire application compiles with standard Go tooling into clean, self-contained binaries.",
        "w2_title": "Scoped .kiw Components",
        "w2_desc": "Combines YAML frontmatter, standard Go HTML templates, and scoped <style> blocks. CSS classes never leak outside their component boundary.",
        "w3_title": "Progressive Scaling",
        "w3_desc": "Start as a lightweight static site, grow into a fullstack web monolith, then seamlessly extract background workers and microservices — all within the same ecosystem.",
        "w4_title": "Enterprise Resilience",
        "w4_desc": "Engineered with production defaults: structured JSON logging, distributed tracing, health checks, circuit breakers, and graceful shutdown out of the box."
      },
      "dsl": {
        "badge": "Component DSL",
        "title": "Write Pages as Clean .kiw Components",
        "lead": "The elegance of Svelte and Astro brought to Go. Frontmatter + HTML + scoped CSS in one intuitive file.",
        "f1_html": "<b>Scoped by Default:</b> Styles remain strictly isolated to the component.",
        "f2_html": "<b>Global CSS via <code class=\"code-inline\">:root</code>:</b> Theme tokens cascade cleanly across layouts.",
        "f3_html": "<b>Built into <code class=\"code-inline\">kiw</code>:</b> Compiled natively via <code class=\"code-inline\">kiw/dsl</code> with zero npm or external bundler dependencies.",
        "f4_html": "<b>Light & Dark Mode:</b> Native token switching via <code class=\"code-inline\">data-theme</code>."
      },
      "matrix": {
        "badge": "Stack Comparison",
        "title": "Why Engineers Choose Krewire",
        "lead": "Compare the simplicity of an all-in-one Go toolchain against fragmented multi-language setups."
      },
      "cta": {
        "badge": "Ready to Build",
        "title": "Start Building with Krewire Today",
        "lead": "Start with a lightweight static site, scale to a web monolith, then to distributed cloud services — one repository, one CLI, one configuration.",
        "button": "Get Started — 5 Min →",
        "star": "★ Star on GitHub"
      }
    },
    "id": {
      "nav": {
        "get_started": "Mulai Cepat",
        "workloads": "8 Workload",
        "dsl": "DSL Komponen",
        "docs": "Dokumentasi",
        "ecosystem": "Ekosistem"
      },
      "hero": {
        "eyebrow": "◈ Ekosistem SDLC Digital End-to-End · v0.1.0 · Open Source",
        "title_html": "Library Go Modular.<br class=\"hidden sm:inline\">Untuk Setiap Workload.",
        "lead_html": "Krewire adalah <b class=\"text-fg dark:text-fg-dark\">ekosistem SDLC digital</b> end-to-end — dari baris pertama spesifikasi hingga operasi produksi — dibangun di atas tiga pilar: <b class=\"text-fg dark:text-fg-dark\">secure</b>, <b class=\"text-fg dark:text-fg-dark\">sustainable</b>, dan <b class=\"text-fg dark:text-fg-dark\">scalable</b>. Dari monolit web binary tunggal hingga TUI interaktif, situs statis berkecepatan tinggi, pipeline buku, worker latar belakang, microservices, cloud infra, dan runtime Go WebAssembly — bangun semuanya dalam Go dengan satu CLI <code class=\"code-inline\">kiw</code> dan bebas kelelahan JS.",
        "cta_start": "Mulai Sekarang — 5 Menit →",
        "cta_workloads": "Jelajahi 8 Workload ↓",
        "meta": "Go 1.27.1+ · Dari Spek ke Produksi · Routing Berbasis Berkas · Scoped CSS · 100% Bebas JS secara Bawaan"
      },
      "trust": {
        "label": "Toolchain Go Terpadu",
        "stdlib": "100% Standard Library Go",
        "npm": "0 npm / 0 Bundler",
        "workloads": "8 Workload dalam 1 Binary",
        "coldstart": "< 1ms Cold Start",
        "config": "Satu Konfigurasi krewire.yaml"
      },
      "pillars": {
        "badge": "Tiga Pilar",
        "title": "Secure. Sustainable. Scalable.",
        "lead": "Krewire adalah ekosistem SDLC digital end-to-end — dari baris pertama spesifikasi hingga sistem produksi yang berjalan. Setiap kapabilitas wajib lolos ketiga pilar sebelum dirilis, dan setiap klaim terlacak ke spek, pengujian, atau gerbang kualitas.",
        "secure_title": "Secure",
        "secure_desc": "Keamanan adalah bawaan, bukan sekadar fitur. Graf dependensi berbasis stdlib, primitif HTTP aman secara bawaan, kontrol selaras OWASP/CWE, kredensial hanya direferensikan — tidak pernah disimpan — serta sandboxing WASM menjaga permukaan serangan tetap minim dan teruji.",
        "sustainable_title": "Sustainable",
        "sustainable_desc": "Perangkat lunak yang dapat dibaca, dijalankan, dan dijangkau oleh satu pengembang bahkan satu dekade mendatang. Binary statis tunggal dengan aset tersemat, baterai opsional tanpa beban saat tak digunakan, komponen teruji yang stabil, dan bebas biaya lisensi — selamanya.",
        "scalable_title": "Scalable",
        "scalable_desc": "Pertumbuhan bersifat aditif, bukan penulisan ulang. Pipeline progresif membawa produk dari situs statis ke monolit hingga worker, service, dan infra — setiap tahap adalah modul opsional dan terbalikkan, dari VPS $5 hingga hyperscale."
      },
      "workloads": {
        "badge": "Engine Workload Terpadu",
        "title": "Library Modular. Delapan Workload.",
        "lead": "Berhenti merakit framework terpisah-pisah untuk web, CLI, antrean latar belakang, dan infra. Krewire menghadirkan ergonomi terpadu, konfigurasi bersama, dan tanpa peralihan konteks.",
        "app_desc": "Monolit web fullstack — kompilasi antarmuka web, routing HTTP, dan handler API menjadi satu binary statis Go mandiri.",
        "cli_desc": "Tool command-line pengembang & antarmuka terminal interaktif didukung pustaka TUI ringan tanpa dependensi luar.",
        "site_desc": "Situs pemasaran statis berkinerja tinggi, landing page, dan blog yang dibangun dengan DSL .kiw dan scoped CSS.",
        "book_desc": "Dokumentasi teknis, pusat panduan, dan referensi API yang dikompilasi menggunakan engine mdbind mandiri.",
        "worker_desc": "Antrean background asinkron, penjadwalan tugas berkala, dan worker tangguh dengan penghentian anggun (graceful shutdown).",
        "service_desc": "Microservice headless, REST API JSON, dan endpoint gRPC yang dirancang untuk konkurensi tinggi dan latensi rendah.",
        "infra_desc": "Penyediaan otomatis bare-metal & VPS, penerbitan sertifikat TLS, dan deployment blue-green tanpa downtime.",
        "runtime_desc": "Runtime WebAssembly reaktif yang mengompilasi Go ke client-side WASM dengan diffing virtual DOM dan hidrasi cepat."
      },
      "steps": {
        "badge": "Pengalaman Pengembang",
        "title": "Dari Nol ke Produksi dalam 4 Perintah",
        "lead": "Tanpa boilerplate rumit, tanpa pergulatan package manager, tanpa magis tersembunyi.",
        "s1_desc": "Pasang tool CLI kiw mandiri lewat skrip satu baris atau go install.",
        "s2_desc": "Buat template proyek siap produksi untuk workload apa pun dalam waktu kurang dari satu detik.",
        "s3_desc": "Server pengembangan lokal hot-reload instan dengan perutean berbasis berkas dan pembaruan scoped CSS.",
        "s4_desc": "Kompilasi ke binary statis tunggal atau bundle aset statis dan rilis ke GitHub Pages atau Docker."
      },
      "why": {
        "badge": "Arsitektur & Filosofi",
        "title": "Go-First. Progresif. Bebas Kembung.",
        "lead": "Dibangun untuk para engineer yang menghargai kecepatan kompilasi, deployment terprediksi, dan arsitektur kode yang terawat.",
        "w1_title": "Toolchain Go Murni",
        "w1_desc": "Tanpa Node.js, tanpa dependensi npm, dan tanpa konfigurasi bundler. Seluruh aplikasi Anda dikompilasi dengan toolchain Go standar menjadi biner mandiri yang bersih.",
        "w2_title": "Komponen .kiw yang Terisolasi",
        "w2_desc": "Menggabungkan frontmatter YAML, template HTML Go standar, dan blok <style> terisolasi. Kelas CSS tidak pernah bocor ke luar batas komponen.",
        "w3_title": "Skalabilitas Progresif",
        "w3_desc": "Mulai sebagai situs statis ringan, berkembang menjadi monolit web fullstack, lalu ekstrak worker antrean dan layanan mikro tanpa repot — semuanya dalam satu ekosistem.",
        "w4_title": "Ketahanan Skala Enterprise",
        "w4_desc": "Dirancang dengan standar siap produksi: structured JSON logging, distributed tracing, health check, circuit breaker, dan graceful shutdown langsung tersedia."
      },
      "dsl": {
        "badge": "DSL Komponen",
        "title": "Tulis Halaman sebagai Komponen .kiw yang Bersih",
        "lead": "Keanggunan Svelte dan Astro dihadirkan ke Go. Frontmatter + HTML + scoped CSS dalam satu berkas intuitif.",
        "f1_html": "<b>Terisolasi secara Bawaan:</b> Gaya tetap terisolasi ketat di dalam komponen.",
        "f2_html": "<b>CSS Global via <code class=\"code-inline\">:root</code>:</b> Token tema mengalir secara rapi di seluruh layout.",
        "f3_html": "<b>Bawaan dari <code class=\"code-inline\">kiw</code>:</b> Dikompilasi secara native melalui <code class=\"code-inline\">kiw/dsl</code> tanpa npm atau dependensi bundler eksternal.",
        "f4_html": "<b>Mode Terang & Gelap:</b> Pergantian token tema native melalui atribut <code class=\"code-inline\">data-theme</code>."
      },
      "matrix": {
        "badge": "Perbandingan Stack",
        "title": "Mengapa Engineer Memilih Krewire",
        "lead": "Bandingkan kesederhanaan toolchain Go terpadu all-in-one dengan tumpukan multi-bahasa yang terfragmentasi."
      },
      "cta": {
        "badge": "Siap Membangun",
        "title": "Mulai Bangun dengan Krewire Hari Ini",
        "lead": "Mulai dari situs statis ringan, skalakan ke monolit web, lalu ke layanan cloud terdistribusi — satu repositori, satu CLI, satu konfigurasi.",
        "button": "Mulai Sekarang — 5 Menit →",
        "star": "★ Star di GitHub"
      }
    }
  };

  function getNestedValue(obj, keyPath) {
    if (!obj || !keyPath) return null;
    var parts = keyPath.split('.');
    var curr = obj;
    for (var i = 0; i < parts.length; i++) {
      if (curr && typeof curr === 'object' && parts[i] in curr) {
        curr = curr[parts[i]];
      } else {
        return null;
      }
    }
    return curr;
  }

  function initI18n() {
    var currentLang = 'en';
    try {
      var saved = localStorage.getItem('krewire-lang');
      if (saved === 'id' || saved === 'en') {
        currentLang = saved;
      } else if (navigator.language && navigator.language.toLowerCase().startsWith('id')) {
        currentLang = 'id';
      }
    } catch (e) {}

    function applyLanguage(lang) {
      currentLang = (lang === 'id') ? 'id' : 'en';
      document.documentElement.lang = currentLang;
      try {
        localStorage.setItem('krewire-lang', currentLang);
      } catch (e) {}

      // Update all toggle buttons
      document.querySelectorAll('.lang-switch .lang-code, [data-lang-toggle] .lang-code').forEach(function (el) {
        el.textContent = currentLang.toUpperCase();
      });

      // Update text nodes
      document.querySelectorAll('[data-i18n]').forEach(function (el) {
        var key = el.getAttribute('data-i18n');
        var val = getNestedValue(translations[currentLang], key) || getNestedValue(translations.en, key);
        if (val) {
          if (el.tagName === 'INPUT' || el.tagName === 'TEXTAREA') {
            el.setAttribute('placeholder', val);
          } else {
            el.textContent = val;
          }
        }
      });

      // Update HTML nodes
      document.querySelectorAll('[data-i18n-html]').forEach(function (el) {
        var key = el.getAttribute('data-i18n-html');
        var val = getNestedValue(translations[currentLang], key) || getNestedValue(translations.en, key);
        if (val) {
          el.innerHTML = val;
        }
      });
    }

    window.krewireI18n = {
      getLocale: function () { return currentLang; },
      setLocale: function (l) { applyLanguage(l); },
      toggle: function () {
        applyLanguage(currentLang === 'en' ? 'id' : 'en');
      }
    };

    applyLanguage(currentLang);

    // Event listener for lang switch clicks
    document.addEventListener('click', function (e) {
      var btn = e.target.closest('[data-lang-toggle], .lang-switch');
      if (btn) {
        e.preventDefault();
        window.krewireI18n.toggle();
      }
    });
  }

  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', initI18n);
  } else {
    initI18n();
  }
})();

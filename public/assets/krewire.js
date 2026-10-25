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
get: function () {
return document.documentElement.dataset.theme || (document.documentElement.classList.contains('dark') ? 'dark' : 'light');
},
set: function (mode) {
document.documentElement.dataset.theme = mode;
document.documentElement.classList.toggle('dark', mode === 'dark');
try {
localStorage.setItem('krewire-theme', mode);
} catch (e) {}
},
toggle: function () {
var cur = window.krewireTheme.get();
var nxt = cur === 'dark' ? 'light' : 'dark';
window.krewireTheme.set(nxt);
}
};
})();
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
(function () {
'use strict';
function getDrawerElements(id) {
if (!id) return { drawer: null, backdrop: null, togglers: [] };
var drawer = document.getElementById(id) || document.querySelector('.' + id) || document.querySelector('[data-drawer="' + id + '"]');
var backdrop = document.getElementById(id + '-backdrop') || document.querySelector('[data-drawer-backdrop="' + id + '"]') || document.querySelector('.drawer-backdrop');
var togglers = document.querySelectorAll('[data-drawer-toggle="' + id + '"], [aria-controls="' + id + '"], #' + id + '-toggle');
return { drawer: drawer, backdrop: backdrop, togglers: togglers };
}
var drawerCtrl = {
open: function (id) {
var els = getDrawerElements(id);
if (!els.drawer) return;
els.drawer.classList.add('drawer-open');
els.drawer.classList.add('is-open');
els.drawer.setAttribute('aria-hidden', 'false');
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
els.drawer.classList.remove('is-open');
els.drawer.setAttribute('aria-hidden', 'true');
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
if (els.drawer.classList.contains('drawer-open') || els.drawer.classList.contains('is-open')) {
this.close(id);
} else {
this.open(id);
}
}
};
window.uiDrawer = drawerCtrl;
window.uiToggleDrawer = function (btn, id) {
if (!id && btn) {
id = btn.getAttribute('data-drawer-toggle') || btn.getAttribute('aria-controls');
}
if (id) {
window.uiDrawer.toggle(id);
}
};
window.uiCloseDrawer = function (id) {
if (id) {
window.uiDrawer.close(id);
}
};
function bindDrawerEvents() {
document.addEventListener('click', function (e) {
var toggleBtn = e.target.closest('[data-drawer-toggle], .nav-toggler, [aria-controls]');
if (toggleBtn) {
var id = toggleBtn.getAttribute('data-drawer-toggle') || toggleBtn.getAttribute('aria-controls');
if (id && (document.getElementById(id) || document.querySelector('.' + id))) {
e.preventDefault();
window.uiDrawer.toggle(id);
return;
}
}
var closeBtn = e.target.closest('[data-drawer-close]');
if (closeBtn) {
var id = closeBtn.getAttribute('data-drawer-close');
if (id) {
e.preventDefault();
window.uiDrawer.close(id);
return;
}
}
var backdrop = e.target.closest('[data-drawer-backdrop]');
if (backdrop) {
var id = backdrop.getAttribute('data-drawer-backdrop');
if (id) {
e.preventDefault();
window.uiDrawer.close(id);
return;
}
}
var drawerLink = e.target.closest('.drawer a');
if (drawerLink) {
var drawer = drawerLink.closest('.drawer');
if (drawer && drawer.id) {
window.uiDrawer.close(drawer.id);
}
}
});
document.addEventListener('keydown', function (e) {
if (e.key === 'Escape') {
document.querySelectorAll('.drawer.drawer-open').forEach(function (d) {
window.uiDrawer.close(d.id);
});
}
});
if (!window._krewireThemeBound) {
window._krewireThemeBound = true;
document.addEventListener('click', function (e) {
var btn = e.target && e.target.closest ? e.target.closest('[data-theme-toggle]') : null;
if (btn && window.krewireTheme && window.krewireTheme.toggle) {
e.preventDefault();
window.krewireTheme.toggle();
}
});
}
}
if (document.readyState === 'loading') {
document.addEventListener('DOMContentLoaded', bindDrawerEvents);
} else {
bindDrawerEvents();
}
})();
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
"eyebrow": "End-to-End Digital SDLC Ecosystem · v0.1.0 · Open Source",
"title_html": "Modular Go Libraries.<br class=\"hidden sm:inline\">Every Workload.",
"lead_html": "Krewire is an end-to-end <b class=\"text-fg dark:text-fg-dark\">digital SDLC ecosystem</b> — from the first line of the spec to production operations — built on three pillars: <b class=\"text-fg dark:text-fg-dark\">secure</b>, <b class=\"text-fg dark:text-fg-dark\">sustainable</b>, and <b class=\"text-fg dark:text-fg-dark\">scalable</b>. From single-binary web monoliths to interactive TUIs, high-speed static sites, book pipelines, background workers, microservices, cloud infra, and Go WebAssembly runtimes — build everything in Go with one CLI <code class=\"code-inline\">kiw</code> and zero JS fatigue.",
"cta_start": "Get Started — 5 Min →",
"cta_contribute": "Contribute →",
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
"lead": "Stop assembling disparate tools for web, CLI, background queues, and infra. Krewire provides unified ergonomics, shared configuration, and zero context switching.",
"app_desc": "Fullstack web monolith — compile your web interface, HTTP routing, and API handlers into a single static Go binary.",
"cli_desc": "Developer command-line tools & rich terminal interfaces powered by the lightweight TUI library with zero dependencies.",
"site_desc": "High-performance static marketing sites, landing pages, and blogs built with the .kiw DSL and scoped CSS styling.",
"book_desc": "Technical documentation, knowledge bases, and API references compiled with the native pkg/book engine.",
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
"lead": "Compare the simplicity of an all-in-one Go toolchain against fragmented multi-language setups.",
"th_dimension": "Architecture Dimension",
"th_traditional": "Traditional Fragmented Stack",
"th_krewire": "Krewire Unified Go",
"row1_dim": "Toolchain & Runtime",
"row1_trad": "2–3 Languages (Go + Node.js/TS + Python/HCL)",
"row1_krewire": "100% Pure Go across all workloads",
"row2_dim": "Package Management",
"row2_trad": "go.mod + package.json + npm/pnpm trees",
"row2_krewire": "Single go.mod (Zero npm / Zero JS)",
"row3_dim": "Developer CLI",
"row3_trad": "Multiple: Air + npm + Cobra + Terraform",
"row3_krewire": "One binary CLI: kiw",
"row4_dim": "Configuration",
"row4_trad": "Scattered (.air.toml, vite.config, tsconfig, Dockerfile)",
"row4_krewire": "Single unified krewire.yaml",
"row5_dim": "Component Styling",
"row5_trad": "Tailwind CLI / PostCSS bundler pipeline",
"row5_krewire": "Native scoped .kiw components",
"row6_dim": "Production Deployment",
"row6_trad": "Heavy multi-stage Docker containers with Node layers",
"row6_krewire": "Single self-contained Go binary or static assets"
},
"ecosystem": {
"badge": "Official Ecosystem",
"title": "4 Git Repositories — Unified Architecture",
"lead": "The entire Krewire ecosystem is organized into exactly 4 focused repositories that compose cleanly.",
"core_badge": "Core Monorepo",
"core_desc": "Core packages, apps, devtool CLI (kiw), and Boost agent templates in Go.",
"book_badge": "Book Engine",
"book_desc": "Standalone Markdown book and docs site builder powering the book workload.",
"portal_badge": "Portal Site",
"portal_desc": "Official production website and documentation portal built with site kind & .kiw DSL.",
"community_badge": "Community",
"community_desc": "Community portal for contributors, releases, guides, and sponsorships."
},
"cta": {
"badge": "Ready to Build",
"title": "Start Building with Krewire Today",
"lead": "Start with a lightweight static site, scale to a web monolith, then to distributed cloud services — one repository, one CLI, one configuration.",
"button": "Get Started — 5 Min →",
"star": "★ Star on GitHub"
},
"footer": {
"workloads_title": "8 Workloads",
"docs_title": "Documentation",
"ecosystem_title": "Ecosystem",
"tagline_html": "Modular Go Libraries for Every Workload.<br>One CLI <code class=\"code-inline\">kiw</code>, one config <code class=\"code-inline\">krewire.yaml</code>.",
"badge": "Compiled Go · Zero Bloat",
"copyright_html": "© 2026 Krewire Contributors — 100% built with Krewire <code class=\"code-inline\">site</code> workload."
},
"docs_reader": {
"toc_title": "Contents",
"search_placeholder": "Filter chapters...",
"copy_button": "Copy",
"copied_button": "✓ Copied!",
"prev_label": "Previous",
"next_label": "Next"
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
"eyebrow": "Ekosistem SDLC Digital End-to-End · v0.1.0 · Open Source",
"title_html": "Library Go Modular.<br class=\"hidden sm:inline\">Untuk Setiap Workload.",
"lead_html": "Krewire adalah <b class=\"text-fg dark:text-fg-dark\">ekosistem SDLC digital</b> end-to-end — dari baris pertama spesifikasi hingga operasi produksi — dibangun di atas tiga pilar: <b class=\"text-fg dark:text-fg-dark\">secure</b>, <b class=\"text-fg dark:text-fg-dark\">sustainable</b>, dan <b class=\"text-fg dark:text-fg-dark\">scalable</b>. Dari monolit web binary tunggal hingga TUI interaktif, situs statis berkecepatan tinggi, pipeline buku, worker latar belakang, microservices, cloud infra, dan runtime Go WebAssembly — bangun semuanya dalam Go dengan satu CLI <code class=\"code-inline\">kiw</code> dan bebas kelelahan JS.",
"cta_start": "Mulai Sekarang — 5 Menit →",
"cta_contribute": "Berkontribusi →",
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
"lead": "Berhenti merakit berbagai perkakas terpisah untuk web, CLI, antrean latar belakang, dan infra. Krewire menghadirkan ergonomi terpadu, konfigurasi bersama, dan tanpa peralihan konteks.",
"app_desc": "Monolit web fullstack — kompilasi antarmuka web, routing HTTP, dan handler API menjadi satu binary statis Go mandiri.",
"cli_desc": "Tool command-line pengembang & antarmuka terminal interaktif didukung pustaka TUI ringan tanpa dependensi luar.",
"site_desc": "Situs pemasaran statis berkinerja tinggi, landing page, dan blog yang dibangun dengan DSL .kiw dan scoped CSS.",
"book_desc": "Dokumentasi teknis, pusat panduan, dan referensi API yang dikompilasi menggunakan engine pkg/book bawaan.",
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
"lead": "Bandingkan kesederhanaan toolchain Go terpadu all-in-one dengan tumpukan multi-bahasa yang terfragmentasi.",
"th_dimension": "Dimensi Arsitektur",
"th_traditional": "Stack Tradisional Terfragmentasi",
"th_krewire": "Krewire Go Terpadu",
"row1_dim": "Toolchain & Runtime",
"row1_trad": "2–3 Bahasa (Go + Node.js/TS + Python/HCL)",
"row1_krewire": "100% Go Murni di semua workload",
"row2_dim": "Manajemen Paket",
"row2_trad": "go.mod + package.json + pohon npm/pnpm",
"row2_krewire": "Satu go.mod (Nol npm / Nol JS)",
"row3_dim": "CLI Pengembang",
"row3_trad": "Banyak tool: Air + npm + Cobra + Terraform",
"row3_krewire": "Satu binary CLI: kiw",
"row4_dim": "Konfigurasi",
"row4_trad": "Tersebar (.air.toml, vite.config, tsconfig, Dockerfile)",
"row4_krewire": "Satu krewire.yaml terpadu",
"row5_dim": "Styling Komponen",
"row5_trad": "Pipeline bundler Tailwind CLI / PostCSS",
"row5_krewire": "Komponen .kiw terisolasi native",
"row6_dim": "Deployment Produksi",
"row6_trad": "Container Docker bertingkat berat dengan layer Node",
"row6_krewire": "Binary Go mandiri tunggal atau aset statis"
},
"ecosystem": {
"badge": "Ekosistem Resmi",
"title": "4 Repositori Git — Arsitektur Terpadu",
"lead": "Seluruh ekosistem Krewire diorganisasi ke dalam tepat 4 repositori terarah yang terkomposisi bersih.",
"core_badge": "Monorepo Inti",
"core_desc": "Paket inti, aplikasi, developer CLI (kiw), dan template agent Boost dalam Go.",
"book_badge": "Engine Buku",
"book_desc": "Compiler buku Markdown & situs dokumentasi mandiri yang menggerakkan workload book.",
"portal_badge": "Situs Portal",
"portal_desc": "Situs produksi resmi dan portal dokumentasi yang dibangun dengan workload site & DSL .kiw.",
"community_badge": "Komunitas",
"community_desc": "Portal komunitas untuk kontributor, rilis, panduan, dan sponsor."
},
"cta": {
"badge": "Siap Membangun",
"title": "Mulai Bangun dengan Krewire Hari Ini",
"lead": "Mulai dari situs statis ringan, skalakan ke monolit web, lalu ke layanan cloud terdistribusi — satu repositori, satu CLI, satu konfigurasi.",
"button": "Mulai Sekarang — 5 Menit →",
"star": "★ Star di GitHub"
},
"footer": {
"workloads_title": "8 Workload",
"docs_title": "Dokumentasi",
"ecosystem_title": "Ekosistem",
"tagline_html": "Library Go Modular untuk Setiap Workload.<br>Satu CLI <code class=\"code-inline\">kiw</code>, satu konfigurasi <code class=\"code-inline\">krewire.yaml</code>.",
"badge": "Go Terkompilasi · Bebas Kembung",
"copyright_html": "© 2026 Kontributor Krewire — 100% dibangun dengan workload <code class=\"code-inline\">site</code> Krewire."
},
"docs_reader": {
"toc_title": "Daftar Isi",
"search_placeholder": "Cari bab...",
"copy_button": "Salin",
"copied_button": "✓ Tersalin!",
"prev_label": "Sebelumnya",
"next_label": "Selanjutnya"
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
function isDocsRoute(pathname) {
var p = pathname || (typeof window !== 'undefined' && window.location ? window.location.pathname : '');
return p.indexOf('/docs') !== -1 || p.startsWith('/en/docs') || p.startsWith('/id/docs');
}
function initI18n() {
var path = window.location.pathname;
var inDocs = isDocsRoute(path) || !!document.querySelector('.docs-reader, .chapter, #sidebar-toc');
var currentLang = 'en';
try {
if (inDocs) {
if (path.startsWith('/id/') || path === '/id') {
currentLang = 'id';
} else {
currentLang = 'en';
}
} else {
var saved = localStorage.getItem('krewire-lang');
if (saved === 'id' || saved === 'en') {
currentLang = saved;
} else if (navigator.language && navigator.language.toLowerCase().startsWith('id')) {
currentLang = 'id';
}
}
} catch (e) {}
function applyLanguage(lang) {
currentLang = (lang === 'id') ? 'id' : 'en';
document.documentElement.lang = currentLang;
try {
localStorage.setItem('krewire-lang', currentLang);
} catch (e) {}
document.querySelectorAll('.lang-switch .lang-code, [data-lang-toggle] .lang-code').forEach(function (el) {
el.textContent = currentLang.toUpperCase();
});
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
document.querySelectorAll('[data-i18n-html]').forEach(function (el) {
var key = el.getAttribute('data-i18n-html');
var val = getNestedValue(translations[currentLang], key) || getNestedValue(translations.en, key);
if (val) {
el.innerHTML = val;
}
});
try {
window.dispatchEvent(new CustomEvent('krewire:langchange', { detail: { lang: currentLang } }));
} catch (e) {}
}
window.krewireI18n = {
getLocale: function () { return currentLang; },
setLocale: function (l) { applyLanguage(l); },
toggle: function () {
applyLanguage(currentLang === 'en' ? 'id' : 'en');
},
t: function (key) {
return getNestedValue(translations[currentLang], key) || getNestedValue(translations.en, key) || '';
}
};
applyLanguage(currentLang);
if (!window._krewireLangBound) {
window._krewireLangBound = true;
document.addEventListener('click', function (e) {
var btn = e.target && e.target.closest ? e.target.closest('[data-lang-toggle], .lang-switch') : null;
if (btn) {
e.preventDefault();
var p = window.location.pathname;
var isDocs = isDocsRoute(p) || !!document.querySelector('.docs-reader, .chapter, #sidebar-toc');

if (isDocs) {
// Mechanism 2: Docs URL Path-based routing for markdown content
var isIdDocs = p.startsWith('/id/') || p === '/id';
var nextLang = isIdDocs ? 'en' : 'id';
try {
localStorage.setItem('krewire-lang', nextLang);
} catch (_) {}

var target;
if (nextLang === 'id') {
var clean = p.replace(/^\/en(\/|$)/, '/');
target = '/id' + (clean.startsWith('/') ? clean : '/' + clean);
} else {
target = p.replace(/^\/id(\/|$)/, '/') || '/docs';
}
var searchAndHash = (window.location.search || '') + (window.location.hash || '');
window.location.href = target + searchAndHash;
return;
}

// Mechanism 1: Landing / Site shell in-place DOM translation (<locale>.json)
var next = currentLang === 'en' ? 'id' : 'en';
applyLanguage(next);
}
});
}
}
if (document.readyState === 'loading') {
document.addEventListener('DOMContentLoaded', initI18n);
} else {
initI18n();
}
})();
;
(function () {
'use strict';
function getCopyText() {
if (window.krewireI18n && window.krewireI18n.t) {
return window.krewireI18n.t('docs_reader.copy_button') || 'Copy';
}
return 'Copy';
}
function getCopiedText() {
if (window.krewireI18n && window.krewireI18n.t) {
return window.krewireI18n.t('docs_reader.copied_button') || '✓ Copied!';
}
return '✓ Copied!';
}
function initCodeBlocks() {
document.querySelectorAll('.chapter pre').forEach(function (pre) {
if (pre.querySelector('.code-copy-btn')) return;
var btn = document.createElement('button');
btn.className = 'code-copy-btn';
btn.setAttribute('aria-label', 'Copy code to clipboard');
btn.innerHTML = '<span class="copy-txt">' + getCopyText() + '</span>';
btn.addEventListener('click', function () {
var code = pre.querySelector('code');
var text = (code ? code.innerText : pre.innerText).replace(/\n$/, '');
function markCopied() {
btn.innerHTML = '<span class="copy-txt copied">' + getCopiedText() + '</span>';
setTimeout(function () {
btn.innerHTML = '<span class="copy-txt">' + getCopyText() + '</span>';
}, 2000);
}
if (navigator.clipboard && navigator.clipboard.writeText) {
navigator.clipboard.writeText(text).then(markCopied);
} else {
try {
var ta = document.createElement('textarea');
ta.value = text;
ta.style.position = 'fixed';
ta.style.opacity = '0';
document.body.appendChild(ta);
ta.select();
document.execCommand('copy');
document.body.removeChild(ta);
markCopied();
} catch (e) {}
}
});
pre.appendChild(btn);
});
}
function initHeadingAnchors() {
document.querySelectorAll('.chapter h2, .chapter h3').forEach(function (h) {
if (!h.id || h.querySelector('.anchor-link')) return;
var a = document.createElement('a');
a.className = 'anchor-link';
a.href = '#' + h.id;
a.innerHTML = '#';
a.setAttribute('aria-label', 'Permalink to ' + h.innerText);
a.style.marginLeft = '0.4rem';
a.style.color = 'var(--primary)';
a.style.textDecoration = 'none';
a.style.opacity = '0.35';
a.style.fontSize = '0.85em';
a.style.transition = 'opacity 0.15s ease';
h.addEventListener('mouseenter', function () { a.style.opacity = '1'; });
h.addEventListener('mouseleave', function () { a.style.opacity = '0.35'; });
h.appendChild(a);
});
}
function initKeyboardNav() {
document.addEventListener('keydown', function (e) {
var tag = (e.target && e.target.tagName) || '';
if (tag === 'INPUT' || tag === 'TEXTAREA' || e.target.isContentEditable) return;
if (e.key === 'ArrowLeft') {
var prev = document.querySelector('nav.pager a.prev');
if (prev && prev.href) window.location.href = prev.href;
} else if (e.key === 'ArrowRight') {
var next = document.querySelector('nav.pager a.next');
if (next && next.href) window.location.href = next.href;
}
});
}
function initDocsLangSwitch() {
var path = window.location.pathname;
var isDocs = (typeof isDocsRoute === 'function' ? isDocsRoute(path) : path.indexOf('/docs') !== -1) || !!document.querySelector('.docs-reader, .chapter, #sidebar-toc');
if (!isDocs) {
return;
}

var isIndonesian = path.startsWith('/id/') || path === '/id';
var curLang = isIndonesian ? 'id' : 'en';

document.querySelectorAll('.lang-switch .lang-code, [data-lang-toggle] .lang-code').forEach(function (el) {
el.textContent = curLang.toUpperCase();
});

if (isIndonesian) {
document.querySelectorAll('.topbar nav .nav-links a, .sidebar-mobile-nav a').forEach(function (a) {
var href = a.getAttribute('href');
if (href && href.startsWith('/docs') && !href.startsWith('/id/')) {
a.setAttribute('href', '/id' + href);
}
});
}

var nav = document.querySelector('.topbar nav');
if (!nav || nav.querySelector('[data-lang-toggle], .lang-switch')) return;

var btn = document.createElement('button');
btn.type = 'button';
btn.className = 'lang-switch';
btn.setAttribute('data-lang-toggle', '');
btn.setAttribute('aria-label', 'Toggle language');
btn.setAttribute('title', 'Switch language (EN/ID)');
btn.innerHTML = '<span class="lang-icon" aria-hidden="true">🌐</span><span class="lang-code">' + curLang.toUpperCase() + '</span>';

var themeToggle = nav.querySelector('[data-theme-toggle]');
if (themeToggle && themeToggle.parentNode) {
themeToggle.parentNode.insertBefore(btn, themeToggle);
} else {
var actions = nav.querySelector('.topbar-actions');
if (actions) {
actions.appendChild(btn);
} else {
nav.appendChild(btn);
}
}
}
function initDocsReaderI18n() {
var title = document.querySelector('.sidebar-title');
if (title && !title.getAttribute('data-i18n')) {
title.setAttribute('data-i18n', 'docs_reader.toc_title');
}
var searchInput = document.getElementById('toc-search');
if (searchInput && !searchInput.getAttribute('data-i18n')) {
searchInput.setAttribute('data-i18n', 'docs_reader.search_placeholder');
}
if (window.krewireI18n && window.krewireI18n.setLocale) {
window.krewireI18n.setLocale(window.krewireI18n.getLocale());
}
window.addEventListener('krewire:langchange', function () {
document.querySelectorAll('.chapter pre .code-copy-btn .copy-txt:not(.copied)').forEach(function (span) {
span.textContent = getCopyText();
});
});
}
function initDocsSidebarDrawer() {
var sidebar = document.getElementById('sidebar');
var toggle = document.getElementById('sidebar-toggle');
var close = document.getElementById('sidebar-close');
var overlay = document.getElementById('sidebar-overlay');
function setOpen(open) {
document.documentElement.classList.toggle('sidebar-open', open);
if (toggle) toggle.setAttribute('aria-expanded', open ? 'true' : 'false');
if (sidebar) sidebar.setAttribute('aria-hidden', open ? 'false' : 'true');
}
if (toggle && !toggle._docsDrawerBound) {
toggle._docsDrawerBound = true;
toggle.addEventListener('click', function (e) {
e.preventDefault();
setOpen(!document.documentElement.classList.contains('sidebar-open'));
});
}
if (close && !close._docsDrawerBound) {
close._docsDrawerBound = true;
close.addEventListener('click', function (e) {
e.preventDefault();
setOpen(false);
});
}
if (overlay && !overlay._docsDrawerBound) {
overlay._docsDrawerBound = true;
overlay.addEventListener('click', function (e) {
e.preventDefault();
setOpen(false);
});
}
if (!window._docsEscBound) {
window._docsEscBound = true;
document.addEventListener('keydown', function (e) {
if (e.key === 'Escape') setOpen(false);
});
}
if (sidebar && !sidebar._docsLinksBound) {
sidebar._docsLinksBound = true;
sidebar.querySelectorAll('a.toc-link, a.toc-sub').forEach(function (link) {
link.addEventListener('click', function () {
if (window.innerWidth < 1024) {
setOpen(false);
}
});
});
}
document.querySelectorAll('[data-group] .toc-toggle').forEach(function (btn) {
if (btn._docsGroupBound) return;
btn._docsGroupBound = true;
btn.addEventListener('click', function (e) {
e.preventDefault();
e.stopPropagation();
var g = btn.closest('[data-group]');
if (!g) return;
var open = g.classList.toggle('open');
btn.setAttribute('aria-expanded', open ? 'true' : 'false');
var link = g.querySelector('.toc-link');
if (link) {
try {
localStorage.setItem('krewire-toc-' + link.getAttribute('href'), open ? '1' : '0');
} catch (_) {}
}
});
});
var toggleAll = document.getElementById('toc-toggle-all');
if (toggleAll && !toggleAll._docsGroupBound) {
toggleAll._docsGroupBound = true;
toggleAll.addEventListener('click', function (e) {
e.preventDefault();
var groups = document.querySelectorAll('[data-group]');
var anyClosed = Array.from(groups).some(function (g) { return !g.classList.contains('open'); });
groups.forEach(function (g) {
g.classList.toggle('open', anyClosed);
var b = g.querySelector('.toc-toggle');
if (b) b.setAttribute('aria-expanded', anyClosed ? 'true' : 'false');
var link = g.querySelector('.toc-link');
if (link) {
try {
localStorage.setItem('krewire-toc-' + link.getAttribute('href'), anyClosed ? '1' : '0');
} catch (_) {}
}
});
});
}
var searchInput = document.getElementById('toc-search');
var searchClear = document.getElementById('toc-search-clear');
var searchKbd = document.getElementById('toc-search-kbd');
var noResults = document.getElementById('toc-no-results');
if (searchInput) {
searchInput.addEventListener('input', function () {
var q = searchInput.value.trim().toLowerCase();
if (searchClear) searchClear.style.display = q ? 'inline-flex' : 'none';
if (searchKbd) searchKbd.style.display = q ? 'none' : '';
var matchCount = 0;
document.querySelectorAll('#sidebar-toc > *').forEach(function (el) {
if (!q) {
el.style.display = '';
if (el.hasAttribute('data-group')) {
el.querySelectorAll('.toc-sub').forEach(function (s) { s.style.display = ''; });
var l = el.querySelector('.toc-link');
var href = l ? l.getAttribute('href') : '';
var v = href ? localStorage.getItem('krewire-toc-' + href) : null;
var hasActive = el.querySelector('.active');
var open = hasActive || v === '1' || (v !== '0' && el.classList.contains('open'));
el.classList.toggle('open', !!open);
var b = el.querySelector('.toc-toggle');
if (b) b.setAttribute('aria-expanded', open ? 'true' : 'false');
}
return;
}
if (el.hasAttribute('data-group')) {
var parentMatch = (el.querySelector('.toc-link') ? el.querySelector('.toc-link').textContent || '' : '').toLowerCase().indexOf(q) !== -1;
var subMatches = 0;
el.querySelectorAll('.toc-sub').forEach(function (s) {
var m = (s.textContent || '').toLowerCase().indexOf(q) !== -1;
s.style.display = (m || parentMatch) ? '' : 'none';
if (m) subMatches++;
});
if (parentMatch || subMatches > 0) {
el.style.display = '';
el.classList.add('open');
var b = el.querySelector('.toc-toggle');
if (b) b.setAttribute('aria-expanded', 'true');
matchCount += (parentMatch ? 1 : 0) + subMatches;
} else {
el.style.display = 'none';
}
} else {
var match = (el.textContent || '').toLowerCase().indexOf(q) !== -1;
el.style.display = match ? '' : 'none';
if (match) matchCount++;
}
});
if (noResults) {
noResults.style.display = (q && matchCount === 0) ? 'block' : 'none';
}
});
if (searchClear) {
searchClear.addEventListener('click', function () {
searchInput.value = '';
searchInput.dispatchEvent(new Event('input'));
searchInput.focus();
});
}
document.addEventListener('keydown', function (e) {
var tag = (e.target && e.target.tagName) || '';
if (tag === 'INPUT' || tag === 'TEXTAREA' || (e.target && e.target.isContentEditable)) return;
if (e.key === '/' && !e.ctrlKey && !e.metaKey && !e.altKey) {
e.preventDefault();
searchInput.focus();
searchInput.select();
}
});
searchInput.addEventListener('keydown', function (e) {
if (e.key === 'Escape') {
if (searchInput.value) {
searchInput.value = '';
searchInput.dispatchEvent(new Event('input'));
} else {
searchInput.blur();
}
}
});
}
}
function initAll() {
initCodeBlocks();
initHeadingAnchors();
initKeyboardNav();
initDocsLangSwitch();
initDocsReaderI18n();
initDocsSidebarDrawer();
}
if (document.readyState === 'loading') {
document.addEventListener('DOMContentLoaded', initAll);
} else {
initAll();
}
})();
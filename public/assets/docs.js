// =============================================================================
// Krewire Documentation — Interactive Reader Enhancements (docs.js)
// Zero external dependencies — pure fast vanilla JS
// =============================================================================
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

  // ── 1. Code Block Copy Buttons & Language Badges ─────────────────────────
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

  // ── 2. Heading Anchor Links ──────────────────────────────────────────────
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

  // ── 3. Keyboard Pager Navigation (Left / Right Arrows) ───────────────────
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

  // ── 4. Documentation Topbar Language Switcher ───────────────────────────
  function initDocsLangSwitch() {
    var path = window.location.pathname;
    var isIndonesian = path.startsWith('/id/') || path === '/id';
    var curLang = isIndonesian ? 'id' : 'en';

    // Synchronize all existing buttons on page
    document.querySelectorAll('.lang-switch .lang-code, [data-lang-toggle] .lang-code').forEach(function (el) {
      el.textContent = curLang.toUpperCase();
    });

    // Localize topbar and mobile nav links if browsing Indonesian docs
    if (isIndonesian) {
      document.querySelectorAll('.topbar nav .nav-links a, .sidebar-mobile-nav a').forEach(function (a) {
        var href = a.getAttribute('href');
        if (href && href.startsWith('/docs') && !href.startsWith('/id/')) {
          a.setAttribute('href', '/id' + href);
        }
      });
    }

    // Delegate language click with full documentation route switching
    document.querySelectorAll('.lang-switch, [data-lang-toggle]').forEach(function (btn) {
      if (btn._docsLangBound) return;
      btn._docsLangBound = true;
      btn.addEventListener('click', function (e) {
        e.preventDefault();
        e.stopPropagation();
        var nextLang = curLang === 'en' ? 'id' : 'en';
        try {
          localStorage.setItem('krewire-lang', nextLang);
        } catch (_) {}

        var target;
        if (nextLang === 'id') {
          var clean = path.replace(/^\/en(\/|$)/, '/');
          target = '/id' + (clean.startsWith('/') ? clean : '/' + clean);
        } else {
          target = path.replace(/^\/id(\/|$)/, '/') || '/docs';
        }
        window.location.href = target;
      });
    });

    var nav = document.querySelector('.topbar nav');
    if (!nav || nav.querySelector('[data-lang-toggle], .lang-switch')) return;

    var btn = document.createElement('button');
    btn.type = 'button';
    btn.className = 'lang-switch';
    btn.setAttribute('data-lang-toggle', '');
    btn.setAttribute('aria-label', 'Toggle language');
    btn.setAttribute('title', 'Switch language (EN/ID)');
    btn.innerHTML = '<span class="lang-icon" aria-hidden="true">🌐</span><span class="lang-code">' + curLang.toUpperCase() + '</span>';
    btn._docsLangBound = true;
    btn.addEventListener('click', function (e) {
      e.preventDefault();
      var nextLang = curLang === 'en' ? 'id' : 'en';
      try {
        localStorage.setItem('krewire-lang', nextLang);
      } catch (_) {}
      var target;
      if (nextLang === 'id') {
        var clean = path.replace(/^\/en(\/|$)/, '/');
        target = '/id' + (clean.startsWith('/') ? clean : '/' + clean);
      } else {
        target = path.replace(/^\/id(\/|$)/, '/') || '/docs';
      }
      window.location.href = target;
    });

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

  // ── 5. Documentation Reader i18n Bindings ──────────────────────────────
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

  // ── 6. Documentation Sidebar Drawer Controller ──────────────────────────
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

    // Close drawer when clicking a navigation link inside sidebar on mobile
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

    // Collapsible sub-chapter groups
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

    // Expand / collapse all groups button
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

    // Quick chapter search/filter
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

  // ── 7. Auto-Run on DOMContentLoaded ──────────────────────────────────────
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

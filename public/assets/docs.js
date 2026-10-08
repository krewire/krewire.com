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
    var nav = document.querySelector('.topbar nav');
    if (!nav || nav.querySelector('[data-lang-toggle], .lang-switch')) return;

    var curLang = 'en';
    if (window.krewireI18n && window.krewireI18n.getLocale) {
      curLang = window.krewireI18n.getLocale();
    } else {
      try {
        curLang = localStorage.getItem('krewire-lang') || 'en';
      } catch (e) {}
    }

    var btn = document.createElement('button');
    btn.type = 'button';
    btn.className = 'lang-switch';
    btn.setAttribute('data-lang-toggle', '');
    btn.setAttribute('aria-label', 'Toggle language');
    btn.setAttribute('title', 'Switch language (EN/ID)');
    btn.innerHTML = '<span class="lang-icon" aria-hidden="true">🌐</span><span class="lang-code">' + curLang.toUpperCase() + '</span>';

    btn.addEventListener('click', function (e) {
      e.preventDefault();
      if (window.krewireI18n && window.krewireI18n.toggle) {
        window.krewireI18n.toggle();
      }
    });

    var themeToggle = nav.querySelector('[data-theme-toggle]');
    if (themeToggle) {
      nav.insertBefore(btn, themeToggle);
    } else {
      nav.appendChild(btn);
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

  // ── 6. Auto-Run on DOMContentLoaded ──────────────────────────────────────
  function initAll() {
    initCodeBlocks();
    initHeadingAnchors();
    initKeyboardNav();
    initDocsLangSwitch();
    initDocsReaderI18n();
  }

  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', initAll);
  } else {
    initAll();
  }
})();

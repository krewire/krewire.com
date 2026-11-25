---
title: "Utilitas Pengujian & Asersi"
description: "Harness asersi pengujian, rantai HTTP fasih, snapshot golden, asersi komponen UI, dan otomasi browser di testing."
date: "2026-10-25"
---

# Utilitas Pengujian & Asersi (`testing`)

Jalur impor: `github.com/krewire/krewire/testing`

`testing` menghadirkan perangkat pengujian murni Go (pure-Go) yang idiomatik untuk ekosistem Krewire. Paket ini mencakup pembangun permintaan HTTP yang fasih (fluent), matcher asersi respons yang dapat dirangkai berantai (chainable), manajemen fixture sementara, perbandingan berkas snapshot golden, inspeksi komponen UI, serta otomasi peramban tanpa kepala (headless browser automation).

---

## 1. Rantai Pengujian HTTP Fasih (Fluent)

Klien HTTP fasih menyederhanakan pengujian handler web dan REST API tanpa boilerplate penanganan galat berulang:

```go
package web_test

import (
	"net/http"
	"testing"

	ftest "github.com/krewire/krewire/testing"
)

func TestProfileEndpoint(t *testing.T) {
	handler := SetupServer()

	// 1. Permintaan tanpa autentikasi ditolak
	req := ftest.GET(t, "/api/profile").Request()
	ftest.Do(t, handler, req).
		Status(http.StatusUnauthorized).
		Header("Content-Type", "application/json").
		Contains("missing authentication token")

	// 2. Permintaan JSON terautentikasi berhasil
	payload := map[string]string{"name": "Alice"}
	req = ftest.JSONRequest(t, "POST", "/api/profile", payload).
		Header("Authorization", "Bearer valid-token").
		Request()

	var res struct {
		Status string `json:"status"`
	}

	ftest.Do(t, handler, req).
		Status(http.StatusCreated).
		Header("Content-Type", "application/json").
		JSON(&res)

	ftest.Equal(t, res.Status, "ok")
}
```

---

## 2. Primitif Asersi

- **`test.Equal(t, got, want)`**: Pemeriksaan kesetaraan mendalam (deep equality).
- **`test.NoError(t, err)`**: Memastikan `err == nil`.
- **`test.HasError(t, err)`**: Memastikan `err != nil`.
- **`test.Contains(t, s, substr)`**: Pemeriksaan substring pada string.
- **`test.True(t, cond, msg)`**: Asersi kondisi boolean.

---

## 3. Pengujian Snapshot Golden

Berkas golden memastikan luaran HTML dan JSON yang dirender tidak mengalami regresi yang tidak disengaja:

```go
package compiler_test

import (
	"testing"

	ftest "github.com/krewire/krewire/testing"
)

func TestTemplateRendering(t *testing.T) {
	output := RenderTemplate("card", data)

	// Membandingkan luaran dengan testdata/card.golden
	ftest.AssertGolden(t, "testdata/card.golden", []byte(output))
}
```

---

## 4. Otomasi Peramban Tanpa Kepala (`testing/browser`)

Paket `browser` menyediakan harness pengujian end-to-end terotomasi tanpa memerlukan runtime Node.js:

```go
package e2e_test

import (
	"net/http/httptest"
	"testing"

	"github.com/krewire/krewire/testing/browser"
)

func TestWebUI_E2E(t *testing.T) {
	srv := httptest.NewServer(SetupRouter())
	defer srv.Close()

	b := browser.New(t, srv.URL)
	b.Navigate("/")
	b.AssertStatus(200)
	b.AssertTextVisible("Welcome to Krewire")
}
```

---

## 5. Keselarasan dengan Tiga Pilar (3S)

- **Aman (S1 - Secure)**: Eksekusi hermetik dalam-proses mencegah pembukaan port jaringan publik selama rangkaian uji terotomasi.
- **Berkelanjutan (S2 - Sustainable)**: Murni Go tanpa memerlukan Node.js, Jest, Cypress, atau daemon pengujian yang berat.
- **Skalabel (S3 - Scalable)**: Kontrak pengujian terpadu yang dibagikan di seluruh API HTTP, pulau WebAssembly, dan perkakas CLI.

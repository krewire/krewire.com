---
title: "Web Engine & Perutean HTTP"
description: "Perutean HTTP radix tree berkinerja tinggi, pipeline middleware, problem details RFC 7807, ekstraksi parameter, dan helper render di pkg/web."
date: "2026-10-25"
---

# Web Engine & Perutean HTTP (`pkg/web`)

Jalur impor: `github.com/krewire/krewire/pkg/web`

Paket `web` membungkus pustaka standar Go `net/http` dengan perutean ekspresif, ekstraksi parameter, rantai middleware, dan serialisasi JSON tanpa dependensi runtime eksternal.

---

## Fitur Utama

- **Perutean Zero Allocations:** Router radix tree berkecepatan tinggi yang mendukung grup rute, parameter path (`/users/:id`), dan wildcard (`/static/*`).
- **Kompatibilitas Pustaka Standar:** Middleware dan handler sepenuhnya mengimplementasikan standar Go `http.Handler` dan `http.HandlerFunc`.
- **RFC 7807 Problem Details:** Standar respons error JSON untuk API publik dan privat.
- **Helper Render:** Fungsi bawaan untuk `web.JSON()`, `web.HTML()`, dan `web.Redirect()`.

---

## Contoh Penggunaan

```go
package main

import (
	"net/http"
	"github.com/krewire/krewire/pkg/web"
)

func main() {
	r := web.NewRouter()

	// Middleware Global
	r.Use(web.Recovery())
	r.Use(web.Logger())

	// Grup Rute
	api := r.Group("/api/v1")
	{
		api.Get("/items", func(w http.ResponseWriter, r *http.Request) {
			items := []string{"buku", "keyboard", "layar"}
			web.JSON(w, http.StatusOK, items)
		})

		api.Get("/items/:id", func(w http.ResponseWriter, r *http.Request) {
			id := web.Param(r, "id")
			web.JSON(w, http.StatusOK, map[string]string{"id": id})
		})
	}

	http.ListenAndServe(":8080", r)
}
```\n
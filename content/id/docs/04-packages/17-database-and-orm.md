---
title: "Database & Multi-Driver ORM"
description: "Panduan referensi untuk ORM multi-driver Krewire: SQLite, MySQL, PostgreSQL, MongoDB, dan Elasticsearch dengan fluent query builder serta migrasi skema."
date: "2026-10-24"
---

# Database & Multi-Driver ORM (`pkg/orm`)

`pkg/orm` menyediakan Object-Relational Mapping (ORM) ekspresif dan query builder agnostik database yang dirancang untuk ekosistem Go yang aman secara tipe (*type-safe*).

## Driver yang Didukung

| Driver | Target Database | Protokol / Karakteristik |
|---|---|---|
| `sqlite` | SQLite 3 | Tersemat (*embedded*), bebas CGO, sangat ideal untuk node edge dan pengembangan lokal |
| `mysql` | MySQL 8.0+ / MariaDB | Protokol relasional standar dengan pooling koneksi otomatis |
| `postgres` | PostgreSQL 14+ | Fitur relasional tingkat lanjut, tipe data JSONB, dan enkripsi SSL |
| `mongodb` | MongoDB 6.0+ | Penyimpanan NoSQL berorientasi dokumen dengan pemetaan BSON |
| `elasticsearch` | Elasticsearch 8+ / OpenSearch | Pencarian teks lengkap (*full-text search*) dan pengindeksan dokumen analitis |

## Contoh Penggunaan Cepat

```go
package main

import (
    "context"
    "fmt"
    
    "github.com/krewire/krewire/pkg/orm"
)

type User struct {
    ID    string `orm:"id,primary"`
    Name  string `orm:"name"`
    Email string `orm:"email,unique"`
}

func main() {
    ctx := context.Background()
    
    // Terhubung ke SQLite (atau postgres/mysql via DSN)
    db, err := orm.Open(orm.Config{
        Driver: "sqlite",
        DSN:    ":memory:",
    })
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Fluent Query Builder
    var users []User
    err = db.Model(&User{}).
        Where("email", "LIKE", "%@krewire.com").
        OrderBy("name", "ASC").
        Limit(10).
        Find(ctx, &users)
}
```

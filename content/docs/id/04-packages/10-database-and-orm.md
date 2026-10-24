---
title: "Database & Multi-Driver ORM"
description: "Reference guide for Krewire's multi-driver ORM: SQLite, MySQL, PostgreSQL, MongoDB, and Elasticsearch with fluent query builder and migration helpers."
date: "2026-10-24"
---

# Database & Multi-Driver ORM (`pkg/orm`)

`pkg/orm` provides an expressive, database-agnostic Object-Relational Mapping (ORM) and query builder designed for Go's type-safe ecosystem.

## Supported Drivers

| Driver | Target Database | Protocol / Architecture |
|---|---|---|
| `sqlite` | SQLite 3 | Embedded, zero-CGO, ideal for edge nodes and local development |
| `mysql` | MySQL 8.0+ / MariaDB | Standard relational protocol with connection pooling |
| `postgres` | PostgreSQL 14+ | Advanced relational features, JSONB, and SSL modes |
| `mongodb` | MongoDB 6.0+ | Document-oriented NoSQL storage with BSON mapping |
| `elasticsearch` | Elasticsearch 8+ / OpenSearch | Full-text search and analytical document indexing |

## Quick Usage

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
    
    // Connect to SQLite (or postgres/mysql via DSN)
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

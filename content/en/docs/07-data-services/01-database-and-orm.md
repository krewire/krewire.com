---
title: "Database & Multi-Driver ORM"
description: "Reference guide for Krewire's multi-driver ORM: SQLite, MySQL, PostgreSQL, MongoDB, and Elasticsearch with fluent query builder, Laravel-like schema migrations, seeders, factories, and generic repositories."
date: "2026-11-11"
---

# Database & Multi-Driver ORM (`pkg/orm`)

`pkg/orm` provides an expressive, database-agnostic Object-Relational Mapping (ORM) and query builder designed for Go's type-safe ecosystem. It eliminates external CGO dependencies and abstracts dialect differences (parameter placeholders, identifier quoting, DDL statements, auto-increment conventions) behind a single fluent API.

## Supported Drivers

| Driver | Target Database | Protocol / Architecture |
|---|---|---|
| `sqlite` | SQLite 3 | Embedded, zero-CGO, WAL mode, foreign keys enforced; bundled in `pkg/orm/sqlite` |
| `mysql` | MySQL 8.0+ / MariaDB | Standard relational protocol with connection pooling |
| `postgres` | PostgreSQL 14+ | Advanced relational features, JSONB, transaction advisory locks |
| `mongodb` | MongoDB 6.0+ | Document-oriented NoSQL storage with BSON mapping |
| `elasticsearch` | Elasticsearch 8+ / OpenSearch | Full-text search and analytical document indexing |

## Models, Timestamps & Soft Deletes

Krewire models provide built-in timestamps, soft deletes, and lifecycle hooks via embedded structs:

```go
package models

import (
    "context"

    "github.com/krewire/krewire/pkg/orm"
)

type User struct {
    orm.BaseModel // Embeds ID (int64), Timestamps (CreatedAt, UpdatedAt), and SoftDeletes (DeletedAt)

    Name     string `json:"name" orm:"column:name"`
    Email    string `json:"email" orm:"column:email;unique"`
    IsActive bool   `json:"is_active" orm:"column:is_active;default:1"`
}

func (User) TableName() string {
    return "users"
}

// Lifecycle Hooks: BeforeCreate, AfterCreate, BeforeUpdate, AfterUpdate, BeforeDelete, AfterDelete, BeforeSave, AfterSave
func (u *User) BeforeCreate(ctx context.Context) error {
    if u.Name == "" {
        u.Name = "Anonymous"
    }
    return nil
}
```

### Soft Deletion & Query Scopes

Queries automatically filter out soft-deleted records (`WHERE deleted_at IS NULL`). Use `.WithTrashed()` to include soft-deleted records or `.OnlyTrashed()` to view only soft-deleted records:

```go
// Normal query: excludes soft-deleted records
var users []models.User
err := db.Model(&models.User{}).FindAll(&users)

// Include soft-deleted records
err = db.Model(&models.User{}).WithTrashed().FindAll(&users)

// Soft delete
_, err = db.Model(&models.User{}).Delete(&user)

// Restore a soft-deleted record
_, err = db.Model(&models.User{}).Restore(&user)

// Permanently delete (hard delete)
_, err = db.Model(&models.User{}).ForceDelete(&user)
```

## Fluent Query Builder

The fluent query builder provides Eloquent-like ergonomics with pure Go type safety:

```go
// Pluck specific column values
var emails []string
err := db.Model(&models.User{}).Where("is_active", "=", 1).Pluck("email", &emails)

// Chunk large datasets
var chunk []models.User
err = db.Model(&models.User{}).OrderBy("id", "ASC").Chunk(100, &chunk, func() error {
    for _, u := range chunk {
        processUser(u)
    }
    return nil
})

// Joins, Group By, Having
err = db.Table("orders").
    Select("users.name", "SUM(orders.amount) as total_spent").
    InnerJoin("users", "orders.user_id = users.id").
    GroupBy("users.name").
    Having("total_spent > ?", 1000).
    FindAll(&report)
```

## Programmatic Migrations (`schema.Blueprint`)

Krewire supports Laravel-like schema migrations in Go using `schema.Blueprint`:

```go
package migrations

import (
    "context"

    "github.com/krewire/krewire/pkg/orm/migrate"
    "github.com/krewire/krewire/pkg/orm/schema"
)

func init() {
    migrate.RegisterGo(migrate.GoMigration{
        Version: "20261111100000",
        Name:    "create_products_table",
        Up: func(ctx context.Context, s *schema.Builder) error {
            return s.Create("products", func(t *schema.Blueprint) {
                t.BigIncrements("id")
                t.String("name", 255).NotNull()
                t.String("sku", 100).Unique()
                t.Decimal("price", 10, 2).Default(0)
                t.Boolean("in_stock").Default(true)
                t.Morphs("commentable") // commentable_id + commentable_type
                t.Timestamps()
                t.SoftDeletes()
            })
        },
        Down: func(ctx context.Context, s *schema.Builder) error {
            return s.DropTableIfExists("products")
        },
    })
}
```

Run migrations via the CLI:
```bash
kiw migrate up       # Apply all pending migrations
kiw migrate down     # Roll back the last migration batch
kiw migrate redo     # Re-apply the last migration atomically
kiw migrate status   # Inspect migration states
```

## Faker Engine (`pkg/orm/faker`)

A pure-Go, deterministic, thread-safe fake data generator for testing and seeding:

```go
import "github.com/krewire/krewire/pkg/orm/faker"

f := faker.New() // or faker.New(42) for deterministic reproducible seeds

name     := f.Name()
email    := f.SafeEmail()
phone    := f.Phone()
address  := f.Address()
company  := f.Company()
sentence := f.Sentence(8)
price    := f.Price(5.0, 250.0)
uuid     := f.UUID()
```

## Seeder & Model Factory Engine (`pkg/orm/seeder`)

Create reproducible test datasets and initial system state using generic model factories and seeders:

### Model Factories

```go
package factories

import (
    "github.com/krewire/krewire/pkg/orm"
    "github.com/krewire/krewire/pkg/orm/faker"
    "github.com/krewire/krewire/pkg/orm/seeder"
    "my-app/models"
)

func UserFactory(db *orm.DB) *seeder.Factory[models.User] {
    return seeder.NewFactory(db, func(f *faker.Faker) *models.User {
        return &models.User{
            Name:     f.Name(),
            Email:    f.SafeEmail(),
            IsActive: true,
        }
    })
}

// Generate in-memory instances:
user := UserFactory(db).Make(ctx)
users := UserFactory(db).Count(10).MakeMany(ctx)

// Persist directly into database:
savedUser, err := UserFactory(db).Create(ctx)
savedUsers, err := UserFactory(db).Count(50).CreateMany(ctx)

// State overrides:
admins, err := UserFactory(db).State(func(u *models.User) {
    u.IsActive = true
    u.Name = "Admin " + u.Name
}).Count(5).CreateMany(ctx)
```

### Seeders

```go
package seeds

import (
    "context"

    "github.com/krewire/krewire/pkg/orm"
    "github.com/krewire/krewire/pkg/orm/seeder"
    "my-app/factories"
)

func init() {
    seeder.Register("DatabaseSeeder", DatabaseSeeder{})
}

type DatabaseSeeder struct{}

func (DatabaseSeeder) Run(ctx context.Context, db *orm.DB) error {
    _, err := factories.UserFactory(db).Count(25).CreateMany(ctx)
    return err
}
```

Run seeders via the CLI:
```bash
kiw seed                      # runs DatabaseSeeder
kiw seed --class UserSeeder   # runs a specific seeder
kiw db:seed                   # Laravel-compatible alias
```

## Generic Repository Pattern (`pkg/orm/repository`)

Krewire provides a type-safe generic Repository pattern (`Repository[T any]`) that decouples persistence concerns from domain services and provides transaction safety:

```go
package repositories

import (
    "context"

    "github.com/krewire/krewire/pkg/orm"
    "github.com/krewire/krewire/pkg/orm/repository"
    "my-app/models"
)

type UserRepository interface {
    repository.Repository[models.User]
}

type userRepositoryImpl struct {
    *repository.BaseRepository[models.User]
}

func NewUserRepository(db *orm.DB) UserRepository {
    return &userRepositoryImpl{
        BaseRepository: repository.New[models.User](db),
    }
}
```

### Repository Operations & Transactions

```go
repo := repositories.NewUserRepository(db)
ctx := context.Background()

// CRUD operations
user, err := repo.FindByID(ctx, 1)
allUsers, err := repo.FindAll(ctx)
err = repo.Create(ctx, &models.User{Name: "Bob", Email: "bob@example.com"})
err = repo.Save(ctx, user) // Inserts if ID==0, updates if ID>0
err = repo.Delete(ctx, 1)

// Pagination with navigation metadata
page, err := repo.Paginate(ctx, 1, 10, func(q *orm.Query) *orm.Query {
    return q.Where("is_active", "=", 1).OrderBy("id", "DESC")
})
// page.Items, page.Total, page.TotalPages, page.HasNext, page.HasPrev

// Atomic database transactions
err = repo.Transaction(ctx, func(txRepo repository.Repository[models.User]) error {
    if err := txRepo.Create(ctx, u1); err != nil {
        return err // automatically rolled back on error
    }
    return txRepo.Create(ctx, u2) // committed automatically on return nil
})
```

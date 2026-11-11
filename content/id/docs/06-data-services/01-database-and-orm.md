---
title: "Database & Multi-Driver ORM"
description: "Panduan referensi untuk ORM multi-driver Krewire: SQLite, MySQL, PostgreSQL, MongoDB, dan Elasticsearch dengan fluent query builder, migrasi skema deklaratif, seeder, factory, dan repository generik."
date: "2026-11-11"
---

# Database & Multi-Driver ORM (`pkg/orm`)

`pkg/orm` menyediakan Object-Relational Mapping (ORM) ekspresif dan query builder agnostik database yang dirancang untuk ekosistem Go yang aman secara tipe (*type-safe*). Pustaka ini meniadakan ketergantungan CGO eksternal dan mengabstraksikan perbedaan dialek SQL (placeholder parameter, pengutipan identifier, DDL, dan konvensi auto-increment) di balik satu API yang bersih.

## Driver yang Didukung

| Driver | Target Database | Protokol / Karakteristik |
|---|---|---|
| `sqlite` | SQLite 3 | Tersemat (*embedded*), bebas CGO, mode WAL, foreign key aktif; terpaket di `pkg/orm/sqlite` |
| `mysql` | MySQL 8.0+ / MariaDB | Protokol relasional standar dengan connection pooling otomatis |
| `postgres` | PostgreSQL 14+ | Fitur relasional tingkat lanjut, tipe JSONB, advisory lock berbasis transaksi |
| `mongodb` | MongoDB 6.0+ | Penyimpanan dokumen NoSQL dengan pemetaan BSON otomatis |
| `elasticsearch` | Elasticsearch 8+ / OpenSearch | Pencarian teks lengkap (*full-text search*) dan pengindeksan dokumen analitis |

## Model, Timestamps & Soft Deletes

Model Krewire menyediakan dukungan otomatis untuk timestamp, penghapusan lunak (*soft delete*), dan lifecycle hook melalui embedded struct:

```go
package models

import (
    "context"

    "github.com/krewire/krewire/pkg/orm"
)

type User struct {
    orm.BaseModel // Menyematkan ID (int64), Timestamps (CreatedAt, UpdatedAt), dan SoftDeletes (DeletedAt)

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

### Soft Deletion & Scope Kueri

Kueri secara otomatis menyaring data yang telah dihapus secara lunak (`WHERE deleted_at IS NULL`). Gunakan `.WithTrashed()` untuk menyertakan data terhapus atau `.OnlyTrashed()` untuk hanya melihat data yang dihapus:

```go
// Kueri normal: mengabaikan data yang dihapus lunak
var users []models.User
err := db.Model(&models.User{}).FindAll(&users)

// Menyertakan data yang dihapus lunak
err = db.Model(&models.User{}).WithTrashed().FindAll(&users)

// Soft delete
_, err = db.Model(&models.User{}).Delete(&user)

// Memulihkan data yang dihapus lunak
_, err = db.Model(&models.User{}).Restore(&user)

// Menghapus secara permanen (hard delete)
_, err = db.Model(&models.User{}).ForceDelete(&user)
```

## Fluent Query Builder

Query builder menyediakan ergonomi kueri yang ekspresif menyerupai Laravel Eloquent dengan keamanan tipe Go murni:

```go
// Mengambil kolom tertentu (Pluck)
var emails []string
err := db.Model(&models.User{}).Where("is_active", "=", 1).Pluck("email", &emails)

// Chunking data besar
var chunk []models.User
err = db.Model(&models.User{}).OrderBy("id", "ASC").Chunk(100, &chunk, func() error {
    for _, u := range chunk {
        prosesPengguna(u)
    }
    return nil
})

// Join, Group By, Having
err = db.Table("orders").
    Select("users.name", "SUM(orders.amount) as total_spent").
    InnerJoin("users", "orders.user_id = users.id").
    GroupBy("users.name").
    Having("total_spent > ?", 1000).
    FindAll(&laporan)
```

## Migrasi Terprogram (`schema.Blueprint`)

Krewire mendukung pembuatan skema secara deklaratif menyerupai Laravel Blueprint di dalam kode Go:

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

Jalankan migrasi melalui CLI:
```bash
kiw migrate up       # Jalankan seluruh migrasi tertunda
kiw migrate down     # Kembalikan (rollback) kelompok migrasi terakhir
kiw migrate redo     # Jalankan ulang migrasi terakhir secara atomik
kiw migrate status   # Periksa status seluruh berkas migrasi
```

## Faker Engine (`pkg/orm/faker`)

Generator data tiruan (*fake data*) deterministik dan aman terhadap goroutine (*thread-safe*) berbasis Go murni:

```go
import "github.com/krewire/krewire/pkg/orm/faker"

f := faker.New() // atau faker.New(42) untuk data deterministik yang dapat direproduksi

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

Buat dataset pengujian yang konsisten dan inisialisasi status sistem menggunakan Model Factory generik dan Seeder:

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

// Buat instans di memori:
user := UserFactory(db).Make(ctx)
users := UserFactory(db).Count(10).MakeMany(ctx)

// Simpan langsung ke database:
savedUser, err := UserFactory(db).Create(ctx)
savedUsers, err := UserFactory(db).Count(50).CreateMany(ctx)

// Modifikasi status (State):
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

Jalankan seeder melalui CLI:
```bash
kiw seed                      # Menjalankan DatabaseSeeder
kiw seed --class UserSeeder   # Menjalankan seeder tertentu
kiw db:seed                   # Alias kompatibel dengan Laravel
```

## Generic Repository Pattern (`pkg/orm/repository`)

Pola generic Repository (`Repository[T any]`) memisahkan urusan persistensi data dari logika bisnis domain dengan keamanan transaksi database penuh:

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

### Operasi Repository & Transaksi Atomik

```go
repo := repositories.NewUserRepository(db)
ctx := context.Background()

// Operasi CRUD
user, err := repo.FindByID(ctx, 1)
allUsers, err := repo.FindAll(ctx)
err = repo.Create(ctx, &models.User{Name: "Budi", Email: "budi@example.com"})
err = repo.Save(ctx, user) // Insert jika ID==0, Update jika ID>0
err = repo.Delete(ctx, 1)

// Paginasi dengan metadata navigasi
page, err := repo.Paginate(ctx, 1, 10, func(q *orm.Query) *orm.Query {
    return q.Where("is_active", "=", 1).OrderBy("id", "DESC")
})
// page.Items, page.Total, page.TotalPages, page.HasNext, page.HasPrev

// Transaksi atomik database
err = repo.Transaction(ctx, func(txRepo repository.Repository[models.User]) error {
    if err := txRepo.Create(ctx, u1); err != nil {
        return err // di-rollback otomatis saat error
    }
    return txRepo.Create(ctx, u2) // di-commit otomatis saat return nil
})
```

// Package db opens the SQLite database and applies the embedded migrations.
package db

import (
	"context"
	"fmt"
	"io/fs"
	"sort"
	"strings"
	"time"

	"github.com/krewire/krewire/pkg/orm"
	_ "modernc.org/sqlite" // pure-Go SQLite driver, registers "sqlite"

	"todo-app/migrations"
)

// Open opens (or creates) the SQLite file at path and migrates it.
func Open(ctx context.Context, path string) (*orm.DB, error) {
	db, err := orm.NewSQLite(orm.SQLiteConfig{
		Path:       path,
		DriverName: "sqlite",
		// One writer keeps SQLite simple and avoids "database is locked".
		MaxOpenConns: 1,
	})
	if err != nil {
		return nil, err
	}
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("db: ping: %w", err)
	}
	if err := migrate(ctx, db); err != nil {
		return nil, err
	}
	return db, nil
}

// migrate runs the "-- +up" section of every migration that is not yet
// recorded in schema_migrations, in file-name (timestamp) order.
func migrate(ctx context.Context, db *orm.DB) error {
	if _, err := db.Exec(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (
  version    TEXT PRIMARY KEY,
  applied_at DATETIME NOT NULL
)`); err != nil {
		return fmt.Errorf("db: schema_migrations: %w", err)
	}
	names, err := fs.Glob(migrations.Files, "*.sql")
	if err != nil {
		return err
	}
	sort.Strings(names)
	for _, name := range names {
		var n int
		if err := db.QueryRow(ctx, `SELECT COUNT(*) FROM schema_migrations WHERE version = ?`, name).Scan(&n); err != nil {
			return err
		}
		if n > 0 {
			continue
		}
		src, err := migrations.Files.ReadFile(name)
		if err != nil {
			return err
		}
		// orm.Tx exposes only Commit/Rollback, so the migration transaction
		// uses the underlying *sql.DB.
		tx, err := db.SQLDB().BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, upSection(string(src))); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("db: migrate %s: %w", name, err)
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO schema_migrations (version, applied_at) VALUES (?, ?)`, name, time.Now().UTC()); err != nil {
			_ = tx.Rollback()
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}

// upSection returns the SQL between "-- +up" and "-- +down".
func upSection(src string) string {
	_, after, found := strings.Cut(src, "-- +up")
	if !found {
		return src
	}
	up, _, _ := strings.Cut(after, "-- +down")
	return up
}

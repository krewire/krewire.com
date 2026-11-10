// Package db opens the SQLite database and applies the embedded migrations.
package db

import (
	"context"

	"github.com/krewire/krewire/pkg/orm"
	"github.com/krewire/krewire/pkg/orm/migrate"
	"github.com/krewire/krewire/pkg/orm/sqlite" // bundled pure-Go SQLite driver

	"todo-app/migrations"
)

// Open opens (or creates) the SQLite file at path and applies every pending
// migration, so a fresh deployment starts with the right schema.
func Open(ctx context.Context, path string) (*orm.DB, error) {
	db, err := sqlite.Open(path)
	if err != nil {
		return nil, err
	}
	if err := migrate.Up(ctx, db, migrations.Files); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}

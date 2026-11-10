// Package migrations embeds the SQL files created by `kiw make migration`.
package migrations

import "embed"

//go:embed *.sql
var Files embed.FS

// Package assets embeds the application's static assets served as-is
package assets

import "embed"

//go:embed public
var Files embed.FS

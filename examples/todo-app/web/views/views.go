// Package views turns the .kiw files in this directory into components the
// web app can render, the same way `kiw build` loads components/ for a site.
package views

import (
	"embed"
	"io/fs"
	"strings"

	"github.com/krewire/krewire/pkg/kiw"
	"github.com/krewire/krewire/pkg/ui"
	"github.com/krewire/krewire/pkg/web"
	"github.com/krewire/krewire/pkg/web/ssg"
)

//go:embed *.kiw
var files embed.FS

// Register parses the Krewire UI catalog (Card, Input, Button, Badge, Alert,
// ...) and this package's own .kiw views, and registers them as components
// on the app. A view with the same name as a UI component wins.
func Register(a *web.App) error {
	var css strings.Builder
	if err := register(a, ui.ComponentsFS, &css); err != nil {
		return err
	}
	if err := register(a, files, &css); err != nil {
		return err
	}
	// Pages link /assets/style.css automatically, but a live web.App does not
	// serve the component <style> blocks yet, so serve them ourselves.
	a.Asset("assets/style.css", css.String())
	return nil
}

func register(a *web.App, fsys fs.FS, css *strings.Builder) error {
	names, err := fs.Glob(fsys, "*.kiw")
	if err != nil {
		return err
	}
	for _, name := range names {
		src, err := fs.ReadFile(fsys, name)
		if err != nil {
			return err
		}
		mod, err := kiw.ParseKiw(string(src))
		if err != nil {
			return err
		}
		style := strings.Join(mod.Styles, "\n")
		css.WriteString(style + "\n")
		a.Component(ssg.Component{
			Name:  strings.TrimSuffix(name, ".kiw"),
			Body:  mod.Body,
			Style: style,
		})
	}
	return nil
}

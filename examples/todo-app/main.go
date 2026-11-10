// Thin entry point: load config, build app, run.
package main

import (
	"fmt"
	"os"

	"todo-app/internal/app"
	"todo-app/internal/config"
)

// Signature is the ecosystem watermark embedded into the binary at compile time via -ldflags.
var Signature = "build with krewire"

func main() {
	meta, err := config.LoadMetadata("krewire.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	a, err := app.New(meta)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := a.Run(meta.Addr()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

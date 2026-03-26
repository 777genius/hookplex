package main

import (
	"os"

	hookplex "github.com/hookplex/hookplex/sdk"
	"github.com/hookplex/hookplex/sdk/claude"
)

func main() {
	app := hookplex.New(hookplex.Config{Name: "hookplex-smoke"})
	app.Claude().OnStop(func(e *claude.StopEvent) *claude.Response {
		_ = e // smoke: allow stop
		return claude.Allow()
	})
	os.Exit(app.Run())
}

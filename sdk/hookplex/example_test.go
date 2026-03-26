package hookplex_test

import (
	hookplex "github.com/hookplex/hookplex/sdk"
	"github.com/hookplex/hookplex/sdk/claude"
	"github.com/hookplex/hookplex/sdk/codex"
)

func ExampleApp_Claude() {
	app := hookplex.New(hookplex.Config{Name: "demo"})
	app.Claude().OnStop(func(*claude.StopEvent) *claude.Response {
		return claude.Allow()
	})
	_ = app
}

func ExampleApp_Codex() {
	app := hookplex.New(hookplex.Config{Name: "demo"})
	app.Codex().OnNotify(func(*codex.NotifyEvent) *codex.Response {
		return codex.Continue()
	})
	_ = app
}

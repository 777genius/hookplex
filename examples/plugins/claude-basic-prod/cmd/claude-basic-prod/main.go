package main

import (
	"os"

	pluginkitai "github.com/plugin-kit-ai/plugin-kit-ai/sdk"
	"github.com/plugin-kit-ai/plugin-kit-ai/sdk/claude"
)

func main() {
	app := pluginkitai.New(pluginkitai.Config{Name: "claude-basic-prod"})
	app.Claude().OnStop(func(e *claude.StopEvent) *claude.Response {
		_ = e
		return claude.Allow()
	})
	app.Claude().OnPreToolUse(func(e *claude.PreToolUseEvent) *claude.PreToolResponse {
		_ = e
		return claude.PreToolAllow()
	})
	app.Claude().OnUserPromptSubmit(func(e *claude.UserPromptEvent) *claude.UserPromptResponse {
		_ = e
		return claude.UserPromptAllow()
	})
	os.Exit(app.Run())
}

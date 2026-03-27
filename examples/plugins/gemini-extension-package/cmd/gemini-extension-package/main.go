package main

import (
	"os"

	pluginkitai "github.com/plugin-kit-ai/plugin-kit-ai/sdk"
)

func main() {
	app := pluginkitai.New(pluginkitai.Config{Name: "gemini-extension-package"})
	os.Exit(app.Run())
}

package main

import (
	"os"

	"github.com/777genius/plugin-kit-ai/cli/internal/exitx"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(exitx.Code(err))
	}
}

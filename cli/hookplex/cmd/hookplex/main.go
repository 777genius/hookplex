package main

import (
	"os"

	"github.com/hookplex/hookplex/cli/internal/exitx"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(exitx.Code(err))
	}
}

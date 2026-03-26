package main

import (
	"fmt"

	"github.com/hookplex/hookplex/cli/internal/validate"
	"github.com/spf13/cobra"
)

var validatePlatform string

var validateCmd = &cobra.Command{
	Use:   "validate [path]",
	Short: "Validate a hookplex project scaffold",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validate.Run(args[0], validatePlatform); err != nil {
			return err
		}
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Validated %s\n", args[0])
		return nil
	},
}

func init() {
	validateCmd.Flags().StringVar(&validatePlatform, "platform", "", `platform override ("codex" or "claude")`)
}

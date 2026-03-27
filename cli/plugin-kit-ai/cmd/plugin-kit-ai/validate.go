package main

import (
	"fmt"

	"github.com/plugin-kit-ai/plugin-kit-ai/cli/internal/validate"
	"github.com/spf13/cobra"
)

var validatePlatform string
var validateStrict bool

var validateCmd = &cobra.Command{
	Use:   "validate [path]",
	Short: "Validate a plugin-kit-ai project scaffold",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		report, err := validate.Validate(args[0], validatePlatform)
		if err != nil {
			return err
		}
		for _, warning := range report.Warnings {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Warning: %s\n", warning.Message)
		}
		if validateStrict && len(report.Warnings) > 0 {
			return fmt.Errorf("validation warnings treated as errors (%d warning(s))", len(report.Warnings))
		}
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Validated %s\n", args[0])
		return nil
	},
}

func init() {
	validateCmd.Flags().StringVar(&validatePlatform, "platform", "", `platform override ("codex" or "claude")`)
	validateCmd.Flags().BoolVar(&validateStrict, "strict", false, "treat validation warnings as errors")
}

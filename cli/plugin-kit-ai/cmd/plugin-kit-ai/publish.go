package main

import (
	"fmt"

	"github.com/777genius/plugin-kit-ai/cli/internal/app"
	"github.com/spf13/cobra"
)

type publishRunner interface {
	Publish(app.PluginPublishOptions) (app.PluginPublishResult, error)
}

var publishCmd = newPublishCmd(pluginService)

func newPublishCmd(runner publishRunner) *cobra.Command {
	var channel string
	var dest string
	var packageRoot string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "publish [path]",
		Short: "Publish a package target into a safe local marketplace root",
		Long: `Publish a package target into a safe local marketplace root through a channel-family workflow.

This first-class publish entrypoint is intentionally bounded to documented local/catalog-safe flows:
- codex-marketplace
- claude-marketplace

For Gemini gallery publication, use publication doctor and the documented repository/release workflow instead of a local marketplace materialization path.`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root := "."
			if len(args) == 1 {
				root = args[0]
			}
			result, err := runner.Publish(app.PluginPublishOptions{
				Root:        root,
				Channel:     channel,
				Dest:        dest,
				PackageRoot: packageRoot,
				DryRun:      dryRun,
			})
			if err != nil {
				return err
			}
			for _, line := range result.Lines {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), line)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&channel, "channel", "", `publish channel ("codex-marketplace" or "claude-marketplace")`)
	cmd.Flags().StringVar(&dest, "dest", "", "destination marketplace root directory")
	cmd.Flags().StringVar(&packageRoot, "package-root", "", "relative package root inside the destination marketplace root (default: plugins/<name>)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "preview the materialized publish result without writing changes")
	_ = cmd.MarkFlagRequired("channel")
	_ = cmd.MarkFlagRequired("dest")
	return cmd
}

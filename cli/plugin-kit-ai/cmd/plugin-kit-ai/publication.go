package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/777genius/plugin-kit-ai/cli/internal/app"
	"github.com/spf13/cobra"
)

var publicationTarget string
var publicationFormat string

var publicationCmd = newPublicationCmd(pluginService)

func newPublicationCmd(runner inspectRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "publication [path]",
		Short: "Show the publication-oriented package and channel view",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root := "."
			if len(args) == 1 {
				root = args[0]
			}
			report, warnings, err := runner.Inspect(app.PluginInspectOptions{
				Root:   root,
				Target: publicationTarget,
			})
			if err != nil {
				return err
			}
			for _, warning := range warnings {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Warning: %s\n", warning.Message)
			}
			switch strings.ToLower(strings.TrimSpace(publicationFormat)) {
			case "", "text":
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "publication %s %s api_version=%s\n",
					report.Publication.Core.Name,
					report.Publication.Core.Version,
					report.Publication.Core.APIVersion,
				)
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "packages: %d channels: %d\n",
					len(report.Publication.Packages),
					len(report.Publication.Channels),
				)
				for _, pkg := range report.Publication.Packages {
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "  package[%s]: family=%s channels=%s inputs=%d managed=%d\n",
						pkg.Target,
						pkg.PackageFamily,
						strings.Join(pkg.ChannelFamilies, ","),
						len(pkg.AuthoredInputs),
						len(pkg.ManagedArtifacts),
					)
				}
				for _, channel := range report.Publication.Channels {
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "  channel[%s]: path=%s targets=%s",
						channel.Family,
						channel.Path,
						strings.Join(channel.PackageTargets, ","),
					)
					if details := inspectChannelDetails(channel.Details); details != "" {
						_, _ = fmt.Fprintf(cmd.OutOrStdout(), " details=%s", details)
					}
					_, _ = fmt.Fprintln(cmd.OutOrStdout())
				}
				return nil
			case "json":
				out, err := json.MarshalIndent(report.Publication, "", "  ")
				if err != nil {
					return err
				}
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), string(out))
				return nil
			default:
				return fmt.Errorf("unsupported format %q (use text or json)", publicationFormat)
			}
		},
	}
	cmd.Flags().StringVar(&publicationTarget, "target", "all", `publication target ("all", "claude", "codex-package", or "gemini")`)
	cmd.Flags().StringVar(&publicationFormat, "format", "text", "output format: text or json")
	return cmd
}

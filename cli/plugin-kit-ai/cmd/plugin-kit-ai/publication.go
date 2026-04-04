package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/777genius/plugin-kit-ai/cli/internal/app"
	"github.com/777genius/plugin-kit-ai/cli/internal/exitx"
	"github.com/777genius/plugin-kit-ai/cli/internal/pluginmanifest"
	"github.com/777genius/plugin-kit-ai/cli/internal/publicationmodel"
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
	cmd.AddCommand(newPublicationDoctorCmd(runner))
	return cmd
}

func newPublicationDoctorCmd(runner inspectRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "doctor [path]",
		Short:         "Inspect publication readiness without mutating files",
		Long:          "Read-only publication readiness check for package-capable targets and authored publish/... channels.",
		Args:          cobra.MaximumNArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
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
			diagnosis := diagnosePublication(report)
			for _, line := range diagnosis.Lines {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), line)
			}
			if diagnosis.Ready {
				return nil
			}
			return exitx.Wrap(errors.New("publication doctor found issues"), 1)
		},
	}
	cmd.Flags().StringVar(&publicationTarget, "target", "all", `publication target ("all", "claude", "codex-package", or "gemini")`)
	return cmd
}

type publicationDiagnosis struct {
	Ready bool
	Lines []string
}

func diagnosePublication(report pluginmanifest.Inspection) publicationDiagnosis {
	lines := []string{
		fmt.Sprintf("Publication: %s %s api_version=%s", report.Publication.Core.Name, report.Publication.Core.Version, report.Publication.Core.APIVersion),
		fmt.Sprintf("Packages: %d", len(report.Publication.Packages)),
		fmt.Sprintf("Channels: %d", len(report.Publication.Channels)),
	}
	if len(report.Publication.Packages) == 0 {
		lines = append(lines,
			"Status: inactive (no publication-capable package targets enabled)",
			"Next:",
			"  enable at least one package-capable target: claude, codex-package, or gemini",
		)
		return publicationDiagnosis{Ready: false, Lines: lines}
	}

	channelTargets := map[string]struct{}{}
	for _, channel := range report.Publication.Channels {
		for _, target := range channel.PackageTargets {
			channelTargets[target] = struct{}{}
		}
		line := fmt.Sprintf("Channel[%s]: path=%s targets=%s", channel.Family, channel.Path, strings.Join(channel.PackageTargets, ","))
		if details := inspectChannelDetails(channel.Details); details != "" {
			line += " details=" + details
		}
		lines = append(lines, line)
	}

	var missing []publicationmodel.Package
	for _, pkg := range report.Publication.Packages {
		lines = append(lines, fmt.Sprintf("Package[%s]: family=%s channels=%s managed=%d",
			pkg.Target,
			pkg.PackageFamily,
			strings.Join(pkg.ChannelFamilies, ","),
			len(pkg.ManagedArtifacts),
		))
		if _, ok := channelTargets[pkg.Target]; !ok {
			missing = append(missing, pkg)
		}
	}
	if len(missing) == 0 {
		lines = append(lines,
			"Status: ready (every publication-capable package target has an authored publication channel)",
			"Next:",
			"  run plugin-kit-ai validate . --strict",
			"  run plugin-kit-ai publication . --format json for CI or automation handoff",
		)
		return publicationDiagnosis{Ready: true, Lines: lines}
	}

	lines = append(lines, "Status: needs_channels (one or more publication-capable package targets have no authored publish/... channel)")
	lines = append(lines, "Next:")
	for _, step := range publicationNextStepsForMissing(missing) {
		lines = append(lines, "  "+step)
	}
	return publicationDiagnosis{Ready: false, Lines: lines}
}

func publicationNextStepsForMissing(missing []publicationmodel.Package) []string {
	stepSet := map[string]struct{}{}
	var steps []string
	for _, pkg := range missing {
		var step string
		switch pkg.Target {
		case "codex-package":
			step = "add publish/codex/marketplace.yaml, then rerun plugin-kit-ai render . and plugin-kit-ai validate . --strict"
		case "claude":
			step = "add publish/claude/marketplace.yaml, then rerun plugin-kit-ai render . and plugin-kit-ai validate . --strict"
		case "gemini":
			step = "add publish/gemini/gallery.yaml, keep gemini-extension.json in the repository or release root, then rerun plugin-kit-ai validate . --strict"
		default:
			continue
		}
		if _, ok := stepSet[step]; ok {
			continue
		}
		stepSet[step] = struct{}{}
		steps = append(steps, step)
	}
	slices.Sort(steps)
	return steps
}

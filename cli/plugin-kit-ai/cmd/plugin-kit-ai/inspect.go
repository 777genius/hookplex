package main

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/777genius/plugin-kit-ai/cli/internal/app"
	"github.com/spf13/cobra"
)

var inspectTarget string
var inspectFormat string

var inspectCmd = &cobra.Command{
	Use:   "inspect [path]",
	Short: "Inspect the discovered package graph and target coverage",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		root := "."
		if len(args) == 1 {
			root = args[0]
		}
		report, warnings, err := pluginService.Inspect(app.PluginInspectOptions{
			Root:   root,
			Target: inspectTarget,
		})
		if err != nil {
			return err
		}
		for _, warning := range warnings {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Warning: %s\n", warning.Message)
		}
		switch strings.ToLower(strings.TrimSpace(inspectFormat)) {
		case "", "text":
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "package %s %s\n", report.Manifest.Name, report.Manifest.Version)
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "targets: %s\n", strings.Join(report.Manifest.Targets, ", "))
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "portable: skills=%d mcp=%t\n", len(report.Portable.Paths("skills")), report.Portable.MCP != nil)
			for _, target := range report.Targets {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "- %s: class=%s production=%s runtime=%s native=%s managed=%s\n",
					target.Target,
					target.TargetClass,
					target.ProductionClass,
					target.RuntimeContract,
					strings.Join(target.TargetNativeKinds, ","),
					strings.Join(target.ManagedArtifacts, ","),
				)
				if len(target.NativeDocPaths) > 0 {
					var docs []string
					for _, kind := range target.TargetNativeKinds {
						if path := strings.TrimSpace(target.NativeDocPaths[kind]); path != "" {
							docs = append(docs, kind+"="+path)
						}
					}
					var remainingKinds []string
					for kind := range target.NativeDocPaths {
						remainingKinds = append(remainingKinds, kind)
					}
					slices.Sort(remainingKinds)
					for _, kind := range remainingKinds {
						path := target.NativeDocPaths[kind]
						if strings.TrimSpace(path) == "" || containsInspectDoc(docs, kind) {
							continue
						}
						docs = append(docs, kind+"="+path)
					}
					if len(docs) > 0 {
						_, _ = fmt.Fprintf(cmd.OutOrStdout(), "  docs=%s\n", strings.Join(docs, ","))
					}
				}
				if len(target.UnsupportedKinds) > 0 {
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "  unsupported=%s\n", strings.Join(target.UnsupportedKinds, ","))
				}
				if len(target.NativeSurfaces) > 0 {
					var tiers []string
					for _, surface := range target.NativeSurfaces {
						tiers = append(tiers, surface.Kind+"="+surface.Tier)
					}
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "  surfaces=%s\n", strings.Join(tiers, ","))
				}
			}
			return nil
		case "json":
			out, err := json.MarshalIndent(report, "", "  ")
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), string(out))
			return nil
		default:
			return fmt.Errorf("unsupported format %q (use text or json)", inspectFormat)
		}
	},
}

func containsInspectDoc(items []string, kind string) bool {
	prefix := kind + "="
	for _, item := range items {
		if strings.HasPrefix(item, prefix) {
			return true
		}
	}
	return false
}

func init() {
	inspectCmd.Flags().StringVar(&inspectTarget, "target", "all", `inspect target ("all", "claude", "codex-package", "codex-runtime", "gemini", "opencode")`)
	inspectCmd.Flags().StringVar(&inspectFormat, "format", "text", "output format: text or json")
}

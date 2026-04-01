package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/777genius/plugin-kit-ai/cli/internal/capabilities"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

type docsManifestEntry struct {
	CommandPath string   `json:"command_path"`
	Slug        string   `json:"slug"`
	FileName    string   `json:"file_name"`
	Use         string   `json:"use"`
	Short       string   `json:"short,omitempty"`
	Long        string   `json:"long,omitempty"`
	Aliases     []string `json:"aliases,omitempty"`
	Deprecated  bool     `json:"deprecated"`
	Hidden      bool     `json:"hidden"`
}

var docsCmd = &cobra.Command{
	Use:    "__docs",
	Short:  "internal docs export helpers",
	Hidden: true,
}

var docsCLIExportCmd = &cobra.Command{
	Use:    "export-cli",
	Short:  "internal Cobra docs export",
	Hidden: true,
	RunE: func(_ *cobra.Command, _ []string) error {
		if strings.TrimSpace(docsCLIExportDir) == "" {
			return fmt.Errorf("--out-dir is required")
		}
		if strings.TrimSpace(docsCLIManifestPath) == "" {
			return fmt.Errorf("--manifest-path is required")
		}
		root := docsRootForExport()
		disableAutoGenTag(root)
		if err := os.MkdirAll(docsCLIExportDir, 0o755); err != nil {
			return err
		}
		if err := doc.GenMarkdownTree(root, docsCLIExportDir); err != nil {
			return err
		}
		manifest := visibleCommandManifest(root)
		body, err := json.MarshalIndent(manifest, "", "  ")
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(docsCLIManifestPath), 0o755); err != nil {
			return err
		}
		return os.WriteFile(docsCLIManifestPath, append(body, '\n'), 0o644)
	},
}

var docsSupportExportCmd = &cobra.Command{
	Use:    "export-support",
	Short:  "internal support and capability export",
	Hidden: true,
	RunE: func(_ *cobra.Command, _ []string) error {
		if err := writeJSON(docsEventsPath, capabilities.All()); err != nil {
			return err
		}
		if err := writeJSON(docsTargetsPath, capabilities.TargetAll()); err != nil {
			return err
		}
		return writeJSON(docsCapabilitiesPath, uniqueCapabilities(capabilities.All()))
	},
}

var (
	docsCLIExportDir     string
	docsCLIManifestPath  string
	docsEventsPath       string
	docsTargetsPath      string
	docsCapabilitiesPath string
)

func init() {
	docsCLIExportCmd.Flags().StringVar(&docsCLIExportDir, "out-dir", "", "directory for generated markdown output")
	docsCLIExportCmd.Flags().StringVar(&docsCLIManifestPath, "manifest-path", "", "path for the generated command manifest json")
	docsSupportExportCmd.Flags().StringVar(&docsEventsPath, "events-path", "", "path for support event json")
	docsSupportExportCmd.Flags().StringVar(&docsTargetsPath, "targets-path", "", "path for target support json")
	docsSupportExportCmd.Flags().StringVar(&docsCapabilitiesPath, "capabilities-path", "", "path for capability summary json")
	docsCmd.AddCommand(docsCLIExportCmd)
	docsCmd.AddCommand(docsSupportExportCmd)
	rootCmd.AddCommand(docsCmd)
}

func docsRootForExport() *cobra.Command {
	return rootCmd
}

func disableAutoGenTag(cmd *cobra.Command) {
	cmd.DisableAutoGenTag = true
	for _, child := range cmd.Commands() {
		disableAutoGenTag(child)
	}
}

func visibleCommandManifest(root *cobra.Command) []docsManifestEntry {
	out := make([]docsManifestEntry, 0)
	var walk func(*cobra.Command)
	walk = func(cmd *cobra.Command) {
		if !cmd.Hidden {
			out = append(out, docsManifestEntry{
				CommandPath: cmd.CommandPath(),
				Slug:        commandSlug(cmd),
				FileName:    commandMarkdownFile(cmd),
				Use:         cmd.Use,
				Short:       cmd.Short,
				Long:        strings.TrimSpace(cmd.Long),
				Aliases:     append([]string(nil), cmd.Aliases...),
				Deprecated:  strings.TrimSpace(cmd.Deprecated) != "",
				Hidden:      cmd.Hidden,
			})
		}
		for _, child := range cmd.Commands() {
			if child.Hidden {
				continue
			}
			walk(child)
		}
	}
	walk(root)
	return out
}

func commandMarkdownFile(cmd *cobra.Command) string {
	return strings.ReplaceAll(cmd.CommandPath(), " ", "_") + ".md"
}

func commandSlug(cmd *cobra.Command) string {
	return strings.ReplaceAll(strings.ToLower(cmd.CommandPath()), " ", "-")
}

func writeJSON(path string, value any) error {
	if strings.TrimSpace(path) == "" {
		return fmt.Errorf("output path is required")
	}
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o644)
}

func uniqueCapabilities(entries []capabilities.Entry) []string {
	seen := map[string]struct{}{}
	out := make([]string, 0)
	for _, entry := range entries {
		for _, capability := range entry.Capabilities {
			if _, ok := seen[capability]; ok {
				continue
			}
			seen[capability] = struct{}{}
			out = append(out, capability)
		}
	}
	slices.Sort(out)
	return out
}

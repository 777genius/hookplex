package platformexec

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/plugin-kit-ai/plugin-kit-ai/cli/internal/pluginmodel"
)

type claudeAdapter struct{}

func (claudeAdapter) ID() string { return "claude" }

func (claudeAdapter) DetectNative(root string) bool {
	return fileExists(filepath.Join(root, ".claude-plugin", "plugin.json")) || fileExists(filepath.Join(root, "hooks", "hooks.json"))
}

func (claudeAdapter) RefineDiscovery(root string, state *pluginmodel.TargetState) error {
	if rel := state.DocPath("package_metadata"); strings.TrimSpace(rel) != "" {
		var discard map[string]any
		if _, ok, err := readYAMLDoc[map[string]any](root, rel); err != nil {
			return fmt.Errorf("parse %s: %w", rel, err)
		} else if ok {
			_ = discard
		}
	}
	return nil
}

func (claudeAdapter) Import(root string, seed ImportSeed) (ImportResult, error) {
	result := ImportResult{
		Manifest: seed.Manifest,
		Launcher: seed.Launcher,
	}
	type meta struct {
		Name        string `json:"name"`
		Version     string `json:"version"`
		Description string `json:"description"`
	}
	if body, err := os.ReadFile(filepath.Join(root, ".claude-plugin", "plugin.json")); err == nil {
		var m meta
		if json.Unmarshal(body, &m) == nil {
			if strings.TrimSpace(m.Name) != "" {
				result.Manifest.Name = m.Name
			}
			if strings.TrimSpace(m.Version) != "" {
				result.Manifest.Version = m.Version
			}
			if strings.TrimSpace(m.Description) != "" {
				result.Manifest.Description = m.Description
			}
		}
	}
	if body, err := os.ReadFile(filepath.Join(root, "hooks", "hooks.json")); err == nil && result.Launcher != nil {
		if entrypoint, ok := inferClaudeEntrypoint(body); ok {
			result.Launcher.Entrypoint = entrypoint
		}
	}
	copied, err := copySingleArtifactIfExists(root, filepath.Join("hooks", "hooks.json"), filepath.Join("targets", "claude", "hooks", "hooks.json"))
	if err != nil {
		return ImportResult{}, err
	}
	result.Artifacts = compactArtifacts(copied)
	return result, nil
}

func (claudeAdapter) Render(root string, graph pluginmodel.PackageGraph, state pluginmodel.TargetState) ([]pluginmodel.Artifact, error) {
	entrypoint := ""
	if graph.Launcher != nil {
		entrypoint = graph.Launcher.Entrypoint
	}
	if strings.TrimSpace(entrypoint) == "" {
		return nil, fmt.Errorf("required launcher missing: %s", pluginmodel.LauncherFileName)
	}
	artifacts, err := renderManagedPluginArtifacts(graph.Manifest.Name, graph.Manifest, graph.Portable, true, filepath.Join(".claude-plugin", "plugin.json"), pluginmodel.NativeExtraDoc{}, "", nil)
	if err != nil {
		return nil, err
	}
	if hookPaths := state.ComponentPaths("hooks"); len(hookPaths) > 0 {
		copied, err := copyArtifacts(root, filepath.Join("targets", "claude", "hooks"), "hooks")
		if err != nil {
			return nil, err
		}
		artifacts = append(artifacts, copied...)
	} else {
		artifacts = append(artifacts, pluginmodel.Artifact{
			RelPath: filepath.Join("hooks", "hooks.json"),
			Content: defaultClaudeHooks(entrypoint),
		})
	}
	copiedKinds := []artifactDir{
		{src: filepath.Join("targets", "claude", "commands"), dst: "commands"},
		{src: filepath.Join("targets", "claude", "contexts"), dst: "contexts"},
	}
	copied, err := copyArtifactDirs(root, copiedKinds...)
	if err != nil {
		return nil, err
	}
	return append(artifacts, copied...), nil
}

func (claudeAdapter) ManagedPaths(root string, graph pluginmodel.PackageGraph, state pluginmodel.TargetState) ([]string, error) {
	return nil, nil
}

func (claudeAdapter) Validate(root string, graph pluginmodel.PackageGraph, state pluginmodel.TargetState) ([]Diagnostic, error) {
	if graph.Launcher == nil {
		return nil, nil
	}
	var diagnostics []Diagnostic
	for _, rel := range state.ComponentPaths("hooks") {
		full := filepath.Join(root, rel)
		body, err := os.ReadFile(full)
		if err != nil {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: SeverityFailure,
				Code:     CodeManifestInvalid,
				Path:     rel,
				Target:   "claude",
				Message:  fmt.Sprintf("Claude hooks file %s is not readable: %v", rel, err),
			})
			continue
		}
		mismatches, err := validateClaudeHookEntrypoints(body, graph.Launcher.Entrypoint)
		if err != nil {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: SeverityFailure,
				Code:     CodeManifestInvalid,
				Path:     rel,
				Target:   "claude",
				Message:  fmt.Sprintf("Claude hooks file %s is invalid JSON: %v", rel, err),
			})
			continue
		}
		for _, mismatch := range mismatches {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: SeverityFailure,
				Code:     CodeEntrypointMismatch,
				Path:     rel,
				Target:   "claude",
				Message:  mismatch,
			})
		}
	}
	return diagnostics, nil
}

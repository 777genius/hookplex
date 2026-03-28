package app

import (
	"fmt"
	"strings"

	"github.com/plugin-kit-ai/plugin-kit-ai/cli/internal/pluginmanifest"
	"github.com/plugin-kit-ai/plugin-kit-ai/cli/internal/runtimecheck"
)

type PluginDoctorOptions struct {
	Root string
}

type PluginDoctorResult struct {
	Ready bool
	Lines []string
}

func (PluginService) Doctor(opts PluginDoctorOptions) (PluginDoctorResult, error) {
	root := strings.TrimSpace(opts.Root)
	if root == "" {
		root = "."
	}
	graph, _, err := pluginmanifest.Discover(root)
	if err != nil {
		return PluginDoctorResult{}, err
	}
	project, err := runtimecheck.Inspect(runtimecheck.Inputs{
		Root:     root,
		Targets:  graph.Manifest.EnabledTargets(),
		Launcher: graph.Launcher,
	})
	if err != nil {
		return PluginDoctorResult{}, err
	}
	diagnosis := runtimecheck.Diagnose(project)
	lines := []string{
		project.ProjectLine(),
		fmt.Sprintf("Status: %s (%s)", diagnosis.Status, diagnosis.Reason),
	}
	if len(diagnosis.Next) > 0 {
		lines = append(lines, "Next:")
		for _, step := range diagnosis.Next {
			lines = append(lines, "  "+step)
		}
	}
	return PluginDoctorResult{
		Ready: diagnosis.Status == runtimecheck.StatusReady,
		Lines: lines,
	}, nil
}

package app

import (
	"fmt"
	"strings"

	"github.com/777genius/plugin-kit-ai/cli/internal/pluginmanifest"
	"github.com/777genius/plugin-kit-ai/cli/internal/runtimecheck"
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
	if requirement := exportRuntimeRequirement(project.Runtime); strings.TrimSpace(requirement) != "" {
		lines = append(lines, "Runtime requirement: "+requirement)
	}
	if hint := exportRuntimeInstallHint(project.Runtime); strings.TrimSpace(hint) != "" {
		lines = append(lines, "Runtime install hint: "+hint)
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

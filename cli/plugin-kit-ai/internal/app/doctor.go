package app

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
	lines = append(lines, doctorEnvironmentLines(root, project)...)
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

type doctorToolSpec struct {
	Label       string
	Commands    []string
	VersionArgs []string
}

func doctorEnvironmentLines(root string, project runtimecheck.Project) []string {
	specs := doctorToolSpecs(root, project)
	if len(specs) == 0 {
		return nil
	}

	lines := []string{"Environment:"}
	missing := false
	for _, spec := range specs {
		path, _, err := doctorFindBinary(spec.Commands)
		if err != nil {
			lines = append(lines, "  "+doctorMissingLine(spec))
			missing = true
			continue
		}
		line := fmt.Sprintf("  %s: ok (%s", spec.Label, path)
		if version := doctorVersion(root, path, spec.VersionArgs); version != "" {
			line += "; " + version
		}
		line += ")"
		lines = append(lines, line)
	}
	if missing {
		lines = append(lines, "  Hint: "+doctorPATHHint())
	}
	return lines
}

func doctorToolSpecs(root string, project runtimecheck.Project) []doctorToolSpec {
	var specs []doctorToolSpec
	if fileExists(filepath.Join(root, "go.mod")) || strings.TrimSpace(project.Runtime) == "go" {
		specs = appendDoctorToolSpec(specs,
			doctorToolSpec{Label: "go", Commands: []string{"go"}, VersionArgs: []string{"version"}},
			doctorToolSpec{Label: "gofmt", Commands: []string{"gofmt"}},
		)
	}

	switch strings.TrimSpace(project.Runtime) {
	case "python":
		specs = appendDoctorToolSpec(specs, doctorToolSpec{
			Label:       "python runtime",
			Commands:    pythonPathNames(),
			VersionArgs: []string{"--version"},
		})
		switch project.Python.Manager {
		case runtimecheck.PythonManagerUV:
			specs = appendDoctorToolSpec(specs, doctorToolSpec{Label: "uv", Commands: []string{"uv"}, VersionArgs: []string{"--version"}})
		case runtimecheck.PythonManagerPoetry:
			specs = appendDoctorToolSpec(specs, doctorToolSpec{Label: "poetry", Commands: []string{"poetry"}, VersionArgs: []string{"--version"}})
		case runtimecheck.PythonManagerPipenv:
			specs = appendDoctorToolSpec(specs, doctorToolSpec{Label: "pipenv", Commands: []string{"pipenv"}, VersionArgs: []string{"--version"}})
		}
	case "node":
		specs = appendDoctorToolSpec(specs, doctorToolSpec{
			Label:       "node",
			Commands:    []string{"node"},
			VersionArgs: []string{"--version"},
		})
		manager := strings.TrimSpace(project.Node.ManagerBinary)
		if manager != "" && manager != "node" {
			specs = appendDoctorToolSpec(specs, doctorToolSpec{
				Label:       manager,
				Commands:    []string{manager},
				VersionArgs: []string{"--version"},
			})
		}
	case "shell":
		if runtime.GOOS == "windows" {
			specs = appendDoctorToolSpec(specs, doctorToolSpec{Label: "bash", Commands: []string{"bash"}, VersionArgs: []string{"--version"}})
		}
	}

	return specs
}

func appendDoctorToolSpec(dst []doctorToolSpec, specs ...doctorToolSpec) []doctorToolSpec {
	for _, spec := range specs {
		if strings.TrimSpace(spec.Label) == "" || len(spec.Commands) == 0 {
			continue
		}
		duplicate := false
		for _, existing := range dst {
			if existing.Label == spec.Label {
				duplicate = true
				break
			}
		}
		if !duplicate {
			dst = append(dst, spec)
		}
	}
	return dst
}

func doctorFindBinary(commands []string) (string, string, error) {
	for _, command := range commands {
		command = strings.TrimSpace(command)
		if command == "" {
			continue
		}
		path, err := runtimecheck.LookPath(command)
		if err == nil {
			return path, command, nil
		}
	}
	return "", "", os.ErrNotExist
}

func doctorMissingLine(spec doctorToolSpec) string {
	if len(spec.Commands) == 1 {
		return fmt.Sprintf("%s: missing from PATH", spec.Label)
	}
	return fmt.Sprintf("%s: missing from PATH (checked: %s)", spec.Label, strings.Join(spec.Commands, ", "))
}

func doctorVersion(root, path string, args []string) string {
	if len(args) == 0 {
		return ""
	}
	out, err := runtimecheck.RunCommand(root, path, args...)
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			return line
		}
	}
	return ""
}

func doctorPATHHint() string {
	hint := "if the runtime is installed but doctor cannot see it here, check PATH for non-interactive shells"
	if runtime.GOOS == "darwin" {
		hint += " (for example ~/.zshenv on macOS)"
	}
	if missing := doctorMissingCommonPATHDirs(); len(missing) > 0 {
		hint += "; common directories missing from PATH: " + strings.Join(missing, ", ")
	}
	return hint + "."
}

func doctorMissingCommonPATHDirs() []string {
	candidates := []string{"/usr/local/bin"}
	switch runtime.GOOS {
	case "darwin":
		candidates = append([]string{"/opt/homebrew/bin", "/usr/local/go/bin"}, candidates...)
	case "linux":
		candidates = append([]string{"/usr/local/sbin"}, candidates...)
	}

	current := map[string]struct{}{}
	for _, entry := range filepath.SplitList(os.Getenv("PATH")) {
		entry = strings.TrimSpace(entry)
		if entry != "" {
			current[entry] = struct{}{}
		}
	}

	var missing []string
	for _, candidate := range candidates {
		if _, ok := current[candidate]; !ok {
			missing = append(missing, candidate)
		}
	}
	return missing
}

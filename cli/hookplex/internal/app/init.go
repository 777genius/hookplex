package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hookplex/hookplex/cli/internal/scaffold"
)

// InitOptions is parsed CLI state for hookplex init.
type InitOptions struct {
	ProjectName string
	Platform    string
	OutputDir   string // empty → ./<project-name> under cwd
	Force       bool
	Extras      bool
}

// InitRunner runs hookplex init.
type InitRunner struct{}

// Run validates options, writes scaffold files, and returns the absolute output directory.
func (InitRunner) Run(opts InitOptions) (outDir string, err error) {
	name := strings.TrimSpace(opts.ProjectName)
	if err := scaffold.ValidateProjectName(name); err != nil {
		return "", err
	}

	p := strings.ToLower(strings.TrimSpace(opts.Platform))
	if _, ok := scaffold.LookupPlatform(p); !ok {
		return "", errUnknownPlatform(opts.Platform)
	}

	out := strings.TrimSpace(opts.OutputDir)
	if out == "" {
		wd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("get working directory: %w", err)
		}
		out = filepath.Join(wd, name)
	} else {
		abs, err := filepath.Abs(out)
		if err != nil {
			return "", fmt.Errorf("resolve output path: %w", err)
		}
		out = abs
	}

	d := scaffold.Data{
		ProjectName: name,
		ModulePath:  scaffold.DefaultModulePath(name),
		Description: "hookplex plugin",
		Platform:    p,
		WithExtras:  opts.Extras,
	}
	if err := scaffold.Write(out, d, opts.Force); err != nil {
		return "", err
	}
	return out, nil
}

func errUnknownPlatform(platform string) error {
	return &unknownPlatformError{platform: platform}
}

type unknownPlatformError struct {
	platform string
}

func (e *unknownPlatformError) Error() string {
	return "unknown platform " + `"` + e.platform + `"`
}

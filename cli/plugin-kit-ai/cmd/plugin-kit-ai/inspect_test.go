package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/777genius/plugin-kit-ai/cli/internal/app"
	"github.com/777genius/plugin-kit-ai/cli/internal/pluginmanifest"
)

type fakeInspectRunner struct {
	report   pluginmanifest.Inspection
	warnings []pluginmanifest.Warning
	err      error
}

func (f fakeInspectRunner) Inspect(app.PluginInspectOptions) (pluginmanifest.Inspection, []pluginmanifest.Warning, error) {
	return f.report, f.warnings, f.err
}

func TestInspectTextShowsLauncherAndGeminiGuidance(t *testing.T) {
	t.Parallel()
	cmd := newInspectCmd(fakeInspectRunner{
		report: pluginmanifest.Inspection{
			Manifest: pluginmanifest.Manifest{
				Name:    "demo",
				Version: "0.1.0",
				Targets: []string{"gemini"},
			},
			Launcher: &pluginmanifest.Launcher{
				Runtime:    "go",
				Entrypoint: "./bin/demo",
			},
			Targets: []pluginmanifest.InspectTarget{{
				Target:            "gemini",
				TargetClass:       "mcp_extension",
				ProductionClass:   "runtime-supported beta extension target",
				RuntimeContract:   "Gemini Go runtime beta lane plus full extension packaging lane; not production-ready",
				TargetNativeKinds: []string{"hooks", "contexts"},
				ManagedArtifacts:  []string{"gemini-extension.json", "hooks/hooks.json"},
			}},
		},
	})
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"--format", "text", "."})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	output := buf.String()
	for _, want := range []string{
		"launcher: runtime=go entrypoint=./bin/demo",
		"next=go test ./...; plugin-kit-ai render --check .; plugin-kit-ai validate . --platform gemini --strict; gemini extensions link .",
		"live_smoke=make test-gemini-runtime-live",
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("inspect output missing %q:\n%s", want, output)
		}
	}
}

func TestInspectTextShowsGeminiPackagingGuidanceWithoutLauncher(t *testing.T) {
	t.Parallel()
	cmd := newInspectCmd(fakeInspectRunner{
		report: pluginmanifest.Inspection{
			Manifest: pluginmanifest.Manifest{
				Name:    "demo",
				Version: "0.1.0",
				Targets: []string{"gemini"},
			},
			Targets: []pluginmanifest.InspectTarget{{
				Target:            "gemini",
				TargetClass:       "mcp_extension",
				ProductionClass:   "runtime-supported beta extension target",
				RuntimeContract:   "Gemini Go runtime beta lane plus full extension packaging lane; not production-ready",
				TargetNativeKinds: []string{"commands", "contexts"},
				ManagedArtifacts:  []string{"gemini-extension.json"},
			}},
		},
	})
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"--format", "text", "."})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	output := buf.String()
	for _, want := range []string{
		"managed=gemini-extension.json",
		"next=render --check + validate --strict keep the packaging lane honest; add --runtime go only when you intentionally want the beta hook lane",
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("inspect output missing %q:\n%s", want, output)
		}
	}
	if strings.Contains(output, "launcher: runtime=") {
		t.Fatalf("inspect output unexpectedly shows launcher:\n%s", output)
	}
}

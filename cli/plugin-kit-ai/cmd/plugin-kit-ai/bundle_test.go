package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/plugin-kit-ai/plugin-kit-ai/cli/internal/app"
)

type fakeBundleInstallRunner struct {
	result app.PluginBundleInstallResult
	err    error
}

func (f fakeBundleInstallRunner) BundleInstall(app.PluginBundleInstallOptions) (app.PluginBundleInstallResult, error) {
	return f.result, f.err
}

func TestBundleInstallHelpIncludesLocalTarballLanguage(t *testing.T) {
	t.Parallel()
	cmd := newBundleCmd(fakeBundleInstallRunner{})
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"install", "--help"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	output := buf.String()
	for _, want := range []string{"local .tar.gz", "Python/Node", "binary-only"} {
		if !strings.Contains(output, want) {
			t.Fatalf("help output missing %q:\n%s", want, output)
		}
	}
}

func TestBundleInstallWritesRunnerOutput(t *testing.T) {
	t.Parallel()
	cmd := newBundleCmd(fakeBundleInstallRunner{
		result: app.PluginBundleInstallResult{
			Lines: []string{
				"Bundle: plugin=demo platform=codex-runtime runtime=python manager=requirements.txt (pip)",
				"Installed path: /tmp/demo",
			},
		},
	})
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"install", "--dest", "/tmp/demo", "/tmp/demo.tar.gz"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	output := buf.String()
	if !strings.Contains(output, "Installed path: /tmp/demo") {
		t.Fatalf("output = %s", output)
	}
}

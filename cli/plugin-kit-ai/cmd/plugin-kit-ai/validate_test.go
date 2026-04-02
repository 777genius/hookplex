package main

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/777genius/plugin-kit-ai/cli/internal/validate"
)

type fakeValidateRunner struct {
	report validate.Report
	err    error
}

func (f fakeValidateRunner) Run(_ string, _ string) (validate.Report, error) {
	return f.report, f.err
}

func TestValidateWritesJSONOutput(t *testing.T) {
	t.Parallel()
	cmd := newValidateCmd(fakeValidateRunner{
		report: validate.Report{
			Platform: "codex-runtime",
			Checks:   []string{"plugin_manifest"},
			Warnings: []validate.Warning{{
				Kind:    validate.WarningManifestUnknownField,
				Path:    "plugin.yaml",
				Message: "unknown plugin.yaml field: extra_field",
			}},
		},
	}.Run)
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"--platform", "codex-runtime", "--format", "json", "."})
	if err := cmd.ExecuteContext(context.Background()); err != nil {
		t.Fatal(err)
	}
	output := buf.String()
	for _, want := range []string{
		`"platform": "codex-runtime"`,
		`"checks": [`,
		`"warnings": [`,
		`"failures": []`,
		`"path": "plugin.yaml"`,
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("json output missing %q:\n%s", want, output)
		}
	}
}

func TestValidateJSONPrintsReportOnFailure(t *testing.T) {
	t.Parallel()
	report := validate.Report{
		Checks: []string{},
		Failures: []validate.Failure{{
			Kind:    validate.FailureManifestMissing,
			Path:    "plugin.yaml",
			Message: "required manifest missing: plugin.yaml",
		}},
	}
	cmd := newValidateCmd(fakeValidateRunner{
		err: &validate.ReportError{Report: report},
	}.Run)
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"--format", "json", "."})
	err := cmd.ExecuteContext(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(buf.String(), `"kind": "manifest_missing"`) {
		t.Fatalf("json output missing failure payload:\n%s", buf.String())
	}
	if !strings.Contains(buf.String(), `"failures": [`) {
		t.Fatalf("json output missing failures array:\n%s", buf.String())
	}
}

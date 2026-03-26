package validate

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidate_CannotInferPlatform(t *testing.T) {
	t.Parallel()
	_, err := Validate(t.TempDir(), "")
	var re *ReportError
	if !errors.As(err, &re) {
		t.Fatalf("expected ReportError, got %v", err)
	}
	if got := re.Report.Failures[0].Kind; got != FailureCannotInferPlatform {
		t.Fatalf("failure kind = %q", got)
	}
	if re.Error() != "could not infer platform" {
		t.Fatalf("error = %q", re.Error())
	}
}

func TestValidate_UnknownPlatform(t *testing.T) {
	t.Parallel()
	_, err := Validate(t.TempDir(), "nope")
	var re *ReportError
	if !errors.As(err, &re) {
		t.Fatalf("expected ReportError, got %v", err)
	}
	if got := re.Report.Failures[0].Kind; got != FailureUnknownPlatform {
		t.Fatalf("failure kind = %q", got)
	}
	if re.Error() != "unknown platform \"nope\"" {
		t.Fatalf("error = %q", re.Error())
	}
}

func TestValidate_RequiredFileMissing(t *testing.T) {
	t.Parallel()
	report, err := Validate(t.TempDir(), "codex")
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Failures) == 0 || report.Failures[0].Kind != FailureRequiredFileMissing {
		t.Fatalf("failures = %+v", report.Failures)
	}
}

func TestValidate_ForbiddenFilePresent(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	mustWriteValidateFile(t, dir, "go.mod", "module example.com/x\n\ngo 1.22\n")
	mustWriteValidateFile(t, dir, "README.md", "# x\n")
	mustWriteValidateFile(t, dir, "AGENTS.md", "repo instructions\n")
	mustWriteValidateFile(t, dir, filepath.Join(".codex", "config.toml"), "notify = [\"./bin/x\", \"notify\"]\n")
	mustWriteValidateFile(t, dir, filepath.Join(".claude-plugin", "plugin.json"), "{}\n")

	report, err := Validate(dir, "codex")
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Failures) == 0 || report.Failures[0].Kind != FailureForbiddenFilePresent {
		t.Fatalf("failures = %+v", report.Failures)
	}
	if !strings.Contains(report.Failures[0].Message, ".claude-plugin/plugin.json") {
		t.Fatalf("message = %q", report.Failures[0].Message)
	}
}

func TestValidate_BuildFailed(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	mustWriteValidateFile(t, dir, "go.mod", "module example.com/x\n\ngo 1.22\n")
	mustWriteValidateFile(t, dir, "README.md", "# x\n")
	mustWriteValidateFile(t, dir, "AGENTS.md", "repo instructions\n")
	mustWriteValidateFile(t, dir, filepath.Join(".codex", "config.toml"), "notify = [\"./bin/x\", \"notify\"]\n")
	mustWriteValidateFile(t, dir, "broken.go", "package main\nfunc main() {\n")

	report, err := Validate(dir, "codex")
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Failures) == 0 || report.Failures[0].Kind != FailureBuildFailed {
		t.Fatalf("failures = %+v", report.Failures)
	}
	if !strings.Contains((&ReportError{Report: report}).Error(), "go build ./...:") {
		t.Fatalf("report error = %q", (&ReportError{Report: report}).Error())
	}
}

func mustWriteValidateFile(t *testing.T, root, rel, body string) {
	t.Helper()
	full := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

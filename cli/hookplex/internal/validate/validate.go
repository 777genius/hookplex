package validate

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type FailureKind string

const (
	FailureUnknownPlatform          FailureKind = "unknown_platform"
	FailureCannotInferPlatform      FailureKind = "cannot_infer_platform"
	FailureRequiredFileMissing      FailureKind = "required_file_missing"
	FailureForbiddenFilePresent     FailureKind = "forbidden_file_present"
	FailureBuildFailed              FailureKind = "build_failed"
	FailureGeneratedContractInvalid FailureKind = "generated_contract_invalid"
)

type Failure struct {
	Kind    FailureKind
	Path    string
	Target  string
	Message string
}

type Report struct {
	Platform string
	Checks   []string
	Failures []Failure
}

type ReportError struct {
	Report Report
}

func (e *ReportError) Error() string {
	if len(e.Report.Failures) == 0 {
		return "validation failed"
	}
	f := e.Report.Failures[0]
	switch f.Kind {
	case FailureRequiredFileMissing:
		return "required file missing: " + f.Path
	case FailureForbiddenFilePresent:
		return fmt.Sprintf("forbidden file present for platform %s: %s", e.Report.Platform, f.Path)
	case FailureBuildFailed:
		return fmt.Sprintf("go build %s: %s", f.Target, f.Message)
	default:
		return f.Message
	}
}

type Rule struct {
	Platform       string
	RequiredFiles  []string
	ForbiddenFiles []string
	BuildTargets   []string
}

func Run(root, platform string) error {
	report, err := Validate(root, platform)
	if err != nil {
		return err
	}
	if len(report.Failures) > 0 {
		return &ReportError{Report: report}
	}
	return nil
}

func Validate(root, platform string) (Report, error) {
	rule, err := resolveRule(root, platform)
	if err != nil {
		return Report{}, err
	}
	report := Report{
		Platform: rule.Platform,
		Checks:   []string{"required_files", "forbidden_files", "build_targets"},
	}
	for _, rel := range rule.RequiredFiles {
		full := filepath.Join(root, rel)
		if _, err := os.Stat(full); err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureRequiredFileMissing,
				Path:    rel,
				Message: "required file missing: " + rel,
			})
		}
	}
	for _, rel := range rule.ForbiddenFiles {
		full := filepath.Join(root, rel)
		if _, err := os.Stat(full); err == nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureForbiddenFilePresent,
				Path:    rel,
				Message: fmt.Sprintf("forbidden file present for platform %s: %s", rule.Platform, rel),
			})
		}
	}
	if len(report.Failures) > 0 {
		return report, nil
	}
	for _, target := range rule.BuildTargets {
		cmd := exec.Command("go", "build", target)
		cmd.Dir = root
		cmd.Env = append(os.Environ(), "GOWORK=off")
		if out, err := cmd.CombinedOutput(); err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureBuildFailed,
				Target:  target,
				Message: fmt.Sprintf("%v\n%s", err, out),
			})
		}
	}
	return report, nil
}

func resolveRule(root, platform string) (Rule, error) {
	if strings.TrimSpace(platform) != "" {
		rule, ok := LookupRule(platform)
		if !ok {
			return Rule{}, &ReportError{Report: Report{
				Failures: []Failure{{
					Kind:    FailureUnknownPlatform,
					Message: fmt.Sprintf("unknown platform %q", platform),
				}},
			}}
		}
		return rule, nil
	}
	if fileExists(filepath.Join(root, "AGENTS.md")) {
		rule, _ := LookupRule("codex")
		return rule, nil
	}
	if fileExists(filepath.Join(root, ".claude-plugin", "plugin.json")) {
		rule, _ := LookupRule("claude")
		return rule, nil
	}
	return Rule{}, &ReportError{Report: Report{
		Failures: []Failure{{
			Kind:    FailureCannotInferPlatform,
			Message: "could not infer platform",
		}},
	}}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

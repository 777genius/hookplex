package hookplexrepo_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestHookplexSkillsInitValidateRender(t *testing.T) {
	bin := buildHookplex(t)
	cases := []struct {
		name         string
		skillName    string
		template     string
		command      string
		mustExist    []string
		mustNotExist []string
	}{
		{
			name:      "go-command",
			skillName: "lint-repo",
			template:  "go-command",
			mustExist: []string{
				filepath.Join("skills", "lint-repo", "SKILL.md"),
				filepath.Join("cmd", "lint-repo", "main.go"),
				filepath.Join("generated", "skills", "claude", "lint-repo", "SKILL.md"),
				filepath.Join("generated", "skills", "codex", "lint-repo", "SKILL.md"),
				filepath.Join("generated", "skills", "codex", "lint-repo", "AGENTS.md"),
				filepath.Join("commands", "lint-repo.md"),
			},
		},
		{
			name:      "cli-wrapper",
			skillName: "format-changed",
			template:  "cli-wrapper",
			command:   "npx prettier@3.4.2 --write .",
			mustExist: []string{
				filepath.Join("skills", "format-changed", "SKILL.md"),
				filepath.Join("skills", "format-changed", "scripts", ".keep"),
				filepath.Join("generated", "skills", "claude", "format-changed", "SKILL.md"),
				filepath.Join("generated", "skills", "codex", "format-changed", "SKILL.md"),
				filepath.Join("generated", "skills", "codex", "format-changed", "AGENTS.md"),
				filepath.Join("commands", "format-changed.md"),
			},
		},
		{
			name:      "docs-only",
			skillName: "review-checklist",
			template:  "docs-only",
			mustExist: []string{
				filepath.Join("skills", "review-checklist", "SKILL.md"),
				filepath.Join("skills", "review-checklist", "references", ".keep"),
				filepath.Join("generated", "skills", "claude", "review-checklist", "SKILL.md"),
				filepath.Join("generated", "skills", "codex", "review-checklist", "SKILL.md"),
				filepath.Join("generated", "skills", "codex", "review-checklist", "AGENTS.md"),
			},
			mustNotExist: []string{
				filepath.Join("commands", "review-checklist.md"),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			args := []string{"skills", "init", tc.skillName, "--output", root, "--template", tc.template}
			if tc.command != "" {
				args = append(args, "--command", tc.command)
			}
			initCmd := exec.Command(bin, args...)
			if out, err := initCmd.CombinedOutput(); err != nil {
				t.Fatalf("hookplex skills init: %v\n%s", err, out)
			}

			validateCmd := exec.Command(bin, "skills", "validate", root)
			if out, err := validateCmd.CombinedOutput(); err != nil {
				t.Fatalf("hookplex skills validate: %v\n%s", err, out)
			}

			renderCmd := exec.Command(bin, "skills", "render", root, "--target", "all")
			if out, err := renderCmd.CombinedOutput(); err != nil {
				t.Fatalf("hookplex skills render: %v\n%s", err, out)
			}

			for _, rel := range tc.mustExist {
				if _, err := os.Stat(filepath.Join(root, rel)); err != nil {
					t.Fatalf("missing %s: %v", rel, err)
				}
			}
			for _, rel := range tc.mustNotExist {
				if _, err := os.Stat(filepath.Join(root, rel)); err == nil {
					t.Fatalf("unexpected %s", rel)
				}
			}
		})
	}
}

func TestHookplexSkillsValidateReportsMultipleProblems(t *testing.T) {
	bin := buildHookplex(t)
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "skills", "broken"), 0o755); err != nil {
		t.Fatal(err)
	}
	body := `---
description: broken skill
execution_mode: nope
supported_agents:
  - claude
  - invalid-agent
allowed_tools:
  - ""
command: echo hi
runtime: nope
---

# Broken Skill

## What it does

Broken on purpose.
`
	if err := os.WriteFile(filepath.Join(root, "skills", "broken", "SKILL.md"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	validateCmd := exec.Command(bin, "skills", "validate", root)
	out, err := validateCmd.CombinedOutput()
	if err == nil {
		t.Fatal("expected validation to fail")
	}
	text := string(out)
	for _, want := range []string{
		"Skill validation found",
		filepath.Join("skills", "broken", "SKILL.md") + ": missing frontmatter field: name",
		"invalid execution_mode",
		"unsupported agent",
		"allowed_tools cannot contain empty values",
		"missing section: When to use",
		"missing section: How to run",
		"missing section: Constraints",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("validation output missing %q:\n%s", want, text)
		}
	}
}

func TestHookplexSkillsExamplesValidateAndRender(t *testing.T) {
	bin := buildHookplex(t)
	root := RepoRoot(t)
	examples := []string{
		filepath.Join(root, "examples", "skills", "go-command-lint"),
		filepath.Join(root, "examples", "skills", "cli-wrapper-formatter"),
		filepath.Join(root, "examples", "skills", "docs-only-review"),
	}
	for _, example := range examples {
		t.Run(filepath.Base(example), func(t *testing.T) {
			validateCmd := exec.Command(bin, "skills", "validate", example)
			if out, err := validateCmd.CombinedOutput(); err != nil {
				t.Fatalf("hookplex skills validate: %v\n%s", err, out)
			}
			renderCmd := exec.Command(bin, "skills", "render", example, "--target", "all")
			if out, err := renderCmd.CombinedOutput(); err != nil {
				t.Fatalf("hookplex skills render: %v\n%s", err, out)
			}
		})
	}
}

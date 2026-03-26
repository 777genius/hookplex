package app

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/hookplex/hookplex/cli/internal/skills/adapters/filesystem"
	"github.com/hookplex/hookplex/cli/internal/skills/adapters/render"
	"github.com/hookplex/hookplex/cli/internal/skills/domain"
)

type InitOptions struct {
	Name        string
	Description string
	Template    filesystem.InitTemplate
	OutputDir   string
	Command     string
	Force       bool
}

type ValidateOptions struct {
	Root string
}

type RenderOptions struct {
	Root   string
	Target string
}

type ValidationFailure struct {
	Path    string
	Message string
}

type ValidationReport struct {
	Skills   []string
	Failures []ValidationFailure
}

type Service struct {
	Repo filesystem.Repository
}

func (s Service) Init(opts InitOptions) (string, error) {
	name := strings.TrimSpace(opts.Name)
	if err := validateName(name); err != nil {
		return "", err
	}
	root := strings.TrimSpace(opts.OutputDir)
	if root == "" {
		root = "."
	}
	desc := strings.TrimSpace(opts.Description)
	if desc == "" {
		desc = "hookplex skill"
	}
	command := strings.TrimSpace(opts.Command)
	if command == "" {
		command = "replace-me"
	}
	if err := s.Repo.InitSkill(root, filesystem.TemplateData{
		SkillName:   name,
		Description: desc,
		Command:     command,
		CommandLine: command,
	}, opts.Template, opts.Force); err != nil {
		return "", err
	}
	return filepath.Join(root, "skills", name), nil
}

func (s Service) Validate(opts ValidateOptions) (ValidationReport, error) {
	names, err := s.Repo.Discover(opts.Root)
	if err != nil {
		return ValidationReport{}, err
	}
	report := ValidationReport{Skills: names}
	for _, name := range names {
		doc, err := s.Repo.LoadSkill(opts.Root, name)
		if err != nil {
			report.Failures = append(report.Failures, ValidationFailure{Path: filepath.Join("skills", name, "SKILL.md"), Message: err.Error()})
			continue
		}
		report.Failures = append(report.Failures, validateDoc(name, doc)...)
	}
	return report, nil
}

func (s Service) Render(opts RenderOptions) ([]domain.Artifact, error) {
	names, err := s.Repo.Discover(opts.Root)
	if err != nil {
		return nil, err
	}
	var renderers []renderer
	switch strings.ToLower(strings.TrimSpace(opts.Target)) {
	case "", "all":
		renderers = []renderer{render.ClaudeRenderer{}, render.CodexRenderer{}}
	case "claude":
		renderers = []renderer{render.ClaudeRenderer{}}
	case "codex":
		renderers = []renderer{render.CodexRenderer{}}
	default:
		return nil, fmt.Errorf("unknown render target %q", opts.Target)
	}
	docs := make(map[string]domain.SkillDocument, len(names))
	var failures []ValidationFailure
	for _, name := range names {
		doc, err := s.Repo.LoadSkill(opts.Root, name)
		if err != nil {
			failures = append(failures, ValidationFailure{Path: filepath.Join("skills", name, "SKILL.md"), Message: err.Error()})
			continue
		}
		docs[name] = doc
		failures = append(failures, validateDoc(name, doc)...)
	}
	if len(failures) > 0 {
		return nil, formatValidationError("cannot render invalid skills", failures)
	}
	var out []domain.Artifact
	for _, name := range names {
		doc := docs[name]
		supportedRenderers := renderersForSkill(doc.Spec, renderers)
		if len(supportedRenderers) == 0 {
			continue
		}
		for _, r := range supportedRenderers {
			artifacts, err := r.Render(name, doc)
			if err != nil {
				return nil, err
			}
			out = append(out, artifacts...)
		}
		if doc.Spec.ExecutionMode == domain.ExecutionCommand {
			cmdBody, err := filesystem.RenderTemplate("command.md.tmpl", filesystem.TemplateData{
				SkillName:            name,
				Description:          doc.Spec.Description,
				CommandLine:          filesystem.CommandLine(doc.Spec),
				Runtime:              string(doc.Spec.Runtime),
				AllowedTools:         doc.Spec.AllowedTools,
				CompatibilitySummary: compatibilitySummary(doc.Spec.Compatibility),
				ExecutionNotes:       executionNotes(doc.Spec),
			})
			if err != nil {
				return nil, err
			}
			out = append(out, domain.Artifact{
				RelPath: filepath.Join("commands", name+".md"),
				Content: cmdBody,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].RelPath < out[j].RelPath })
	return out, nil
}

func (s Service) WriteArtifacts(root string, artifacts []domain.Artifact) error {
	return s.Repo.WriteArtifacts(root, artifacts)
}

type renderer interface {
	Render(name string, doc domain.SkillDocument) ([]domain.Artifact, error)
	Target() string
}

func validateName(name string) error {
	if name == "" {
		return fmt.Errorf("skill name is empty")
	}
	for _, r := range name {
		if !(r == '-' || r == '_' || r >= '0' && r <= '9' || r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') {
			return fmt.Errorf("invalid skill name %q", name)
		}
	}
	return nil
}

func validateDoc(name string, doc domain.SkillDocument) []ValidationFailure {
	var failures []ValidationFailure
	skillPath := filepath.Join("skills", name, "SKILL.md")
	if strings.TrimSpace(doc.Spec.Name) == "" {
		failures = append(failures, ValidationFailure{Path: skillPath, Message: "missing frontmatter field: name"})
	} else if strings.TrimSpace(doc.Spec.Name) != name {
		failures = append(failures, ValidationFailure{Path: skillPath, Message: fmt.Sprintf("frontmatter name %q must match skill directory %q", doc.Spec.Name, name)})
	} else if err := validateName(strings.TrimSpace(doc.Spec.Name)); err != nil {
		failures = append(failures, ValidationFailure{Path: skillPath, Message: err.Error()})
	}
	if strings.TrimSpace(doc.Spec.Description) == "" {
		failures = append(failures, ValidationFailure{Path: skillPath, Message: "missing frontmatter field: description"})
	}
	switch doc.Spec.ExecutionMode {
	case domain.ExecutionDocsOnly, domain.ExecutionCommand:
	default:
		failures = append(failures, ValidationFailure{Path: skillPath, Message: fmt.Sprintf("invalid execution_mode %q (expected %q or %q)", doc.Spec.ExecutionMode, domain.ExecutionDocsOnly, domain.ExecutionCommand)})
	}
	if len(doc.Spec.SupportedAgents) == 0 {
		failures = append(failures, ValidationFailure{Path: skillPath, Message: "missing frontmatter field: supported_agents"})
	}
	for _, tool := range doc.Spec.AllowedTools {
		if strings.TrimSpace(tool) == "" {
			failures = append(failures, ValidationFailure{Path: skillPath, Message: "allowed_tools cannot contain empty values"})
		}
	}
	for _, input := range doc.Spec.Inputs {
		if strings.TrimSpace(input) == "" {
			failures = append(failures, ValidationFailure{Path: skillPath, Message: "inputs cannot contain empty values"})
		}
	}
	for _, output := range doc.Spec.Outputs {
		if strings.TrimSpace(output) == "" {
			failures = append(failures, ValidationFailure{Path: skillPath, Message: "outputs cannot contain empty values"})
		}
	}
	for _, require := range doc.Spec.Compatibility.Requires {
		if strings.TrimSpace(require) == "" {
			failures = append(failures, ValidationFailure{Path: skillPath, Message: "compatibility.requires cannot contain empty values"})
		}
	}
	for _, osName := range doc.Spec.Compatibility.SupportedOS {
		if strings.TrimSpace(osName) == "" {
			failures = append(failures, ValidationFailure{Path: skillPath, Message: "compatibility.supported_os cannot contain empty values"})
		}
	}
	for _, note := range doc.Spec.Compatibility.Notes {
		if strings.TrimSpace(note) == "" {
			failures = append(failures, ValidationFailure{Path: skillPath, Message: "compatibility.notes cannot contain empty values"})
		}
	}
	for key, hint := range doc.Spec.AgentHints {
		switch domain.Agent(key) {
		case domain.AgentClaude, domain.AgentCodex:
		default:
			failures = append(failures, ValidationFailure{Path: skillPath, Message: fmt.Sprintf("unsupported agent_hints key %q", key)})
			continue
		}
		if !containsAgent(doc.Spec.SupportedAgents, domain.Agent(key)) {
			failures = append(failures, ValidationFailure{Path: skillPath, Message: fmt.Sprintf("agent_hints.%s requires %q in supported_agents", key, key)})
		}
		for _, note := range hint.Notes {
			if strings.TrimSpace(note) == "" {
				failures = append(failures, ValidationFailure{Path: skillPath, Message: fmt.Sprintf("agent_hints.%s.notes cannot contain empty values", key)})
			}
		}
	}
	for _, agent := range doc.Spec.SupportedAgents {
		switch agent {
		case domain.AgentClaude, domain.AgentCodex:
		default:
			failures = append(failures, ValidationFailure{Path: skillPath, Message: fmt.Sprintf("unsupported agent %q (supported: %q, %q)", agent, domain.AgentClaude, domain.AgentCodex)})
		}
	}
	if doc.Spec.ExecutionMode == domain.ExecutionCommand {
		if strings.TrimSpace(doc.Spec.Command) == "" {
			failures = append(failures, ValidationFailure{Path: skillPath, Message: "execution_mode=command requires command"})
		}
		if wd := strings.TrimSpace(doc.Spec.WorkingDir); wd != "" {
			if filepath.IsAbs(wd) {
				failures = append(failures, ValidationFailure{Path: skillPath, Message: "working_dir must be relative to the skill root"})
			}
		}
		if timeout := strings.TrimSpace(doc.Spec.Timeout); timeout != "" {
			if _, err := time.ParseDuration(timeout); err != nil {
				failures = append(failures, ValidationFailure{Path: skillPath, Message: fmt.Sprintf("timeout must be a valid duration: %v", err)})
			}
		}
		switch doc.Spec.Runtime {
		case domain.RuntimeGo, domain.RuntimeShell, domain.RuntimePython, domain.RuntimeNode, domain.RuntimeDeno, domain.RuntimeExternal, domain.RuntimeGeneric:
		default:
			failures = append(failures, ValidationFailure{Path: skillPath, Message: fmt.Sprintf("execution_mode=command requires valid runtime (got %q)", doc.Spec.Runtime)})
		}
	}
	requiredSections := []string{"## What it does", "## When to use", "## How to run", "## Constraints"}
	for _, section := range requiredSections {
		if !strings.Contains(doc.Body, section) {
			failures = append(failures, ValidationFailure{Path: skillPath, Message: "missing section: " + strings.TrimPrefix(section, "## ")})
		}
	}
	return failures
}

func compatibilitySummary(spec domain.CompatibilitySpec) []string {
	var out []string
	if len(spec.Requires) > 0 {
		out = append(out, "Requires: "+strings.Join(spec.Requires, ", "))
	}
	if len(spec.SupportedOS) > 0 {
		out = append(out, "Supported OS: "+strings.Join(spec.SupportedOS, ", "))
	}
	if spec.RepoRequired {
		out = append(out, "Requires a repository checkout")
	}
	if spec.NetworkRequired {
		out = append(out, "May require network access")
	}
	out = append(out, spec.Notes...)
	return out
}

func executionNotes(spec domain.SkillSpec) []string {
	var out []string
	if spec.SafeToRetry != nil {
		out = append(out, "Safe to retry: "+yesNo(*spec.SafeToRetry))
	}
	if spec.WritesFiles != nil {
		out = append(out, "Writes files: "+yesNo(*spec.WritesFiles))
	}
	if spec.ProducesJSON != nil {
		out = append(out, "Produces JSON: "+yesNo(*spec.ProducesJSON))
	}
	return out
}

func yesNo(v bool) string {
	if v {
		return "yes"
	}
	return "no"
}

func renderersForSkill(spec domain.SkillSpec, candidates []renderer) []renderer {
	allowed := make(map[string]struct{}, len(spec.SupportedAgents))
	for _, agent := range spec.SupportedAgents {
		allowed[string(agent)] = struct{}{}
	}
	out := make([]renderer, 0, len(candidates))
	for _, candidate := range candidates {
		if _, ok := allowed[candidate.Target()]; ok {
			out = append(out, candidate)
		}
	}
	return out
}

func containsAgent(agents []domain.Agent, want domain.Agent) bool {
	for _, agent := range agents {
		if agent == want {
			return true
		}
	}
	return false
}

func formatValidationError(prefix string, failures []ValidationFailure) error {
	var b strings.Builder
	b.WriteString(prefix)
	b.WriteString(":\n")
	for _, failure := range failures {
		b.WriteString("- ")
		b.WriteString(failure.Path)
		b.WriteString(": ")
		b.WriteString(failure.Message)
		b.WriteString("\n")
	}
	return fmt.Errorf(strings.TrimRight(b.String(), "\n"))
}

// Package scaffold writes the current-state hookplex init project tree.
package scaffold

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

//go:embed templates/*.tmpl
var tmplFS embed.FS

// Data is passed to all templates.
type Data struct {
	ModulePath  string
	ProjectName string
	Description string
	Platform    string
	WithExtras  bool
}

type TemplateFile struct {
	Path     string
	Template string
	Extra    bool
}

type PlatformDefinition struct {
	Name  string
	Files []TemplateFile
}

type PlannedFile struct {
	RelPath  string
	Template string
}

type ProjectPlan struct {
	Platform string
	Files    []PlannedFile
	Data     Data
}

var nameRe = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{0,63}$`)

// ValidateProjectName returns an error if name is not a safe directory/binary segment.
func ValidateProjectName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("project name is empty")
	}
	if !nameRe.MatchString(name) {
		return fmt.Errorf("invalid project name %q: use letters, digits, underscore, hyphen; start with a letter; max 64 characters", name)
	}
	return nil
}

// DefaultModulePath returns example.com/<name> for generated go.mod.
func DefaultModulePath(name string) string {
	return "example.com/" + name
}

// Paths lists relative paths created by Write (for tests and docs).
func Paths(platform, name string, extras bool) []string {
	def, ok := LookupPlatform(platform)
	if !ok {
		return nil
	}
	return planPaths(expandTemplateFiles(def.Files, Data{ProjectName: name, Platform: def.Name, WithExtras: extras}))
}

func BuildPlan(d Data) (ProjectPlan, error) {
	if err := ValidateProjectName(d.ProjectName); err != nil {
		return ProjectPlan{}, err
	}
	if strings.TrimSpace(d.ModulePath) == "" {
		d.ModulePath = DefaultModulePath(d.ProjectName)
	}
	if strings.TrimSpace(d.Description) == "" {
		d.Description = "hookplex plugin"
	}
	p, ok := LookupPlatform(d.Platform)
	if !ok {
		return ProjectPlan{}, fmt.Errorf("unknown platform %q", d.Platform)
	}
	d.Platform = p.Name
	out := ProjectPlan{
		Platform: p.Name,
		Data:     d,
		Files:    expandTemplateFiles(p.Files, d),
	}
	return out, nil
}

// Write creates the plugin tree at root (must exist or be creatable).
func Write(root string, d Data, force bool) error {
	plan, err := BuildPlan(d)
	if err != nil {
		return err
	}
	return Apply(root, plan, force)
}

func Apply(root string, plan ProjectPlan, force bool) error {
	info, err := os.Stat(root)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if mkErr := os.MkdirAll(root, 0o755); mkErr != nil {
			return mkErr
		}
	} else if !info.IsDir() {
		return fmt.Errorf("output path %q is not a directory", root)
	} else {
		entries, rerr := os.ReadDir(root)
		if rerr != nil {
			return rerr
		}
		if len(entries) > 0 && !force {
			return fmt.Errorf("directory %q is not empty (use --force to overwrite files)", root)
		}
	}

	for _, file := range plan.Files {
		if err := writeOne(root, file.RelPath, file.Template, plan.Data, force); err != nil {
			return err
		}
	}
	return nil
}

func planPaths(tasks []PlannedFile) []string {
	out := make([]string, 0, len(tasks))
	for _, task := range tasks {
		out = append(out, task.RelPath)
	}
	return out
}

func expandTemplateFiles(files []TemplateFile, d Data) []PlannedFile {
	out := make([]PlannedFile, 0, len(files))
	for _, file := range files {
		if file.Extra && !d.WithExtras {
			continue
		}
		out = append(out, PlannedFile{
			RelPath:  expandPathTemplate(file.Path, d),
			Template: file.Template,
		})
	}
	return out
}

func expandPathTemplate(path string, d Data) string {
	path = strings.ReplaceAll(path, "{{.ProjectName}}", d.ProjectName)
	path = strings.ReplaceAll(path, "{{.Platform}}", d.Platform)
	return path
}

func writeOne(root, rel, tplName string, d Data, force bool) error {
	raw, err := tmplFS.ReadFile(filepath.Join("templates", tplName))
	if err != nil {
		return err
	}
	t, err := template.New(tplName).Parse(string(raw))
	if err != nil {
		return fmt.Errorf("parse %s: %w", tplName, err)
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, d); err != nil {
		return fmt.Errorf("execute %s: %w", tplName, err)
	}
	full := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		return err
	}
	if _, err := os.Stat(full); err == nil && !force {
		return fmt.Errorf("refusing to overwrite existing file %s (use --force)", rel)
	}
	mode := fs.FileMode(0o644)
	if strings.HasSuffix(rel, ".sh") {
		mode = 0o755
	}
	return os.WriteFile(full, buf.Bytes(), mode)
}

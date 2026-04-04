package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/777genius/plugin-kit-ai/cli/internal/pluginmanifest"
	"github.com/777genius/plugin-kit-ai/cli/internal/publicationmodel"
)

func TestPublicationTextShowsPackagesAndChannels(t *testing.T) {
	t.Parallel()
	cmd := newPublicationCmd(fakeInspectRunner{
		report: pluginmanifest.Inspection{
			Publication: publicationmodel.Model{
				Core: publicationmodel.Core{
					APIVersion:  "v1",
					Name:        "demo",
					Version:     "0.1.0",
					Description: "demo plugin",
				},
				Packages: []publicationmodel.Package{
					{
						Target:          "codex-package",
						PackageFamily:   "codex-plugin",
						ChannelFamilies: []string{"codex-marketplace"},
						AuthoredInputs:  []string{"plugin.yaml", "publish/codex/marketplace.yaml"},
						ManagedArtifacts: []string{
							".codex-plugin/plugin.json",
							".agents/plugins/marketplace.json",
						},
					},
				},
				Channels: []publicationmodel.Channel{
					{
						Family:         "codex-marketplace",
						Path:           "publish/codex/marketplace.yaml",
						PackageTargets: []string{"codex-package"},
						Details: map[string]string{
							"marketplace_name":      "local-repo",
							"source_root":           "./",
							"category":              "Productivity",
							"installation_policy":   "AVAILABLE",
							"authentication_policy": "ON_INSTALL",
						},
					},
				},
			},
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
		"publication demo 0.1.0 api_version=v1",
		"packages: 1 channels: 1",
		"package[codex-package]: family=codex-plugin channels=codex-marketplace inputs=2 managed=2",
		"channel[codex-marketplace]: path=publish/codex/marketplace.yaml targets=codex-package",
		"details=authentication_policy=ON_INSTALL,category=Productivity,installation_policy=AVAILABLE,marketplace_name=local-repo,source_root=./",
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("publication output missing %q:\n%s", want, output)
		}
	}
}

func TestPublicationJSONEmitsPublicationModelOnly(t *testing.T) {
	t.Parallel()
	cmd := newPublicationCmd(fakeInspectRunner{
		report: pluginmanifest.Inspection{
			Publication: publicationmodel.Model{
				Core: publicationmodel.Core{
					APIVersion:  "v1",
					Name:        "gemini-publish",
					Version:     "0.1.0",
					Description: "gemini publish",
				},
				Packages: []publicationmodel.Package{
					{
						Target:          "gemini",
						PackageFamily:   "gemini-extension",
						ChannelFamilies: []string{"gemini-gallery"},
					},
				},
				Channels: []publicationmodel.Channel{
					{
						Family:         "gemini-gallery",
						Path:           "publish/gemini/gallery.yaml",
						PackageTargets: []string{"gemini"},
						Details: map[string]string{
							"distribution":          "github_release",
							"repository_visibility": "public",
							"github_topic":          "gemini-cli-extension",
							"manifest_root":         "release_archive_root",
						},
					},
				},
			},
		},
	})
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"--format", "json", "."})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	var payload map[string]any
	if err := json.Unmarshal(buf.Bytes(), &payload); err != nil {
		t.Fatalf("json parse: %v\n%s", err, buf.Bytes())
	}
	if payload["core"] == nil || payload["packages"] == nil || payload["channels"] == nil {
		t.Fatalf("publication payload = %+v", payload)
	}
	if _, found := payload["manifest"]; found {
		t.Fatalf("publication json unexpectedly includes full inspect payload: %+v", payload)
	}
}

func TestPublicationHelpMentionsSupportedTargets(t *testing.T) {
	t.Parallel()
	cmd := newPublicationCmd(fakeInspectRunner{})
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"--help"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	output := buf.String()
	for _, want := range []string{
		`publication target ("all", "claude", "codex-package", or "gemini")`,
		`output format: text or json`,
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("help output missing %q:\n%s", want, output)
		}
	}
}

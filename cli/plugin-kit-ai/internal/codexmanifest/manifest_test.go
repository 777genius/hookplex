package codexmanifest

import (
	"strings"
	"testing"
)

func TestParseInterfaceDocRejectsInvalidDefaultPromptShape(t *testing.T) {
	t.Parallel()
	if _, err := ParseInterfaceDoc([]byte(`{"defaultPrompt":"Run the demo"}`)); err == nil || !strings.Contains(err.Error(), "interface.defaultPrompt must be an array of strings") {
		t.Fatalf("ParseInterfaceDoc error = %v", err)
	}
	if _, err := ParseInterfaceDoc([]byte(`{"defaultPrompt":[""]}`)); err == nil || !strings.Contains(err.Error(), "interface.defaultPrompt[0] must not be empty") {
		t.Fatalf("ParseInterfaceDoc error = %v", err)
	}
}

func TestParseInterfaceDocRejectsNonObjectJSON(t *testing.T) {
	t.Parallel()
	if _, err := ParseInterfaceDoc([]byte(`["demo"]`)); err == nil || !strings.Contains(err.Error(), "Codex interface doc must be a JSON object") {
		t.Fatalf("ParseInterfaceDoc error = %v", err)
	}
}

func TestDecodeImportedPluginManifestRejectsInvalidInterfaceDefaultPrompt(t *testing.T) {
	t.Parallel()
	if _, err := DecodeImportedPluginManifest([]byte(`{"name":"demo","version":"0.1.0","description":"demo","interface":{"defaultPrompt":"Run the demo"}}`)); err == nil || !strings.Contains(err.Error(), "interface.defaultPrompt must be an array of strings") {
		t.Fatalf("DecodeImportedPluginManifest error = %v", err)
	}
}

func TestParseAppManifestDocRejectsNonObjectJSON(t *testing.T) {
	t.Parallel()
	if _, err := ParseAppManifestDoc([]byte(`["demo"]`)); err == nil || !strings.Contains(err.Error(), "Codex app manifest must be a JSON object") {
		t.Fatalf("ParseAppManifestDoc error = %v", err)
	}
}

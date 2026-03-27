package targetcontracts

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestMarkdownStaysInSyncWithGeneratedDoc(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", "..", "..", ".."))
	body, err := os.ReadFile(filepath.Join(root, "docs", "generated", "target_support_matrix.md"))
	if err != nil {
		t.Fatal(err)
	}
	got := string(body)
	want := string(Markdown(All()))
	if got != want {
		t.Fatalf("target support matrix drifted\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}

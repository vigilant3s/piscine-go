package textfix

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGolden(t *testing.T) {
	// testdata is a sibling of textfix: go_reloaded/textfix -> go_reloaded/testdata
	base := filepath.Join("..", "testdata")

	entries, err := os.ReadDir(base)
	if err != nil {
		t.Fatalf("read testdata dir: %v", err)
	}

	for _, e := range entries {
		name := e.Name()
		if !strings.HasSuffix(name, ".in") {
			continue
		}
		inPath := filepath.Join(base, name)
		outPath := filepath.Join(base, strings.TrimSuffix(name, ".in")+".out")

		inBytes, err := os.ReadFile(inPath)
		if err != nil {
			t.Fatalf("read input %s: %v", inPath, err)
		}
		wantBytes, err := os.ReadFile(outPath)
		if err != nil {
			t.Fatalf("read expected %s: %v", outPath, err)
		}

		got := ProcessText(string(inBytes))
		want := string(wantBytes)

		if got != want {
			t.Fatalf("%s\n--- WANT ---\n%q\n--- GOT ----\n%q\n", name, want, got)
		}
	}
}

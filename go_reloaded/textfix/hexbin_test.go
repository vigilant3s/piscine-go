package textfix

import "testing"

func TestHexBinSimple(t *testing.T) {
	in := "Add 1E (hex) and 10 (bin)."
	got := ProcessText(in)
	want := "Add 30 and 2."
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

package version

import "testing"

func TestString(t *testing.T) {
	if got := String(); got != "tally dev (none)" {
		t.Fatalf("got %q", got)
	}
	old := Version
	defer func() { Version = old }()
	Version = "1.2.3"
	if got := String(); got != "tally 1.2.3 (none)" {
		t.Fatalf("got %q", got)
	}
}

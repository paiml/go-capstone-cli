package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain_(t *testing.T) {
	got := -1
	exit = func(c int) { got = c }
	old := os.Args
	defer func() { os.Args = old }()
	os.Args = []string{"tally", "sum", "2", "3"}
	main()
	if got != 0 {
		t.Fatalf("exit %d", got)
	}
}

// TestBuiltBinary is the integration test: build the real binary with a
// stamped version and run it as a user would.
func TestBuiltBinary(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "tally")
	build := exec.Command("go", "build", "-ldflags", "-X github.com/paiml/go-capstone-cli/version.Version=9.9.9", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}
	out, err := exec.Command(bin, "version").Output()
	if err != nil || !strings.HasPrefix(string(out), "tally 9.9.9") {
		t.Fatalf("version = %q, %v", out, err)
	}
	if err := exec.Command(bin, "sum", "x").Run(); err == nil {
		t.Fatal("sum x should exit non-zero")
	} else if ee, ok := err.(*exec.ExitError); !ok || ee.ExitCode() != 1 {
		t.Fatalf("exit = %v", err)
	}
}

package logx

import (
	"bytes"
	"strings"
	"testing"
)

func TestLevels(t *testing.T) {
	var quiet, loud bytes.Buffer
	New(&quiet, false).Debug("hidden", "k", 1)
	New(&loud, true).Debug("shown", "k", 1)
	if quiet.Len() != 0 {
		t.Fatalf("quiet logged %q", quiet.String())
	}
	if got := strings.TrimSpace(loud.String()); got != "level=DEBUG msg=shown k=1" {
		t.Fatalf("got %q", got)
	}
}

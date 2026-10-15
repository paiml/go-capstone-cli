package cli

import (
	"bytes"
	"strings"
	"testing"
)

func run(in string, env string, args ...string) (int, string, string) {
	var o, e bytes.Buffer
	code := Run(Env{In: strings.NewReader(in), Out: &o, Err: &e, Getenv: func(k string) string {
		if k == "TALLY_UNIT" {
			return env
		}
		return ""
	}}, args)
	return code, o.String(), e.String()
}

func TestRun(t *testing.T) {
	cases := []struct {
		name       string
		in, env    string
		args       []string
		code       int
		out, errIn string
	}{
		{"sum", "", "", []string{"sum", "1", "2", "39"}, 0, "42\n", ""},
		{"sum unit flag", "", "", []string{"sum", "1", "--unit", "kg"}, 0, "1 kg\n", ""},
		{"sum unit env", "", "ms", []string{"sum", "5"}, 0, "5 ms\n", ""},
		{"sum verbose", "", "", []string{"-v", "sum", "7"}, 0, "7\n", "level=DEBUG msg=add n=7"},
		{"sum bad", "", "", []string{"sum", "x"}, 1, "", `"x" is not a whole number`},
		{"sum none", "", "", []string{"sum"}, 2, "", "requires at least 1 arg"},
		{"count", "a b\nc\n", "", []string{"count"}, 0, "2 lines 3 words\n", ""},
		{"version", "", "", []string{"version"}, 0, "tally dev (none)\n", ""},
		{"unknown", "", "", []string{"nope"}, 2, "", "unknown command"},
	}
	for _, c := range cases {
		code, out, errOut := run(c.in, c.env, c.args...)
		if code != c.code || out != c.out || !strings.Contains(errOut, c.errIn) {
			t.Errorf("%s: code=%d out=%q err=%q", c.name, code, out, errOut)
		}
	}
}

func TestOS(t *testing.T) {
	e := OS()
	if e.In == nil || e.Out == nil || e.Err == nil || e.Getenv == nil {
		t.Fatal("incomplete")
	}
}

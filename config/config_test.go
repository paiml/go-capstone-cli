package config

import "testing"

func TestParseFile(t *testing.T) {
	got := ParseFile("# c\n\ngreeting = hi\nbadline\nname=Ada\n")
	if len(got) != 2 || got["greeting"] != "hi" || got["name"] != "Ada" {
		t.Fatalf("got %v", got)
	}
}

func TestResolvePrecedence(t *testing.T) {
	file := map[string]string{"k": "f"}
	cases := []struct {
		name, flag, env string
		file            map[string]string
		want            string
		src             Source
	}{
		{"flag wins", "a", "b", file, "a", Flag},
		{"env next", "", "b", file, "b", Env},
		{"file next", "", "", file, "f", File},
		{"default last", "", "", nil, "d", Default},
	}
	for _, c := range cases {
		v, s := Resolve("k", c.flag, c.env, c.file, "d")
		if v != c.want || s != c.src {
			t.Errorf("%s: got %q %q", c.name, v, s)
		}
	}
}

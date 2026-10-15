// Package config resolves one setting from three sources: flag, then
// environment, then file, then a default.
package config

import (
	"bufio"
	"strings"
)

// Source says where a resolved value came from.
type Source string

// The sources, in precedence order.
const (
	Flag    Source = "flag"
	Env     Source = "env"
	File    Source = "file"
	Default Source = "default"
)

// ParseFile reads simple key=value lines. Blank lines and # comments are skipped.
func ParseFile(text string) map[string]string {
	out := map[string]string{}
	sc := bufio.NewScanner(strings.NewReader(text))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		k, v, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		out[strings.TrimSpace(k)] = strings.TrimSpace(v)
	}
	return out
}

// Resolve returns the value for key and where it came from.
// flagVal is "" when the flag was not set.
func Resolve(key, flagVal, envVal string, file map[string]string, def string) (string, Source) {
	if flagVal != "" {
		return flagVal, Flag
	}
	if envVal != "" {
		return envVal, Env
	}
	if v, ok := file[key]; ok {
		return v, File
	}
	return def, Default
}

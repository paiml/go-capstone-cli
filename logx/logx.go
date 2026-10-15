// Package logx builds the structured logger: text to stderr, level by flag.
package logx

import (
	"io"
	"log/slog"
)

// New returns a logger writing key=value lines without timestamps, so output
// is stable enough to assert on.
func New(w io.Writer, verbose bool) *slog.Logger {
	level := slog.LevelWarn
	if verbose {
		level = slog.LevelDebug
	}
	h := slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})
	return slog.New(h)
}

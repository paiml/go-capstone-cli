// Package version holds build stamps. The linker overwrites them:
//
//	go build -ldflags "-X github.com/paiml/go-capstone-cli/version.Version=1.2.3"
package version

import "fmt"

// Version and Commit are set with -ldflags -X at build time.
var (
	Version = "dev"
	Commit  = "none"
)

// String is the one-line version.
func String() string { return fmt.Sprintf("tally %s (%s)", Version, Commit) }

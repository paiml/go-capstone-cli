// Command tally adds numbers and counts lines.
package main

import (
	"os"

	"github.com/paiml/go-capstone-cli/cli"
)

var exit = os.Exit

func main() { exit(cli.Run(cli.OS(), os.Args[1:])) }

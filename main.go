// Command stamp-cli stands in for release-cli in architecture spike C. It exists
// only to report its linker-injected stamps so the composite action's version and
// protocol guard can be exercised against a real GitHub release.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// version is the semantic version injected at build time.
var version = "dev"

// commit is the source commit injected at build time.
var commit = "unknown"

// protocolLiteral is the CLI/workflow protocol integer. It is a source literal so
// a CI check can prove it equals the action's expected value; only a deliberate
// release-unit change moves it.
const protocolLiteral = 1

// stamps is the machine-readable identity of this binary.
type stamps struct {
	// Version is the semantic version.
	Version string `json:"version"`
	// Commit is the source commit.
	Commit string `json:"commit"`
	// Protocol is the workflow/binary contract integer.
	Protocol int `json:"protocol"`
}

// main prints the stamps as one JSON document.
func main() {
	if len(os.Args) > 1 && os.Args[1] != "version" {
		fmt.Fprintf(os.Stderr, "usage: %s version [--json]\n", os.Args[0])
		os.Exit(2)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(stamps{Version: version, Commit: commit, Protocol: protocolLiteral}); err != nil {
		fmt.Fprintf(os.Stderr, "encode stamps: %v\n", err)
		os.Exit(1)
	}
}

// protocolString renders the protocol for shell comparison.
func protocolString() string { return strconv.Itoa(protocolLiteral) }

// contract reports the stamped contract as a single line for the guard's
// human-readable output and for release-note evidence.
func contract() string {
	return fmt.Sprintf("version=%s protocol=%s commit=%s", version, protocolString(), commit)
}

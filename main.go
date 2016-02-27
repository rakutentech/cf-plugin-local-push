package main

import (
	"os"

	"github.com/cloudfoundry/cli/plugin"
)

func main() {
	localPush := LocalPush{
		OutStream: os.Stdout,
	}

	// Start plugin
	plugin.Start(&localPush)
}

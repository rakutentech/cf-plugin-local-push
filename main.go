package main

import (
	"os"

	"github.com/cloudfoundry/cli/plugin"
)

func main() {
	localPush := LocalPush{
		OutStream: os.Stdout,
		InStream:  os.Stdin,
	}

	// Start plugin
	plugin.Start(&localPush)
}

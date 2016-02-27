package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

const Name string = "local-push"

var Version = plugin.VersionType{
	Major: 0,
	Minor: 1,
	Build: 0,
}

func VersionStr() string {
	return fmt.Sprintf("%d.%d.%d", Version.Major, Version.Minor, Version.Build)
}

package main

import (
	"os"
	"os/exec"
)

// docker runs docker command
func docker(args ...string) error {
	cmd := exec.Command("docker", args...)
	cmd.Stderr = os.Stdout // ioutil.Discard
	cmd.Stdout = os.Stdout // ioutil.Discard

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

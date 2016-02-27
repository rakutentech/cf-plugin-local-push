package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/tcnksm/go-input"
)

// Exit codes are int values that represent an exit code
// for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// EnvDebug is environmental variable for enabling debug
// output
const EnvDebug = "DEBUG_PLUGIN"

const (
	// Dockerfile is file name of Dockerfile
	Dockerfile = "Dockerfile"
)

// dockerfileText is used for build docker image for target application.
var dockerfileText = `FROM tcnksm/cf-buildstep:latest
ENV USER vcap
ADD . /app
RUN /build/builder
CMD /start web`

// Debugf prints debug output when EnvDebug is given
func Debugf(format string, args ...interface{}) {
	if env := os.Getenv(EnvDebug); len(env) != 0 {
		fmt.Fprintf(os.Stdout, "[DEBUG] "+format+"\n", args...)
	}
}

// LocalPush
type LocalPush struct {
	OutStream io.Writer
	InStream  io.Reader
}

// Run starts plugin process.
func (p *LocalPush) Run(cliConn plugin.CliConnection, arg []string) {
	Debugf("Run local-push plugin")
	Debugf("Arg: %#v", arg)

	// Ensure local-push is called.
	// Plugin is also called when install/uninstall via cf command.
	// Ignore such other calls.
	if len(arg) < 1 || arg[0] != Name {
		os.Exit(ExitCodeOK)
	}

	// Read CLI context.
	ctx, err := NewCLIContext(cliConn)
	if err != nil {
		fmt.Fprintf(p.OutStream, "Failed to read cf command context: %s\n", err)
		os.Exit(ExitCodeError)
	}

	// Call run instead of doing the work here so we can use
	// `defer` statements within the function and have them work properly.
	// (defers aren't called with os.Exit)
	os.Exit(p.run(ctx, arg))
}

// run runs local-push.
func (p *LocalPush) run(ctx *CLIContext, args []string) int {

	ui := &input.UI{
		Writer: p.OutStream,
		Reader: p.InStream,
	}

	// Check docker is installed or not.
	if _, err := exec.LookPath("docker"); err != nil {
		fmt.Fprintf(p.OutStream, "docker command is not found in your $PATH. Install it before.\n")
		return ExitCodeError
	}

	// Check Dockerfile is exist or not
	if _, err := os.Stat(Dockerfile); !os.IsNotExist(err) {
		fmt.Fprintf(p.OutStream, "Dockerfile is already exist\n")
		query := "Overwrite Dockerfile? [yN]"
		ans, err := ui.Ask(query, &input.Options{
			Default:     "N",
			HideDefault: true,
			HideOrder:   true,
			Required:    true,
			Loop:        true,
			ValidateFunc: func(s string) error {
				if s != "y" && s != "N" {
					return fmt.Errorf("input must be 'y' or 'N'")
				}
				return nil
			},
		})

		if err != nil {
			fmt.Fprintf(p.OutStream, "Failed to read input: %s\n", err)
			return ExitCodeError
		}

		// Stop execution
		if ans != "y" {
			fmt.Fprintf(p.OutStream, "Aborting\n")
			return ExitCodeOK
		}
	}

	fmt.Fprintf(p.OutStream, "(cf-local-push) Generate Dockerfile\n")
	f, err := os.Create("Dockerfile")
	if err != nil {
		fmt.Fprintf(p.OutStream, "%s\n", err)
		return ExitCodeError
	}

	if _, err := f.Write([]byte(dockerfileText)); err != nil {
		fmt.Fprintf(p.OutStream, "%s\n", err)
		return ExitCodeError
	}

	fmt.Fprintf(p.OutStream, "(cf-local-push) Start building docker image\n")
	if err := docker("build", "-t", "image-localpush", "."); err != nil {
		fmt.Fprintf(p.OutStream, "%s\n", err)
		return ExitCodeError
	}

	fmt.Fprintf(p.OutStream, "(cf-local-push) Start running docker container\n")
	if err := docker("run", "-p", "8080:8080", "-e", "PORT=8080", "image-localpush"); err != nil {
		fmt.Fprintf(p.OutStream, "%s\n", err)
		return ExitCodeError
	}

	return 0
}

func (p *LocalPush) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:    Name,
		Version: Version,
		Commands: []plugin.Command{
			{
				Name:     "local-push",
				HelpText: "Push cf app on local Docker container",
				UsageDetails: plugin.Usage{
					Usage: p.Usage(),
				},
			},
		},
	}
}

func (p *LocalPush) Usage() string {
	text := `cf local-push 

local-push command tells cf to deploy your app on local docker
container.
`
	return text
}

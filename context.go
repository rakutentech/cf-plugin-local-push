package main

import "github.com/cloudfoundry/cli/plugin"

// CLIContext is the context which can be retrieved
// from cf command.
type CLIContext struct {
	User     string
	Token    string
	Endpoint string

	// Embeded because some value is needed to
	// be retrieved dynamically.
	plugin.CliConnection
}

// NewCLIContext retrieved current cf command context
func NewCLIContext(cliConn plugin.CliConnection) (*CLIContext, error) {
	user, err := cliConn.Username()
	if err != nil {
		return nil, err
	}

	endpoint, err := cliConn.ApiEndpoint()
	if err != nil {
		return nil, err
	}

	return &CLIContext{
		User:          user,
		Endpoint:      endpoint,
		CliConnection: cliConn,
	}, nil
}

package libanvil

import (
	"context"
	"fmt"
	"io"
	"os/exec"
)

type AnvilRuntime struct {
	ctx          context.Context
	ctxCancel    context.CancelFunc
	cmd          *exec.Cmd
	accsFilename string
}

type RunAnvilOptions struct {
	Executable   string
	StdoutWriter io.Writer
	StderrWriter io.Writer
	CLIArgs      []string
}

var defaultOptions = &RunAnvilOptions{
	Executable:   "anvil",
	StdoutWriter: nil,
	StderrWriter: nil,
	CLIArgs:      nil,
}

func RunAnvil(options *RunAnvilOptions) (*AnvilRuntime, error) {
	if options == nil {
		options = defaultOptions
	}

	var executable string
	if options.Executable != "" {
		executable = options.Executable
	} else {
		executable = defaultOptions.Executable
	}

	var args []string
	if options.CLIArgs != nil {
		args = options.CLIArgs
	} else {
		args = make([]string, 0)
	}

	ctx, ctxCancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, executable, args...)

	w := newIntermediateWriter(options.StdoutWriter)
	cmd.Stdout = w

	if options.StderrWriter != nil {
		cmd.Stderr = options.StderrWriter
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("error running anvil: %w", err)
	}

	return &AnvilRuntime{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		cmd:       cmd,
	}, nil
}

func (a *AnvilRuntime) Stop() {
	a.ctxCancel()
}

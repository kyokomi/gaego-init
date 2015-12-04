package main

import (
	"flag"
	"fmt"
	"io"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		app string
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&app, "app", "", "appEngine application name")
	flags.StringVar(&app, "a", "", "appEngine application name(Short)")

	flVersion := flags.Bool("version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if *flVersion {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if app == "" {
		fmt.Fprintf(cli.errStream, "required pkgName and appName\n")
		return ExitCodeError
	}

	if err := NewBuilder(app).Run(); err != nil {
		fmt.Fprintf(cli.errStream, "run error %s\n", err.Error())
		return ExitCodeError
	}

	return ExitCodeOK
}

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if exitCode := start(os.Stdout); exitCode != 0 {
		os.Exit(exitCode)
	}
}

func start(stdout io.Writer) int {
	options := parseCLIFlags()
	validationErrors := options.validate()
	if len(validationErrors) != 0 {
		err := printValidationErrors(validationErrors, stdout)
		if err != nil {
			fmt.Fprintf(stdout, "failed to print validation errors: %s\n", err.Error())
			return 3
		}
		return 4
	}
	m, err := process(options)
	if err != nil {
		fmt.Fprintf(stdout, "could not process: %s\n", err.Error())
		return 5
	}
	err = printDiffOutput(m, options, stdout)
	if err != nil {
		fmt.Fprintf(stdout, "failed to print output: %s\n", err.Error())
		return 6
	}
	if len(m.Diff) != 0 {
		return 7
	}
	return 0
}

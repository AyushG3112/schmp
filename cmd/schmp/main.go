package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if exitCode := start(os.Stdout, os.Stderr); exitCode != 0 {
		os.Exit(exitCode)
	}
}

// start runs the CLI, writing comparison results to stdout and any diagnostics
// to stderr. It returns the process exit code: 0 on a clean no-diff run, 7 when
// schema differences are found (handy as a CI gate), and 3-6 for the various
// failure modes.
func start(stdout, stderr io.Writer) int {
	options := parseCLIFlags()

	validationErrors := options.validate()
	if len(validationErrors) != 0 {
		if err := printValidationErrors(validationErrors, stderr); err != nil {
			fmt.Fprintf(stderr, "failed to print validation errors: %s\n", err.Error())
			return 3
		}
		return 4
	}

	result, err := process(options)
	if err != nil {
		fmt.Fprintf(stderr, "could not process: %s\n", err.Error())
		return 5
	}

	if err := printDiffOutput(result, options, stdout); err != nil {
		fmt.Fprintf(stderr, "failed to print output: %s\n", err.Error())
		return 6
	}

	if len(result.Diff) != 0 {
		return 7
	}
	return 0
}

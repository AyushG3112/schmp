package main

import (
	"fmt"
	"os"
)

func main() {
	if exitCode := start(); exitCode != 0 {
		os.Exit(exitCode)
	}
}

func start() int {
	options := parseCLIFlags()
	validationErrors := options.validate()
	printValidationErrors(validationErrors)
	if len(validationErrors) != 0 {
		return 3
	}
	m, err := process(options)
	if err != nil {
		fmt.Printf("could not process: %s\n", err.Error())
		return 4
	}
	err = printDiffOutput(m, options)
	if err != nil {
		fmt.Printf("failed to print output: %s\n", err.Error())
		return 5
	}
	if len(m.Diff) != 0 {
		return 6
	}
	return 0
}

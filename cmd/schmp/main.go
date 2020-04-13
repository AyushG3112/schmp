package main

import (
	"fmt"
	"os"
)

func main() {
	options := parseCLIFlags()
	validationErrors := options.validate()
	printValidationErrors(validationErrors)
	if len(validationErrors) != 0 {
		os.Exit(3)
	}
	m, err := process(options)
	if err != nil {
		fmt.Printf("could not process: %s\n", err.Error())
		os.Exit(4)
	}
	err = printDiffOutput(m, options)
	if err != nil {
		fmt.Printf("failed to print output: %s\n", err.Error())
		os.Exit(5)
	}
	if len(m) != 0 {
		os.Exit(6)
	}
}

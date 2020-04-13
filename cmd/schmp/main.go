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
		os.Exit(1)
	}
	m, err := process(options)
	if err != nil {
		fmt.Printf("could not process: %s\n", err.Error())
		os.Exit(3)
	}
	err = printDiffOutput(m, options)
	if err != nil {
		fmt.Printf("failed to print output: %s\n", err.Error())
		os.Exit(4)
	}
	if len(m) != 0 {
		os.Exit(5)
	}
}

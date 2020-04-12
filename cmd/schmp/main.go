package main

import (
	"fmt"
	"os"
)

func main() {
	options := parseCLIFlags()
	printValidationErrorsAndExit(options.validate())
	m, err := process(options)
	if err != nil {
		fmt.Printf("could not process: %s\n", err.Error())
		os.Exit(1)
	}
	printDiffOutputAndExit(m, options)
}

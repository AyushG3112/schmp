package main

import (
	"fmt"
	"os"
)

func printDiffOutputAndExit(m map[string][]string, options cliOptions) error {
	if len(m) == 0 {
		fmt.Println("No Diff, all files are exactly the same!")
		os.Exit(0)
	}
	fmt.Println(m)
	os.Exit(2)
	return nil
}

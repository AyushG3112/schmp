package main

import (
	"fmt"
)

func printDiffOutput(m map[string][]string, options cliOptions) error {
	if len(m) == 0 {
		fmt.Println("No Diff, all files are exactly the same!")
		return nil
	}
	fmt.Println(m)
	return nil
}

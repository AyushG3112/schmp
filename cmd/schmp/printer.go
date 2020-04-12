package main

import (
	"fmt"
	"os"
)

func printValidationErrorsAndExit(valErrors []string) {
	if len(valErrors) == 0 {
		return
	}
	fmt.Println("Could not process due to following errors: ")
	for _, v := range valErrors {
		fmt.Printf(" - %s\n", v)
	}
	os.Exit(1)
	return
}

package main

import (
	"fmt"
	"io"
)

func printValidationErrors(valErrors []string, w io.Writer) error {
	if len(valErrors) == 0 {
		return nil
	}
	if _, err := fmt.Fprintln(w, "Could not process due to following errors: "); err != nil {
		return err
	}
	for _, v := range valErrors {
		if _, err := fmt.Fprintf(w, " - %s\n", v); err != nil {
			return err
		}
	}
	return nil
}

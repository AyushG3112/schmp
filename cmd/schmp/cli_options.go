package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AyushG3112/schmp"
)

type cliOptions struct {
	mode    string
	outType string
	inFiles []string
}

func isOutTypeAllowed(outType string) bool {
	return outType == "stdout" || outType == "json"
}

func (c *cliOptions) validate() []string {
	valErrors := make([]string, 0)
	if !isOutTypeAllowed(c.outType) {
		valErrors = append(valErrors, fmt.Sprintf(`out-type "%s" is not allowed`, c.outType))
	}

	return valErrors
}

func (c *cliOptions) toProcessingOptions() (schmp.ProcessingOptions, error) {
	var err error
	options := schmp.ProcessingOptions{
		Mode: c.mode,
	}

	files := make([]io.ReadCloser, 0)
	for _, v := range c.inFiles {
		file, er := os.Open(v)
		if er != nil {
			err = er
			break
		}
		files = append(files, file)
	}

	if err != nil {
		for _, v := range files {
			v.Close()
		}
		return options, err
	}

	options.Sources = make([]io.Reader, len(files))

	for i, v := range files {
		options.Sources[i] = v
	}
	return options, err
}

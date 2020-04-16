package main

import (
	"io"

	"github.com/ayushg3112/schmp"
)

func process(options cliOptions) (schmp.ComparisonOutput, error) {
	processingOptions, err := options.toProcessingOptions()
	if err != nil {
		return schmp.ComparisonOutput{}, err
	}
	defer closeFileHandles(processingOptions.Sources)
	return schmp.Compare(processingOptions)
}

func closeFileHandles(files []io.Reader) {
	for _, r := range files {
		if v, ok := r.(io.ReadCloser); ok {
			v.Close()
		}
	}
}

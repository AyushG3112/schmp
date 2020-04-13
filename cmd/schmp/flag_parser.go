package main

import (
	"github.com/spf13/pflag"
)

func parseCLIFlags() cliOptions {
	mode := pflag.StringP("mode", "m", "json", "file format")
	outType := pflag.StringP("out-type", "o", "stdout", "Output format")
	outFile := pflag.String("out-file", "", "Output file. Only used if `--out-type` is not stdout")
	inFiles := pflag.StringArrayP("file", "f", []string{}, "Files to compare. Use this flag multiple times, once for each file.")

	pflag.Parse()

	options := cliOptions{
		mode:    *mode,
		outType: *outType,
		outFile: *outFile,
		inFiles: *inFiles,
	}
	return options
}

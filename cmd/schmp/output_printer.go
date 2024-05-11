package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/ayushg3112/schmp"
)

type typeData struct {
	File string
	Type string
}

type outputFormat struct {
	Field string
	Types []typeData
}

func (o *outputFormat) printStdout(stdout io.Writer) error {
	_, err := fmt.Fprintf(stdout, "%s:\n", o.Field)
	if err != nil {
		return err
	}
	for _, v := range o.Types {
		_, err = fmt.Fprintf(stdout, "\t%s: %s\n", v.File, v.Type)
		if err != nil {
			return err
		}
	}
	return nil
}

type outputList struct {
	outputs []outputFormat
}

func (o *outputList) printStdout(stdout io.Writer) error {
	if len(o.outputs) == 0 {
		_, err := fmt.Fprintln(stdout, "No Diff, all files are exactly the same!")
		return err
	}
	_, err := fmt.Fprintf(stdout, "%d differences found:\n\n", len(o.outputs))
	if err != nil {
		return err
	}
	for i, v := range o.outputs {
		err = v.printStdout(stdout)
		if err != nil {
			return err
		}
		if i != len(o.outputs)-1 {
			_, err = fmt.Fprint(stdout, "\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *outputList) writeJSON(stdout io.Writer, outFile string) error {
	if len(o.outputs) == 0 {
		_, err := fmt.Fprintln(stdout, "No Diff, all files are exactly the same!")
		if err != nil {
			return err
		}
		return nil
	}
	bytes, err := json.MarshalIndent(o.outputs, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(outFile, bytes, 0644)
}

func (o *outputList) fromComparisonDiff(diff map[string][]string, options cliOptions) {
	op := make([]outputFormat, len(diff))
	i := 0
	for k, v := range diff {
		op[i].Field = k
		op[i].Types = make([]typeData, len(v))
		for idx, t := range v {
			op[i].Types[idx].File = options.inFiles[idx]
			op[i].Types[idx].Type = getOutputType(t)
		}
		i++
	}

	sort.Slice(op, func(i, j int) bool {
		return op[i].Field < op[j].Field
	})
	o.outputs = op
}

func getOutputType(t string) string {
	if strings.HasPrefix(t, "map") {
		return "object"
	}
	if strings.HasPrefix(t, "[]") {
		return "array"
	}
	if strings.HasPrefix(t, "int") {
		return "integer"
	}
	if strings.HasPrefix(t, "float") {
		return "float"
	}
	if strings.HasPrefix(t, "<nil>") {
		return "null"
	}
	return t
}

func printDiffOutput(result schmp.ComparisonOutput, options cliOptions, stdout io.Writer) error {
	m := result.Diff
	o := outputList{}
	o.fromComparisonDiff(m, options)
	if options.outType == "stdout" {
		return o.printStdout(stdout)
	}

	if options.outType == "json" {
		return o.writeJSON(stdout, options.outFile)
	}
	return fmt.Errorf("outType '%s' not supported", options.outType)
}

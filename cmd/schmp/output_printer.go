package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

// writeJSON renders the diff as indented JSON. With an out-file it is written
// there (always, even for an empty "[]" diff, so the artifact a pipeline expects
// always exists); without one it goes to stdout.
func (o *outputList) writeJSON(stdout io.Writer, outFile string) error {
	bytes, err := json.MarshalIndent(o.outputs, "", "  ")
	if err != nil {
		return err
	}
	if outFile == "" {
		_, err := fmt.Fprintf(stdout, "%s\n", bytes)
		return err
	}
	return os.WriteFile(outFile, bytes, 0644)
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

// getOutputType maps the Go type names produced by the comparator to friendly
// names (object, array, integer, float, null). A heterogeneous array surfaces
// as a " | "-joined set of type names, so each member is mapped individually.
func getOutputType(t string) string {
	if t == "" {
		return ""
	}
	parts := strings.Split(t, " | ")
	for i, part := range parts {
		parts[i] = friendlyTypeName(part)
	}
	return strings.Join(parts, " | ")
}

func friendlyTypeName(t string) string {
	switch {
	case strings.HasPrefix(t, "map"):
		return "object"
	case strings.HasPrefix(t, "[]"):
		return "array"
	case strings.HasPrefix(t, "int"), strings.HasPrefix(t, "uint"):
		return "integer"
	case strings.HasPrefix(t, "float"):
		return "float"
	case t == "<nil>":
		return "null"
	default:
		return t
	}
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

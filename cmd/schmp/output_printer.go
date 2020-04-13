package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type typeData struct {
	File string
	Type string
}

type outputFormat struct {
	Field string
	Types []typeData
}

func (o *outputFormat) printStdout() {
	fmt.Printf("%s:\n", o.Field)
	for _, v := range o.Types {
		fmt.Printf("\t%s: %s\n", v.File, v.Type)
	}
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
	if strings.HasPrefix(t, "nil") {
		return "null"
	}
	return t
}

func toOutputFormat(m map[string][]string, options cliOptions) []outputFormat {
	op := make([]outputFormat, len(m))
	i := 0
	for k, v := range m {
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
	return op
}

func printDiffOutput(m map[string][]string, options cliOptions) error {
	if len(m) == 0 {
		fmt.Println("No Diff, all files are exactly the same!")
		return nil
	}
	output := toOutputFormat(m, options)
	if options.outType == "stdout" {
		fmt.Printf("%d differences found:\n\n", len(output))
		for _, v := range output {
			v.printStdout()
			fmt.Print("\n")
		}
		return nil
	}

	if options.outType == "json" {
		bytes, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(options.outFile, bytes, 0644)
		return err

	}
	return nil
}

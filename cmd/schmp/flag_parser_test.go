package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestFlagsAreParsedCorrectly(t *testing.T) {
	os.Args = []string{"schmp", "-f", "testdata/yaml/yaml_1.yaml", "--file", "testdata/yaml/yaml_2.yaml", "--mode", "yaml", "-o", "json", "--out-file", "sample.json"}
	cliOptions := parseCLIFlags()
	if cliOptions.mode != "yaml" {
		t.Fatalf("could not set mode")
	}
	if cliOptions.outType != "json" {
		t.Fatalf("could not set outType")
	}
	if cliOptions.outFile != "sample.json" {
		t.Fatalf("could not set outFile")
	}
	if !reflect.DeepEqual(cliOptions.inFiles, []string{"testdata/yaml/yaml_1.yaml", "testdata/yaml/yaml_2.yaml"}) {
		fmt.Println("could not set inFiles")
	}
}

func TestCheckDefaultFlags(t *testing.T) {
	os.Args = []string{"schmp"}
	cliOptions := parseCLIFlags()
	if cliOptions.mode != "json" {
		t.Fatalf("default mode is not `json`")
	}
	if cliOptions.outType != "stdout" {
		t.Fatalf("default mode is not `stdout`")
	}
}

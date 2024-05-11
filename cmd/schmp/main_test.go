package main

import (
	"fmt"
	"os"
	"testing"
)

func TestStartCLIFlagsValidationErrorsPrintFailed(t *testing.T) {
	m := &mockWriterError{msg: "TestStartCLIFlagsValidationErrorsPrintFailed"}
	os.Args = []string{"schmp", "-o", "random"}
	exitCode := start(m)
	expected := "failed to print validation errors: mock error: " + m.msg + "\n"
	if exitCode != 3 || expected != m.data {
		t.Log(m.data)
		t.Log(expected)
		t.Fatalf("Invalid exit code: %d", exitCode)
	}
}

func TestStartCLIFlagsValidationErrors(t *testing.T) {
	m := &mockWriter{}
	os.Args = []string{"schmp", "-o", "random"}
	exitCode := start(m)
	expected := "Could not process due to following errors: \n - out-type \"random\" is not allowed\n"
	if exitCode != 4 || m.data != expected {
		t.Log(m.data)
		t.Log(expected)
		t.Fatalf("Invalid exit code: %d", exitCode)
	}
}

func TestStartProcessFailedError(t *testing.T) {
	m := &mockWriter{}
	os.Args = []string{"schmp", "-m", "random"}
	exitCode := start(m)
	expected := "could not process: mode \"random\" is not supported\n"
	if exitCode != 5 || m.data != expected {
		t.Log(m.data)
		t.Log(expected)
		t.Fatalf("Invalid exit code: %d", exitCode)
	}
}

func TestSuccessfulResultOnDiffStdout(t *testing.T) {
	m := &mockWriter{}
	os.Args = []string{"schmp", "-f", "../../testdata/json/json_1.json", "-f", "../../testdata/json/json_2.json"}
	exitCode := start(m)
	if exitCode != 7 {
		t.Fatalf("Invalid exit code: %d", exitCode)
	}
}

func TestSuccessfulResultOnNoDiffStdout(t *testing.T) {
	m := &mockWriter{}
	os.Args = []string{"schmp", "-f", "../../testdata/json/json_1.json", "-f", "../../testdata/json/json_1.json"}
	exitCode := start(m)
	expected := "No Diff, all files are exactly the same!\n"
	if exitCode != 0 || expected != m.data {
		fmt.Println(m)
		t.Fatalf("Invalid exit code: %d", exitCode)
	}
}

// func start(stdout io.Writer) int {
// 	options := parseCLIFlags()
// 	validationErrors := options.validate()
// 	printValidationErrors(validationErrors, stdout)
// 	if len(validationErrors) != 0 {
// 		return 3
// 	}
// 	m, err := process(options)
// 	if err != nil {
// 		fmt.Fprintf(stdout, "could not process: %s\n", err.Error())
// 		return 4
// 	}
// 	err = printDiffOutput(m, options, stdout)
// 	if err != nil {
// 		fmt.Fprintf(stdout, "failed to print output: %s\n", err.Error())
// 		return 5
// 	}
// 	if len(m.Diff) != 0 {
// 		return 6
// 	}
// 	return 0
// }

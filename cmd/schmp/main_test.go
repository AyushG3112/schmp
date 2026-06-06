package main

import (
	"fmt"
	"os"
	"testing"
)

func TestStartCLIFlagsValidationErrorsPrintFailed(t *testing.T) {
	stderr := &mockWriterError{msg: "TestStartCLIFlagsValidationErrorsPrintFailed"}
	os.Args = []string{"schmp", "-o", "random"}
	exitCode := start(&mockWriter{}, stderr)
	expected := "failed to print validation errors: mock error: " + stderr.msg + "\n"
	if exitCode != 3 || expected != stderr.data {
		t.Log(stderr.data)
		t.Log(expected)
		t.Fatalf("Invalid exit code: %d", exitCode)
	}
}

func TestStartCLIFlagsValidationErrors(t *testing.T) {
	stderr := &mockWriter{}
	os.Args = []string{"schmp", "-o", "random"}
	exitCode := start(&mockWriter{}, stderr)
	expected := "Could not process due to following errors: \n - out-type \"random\" is not allowed\n"
	if exitCode != 4 || stderr.data != expected {
		t.Log(stderr.data)
		t.Log(expected)
		t.Fatalf("Invalid exit code: %d", exitCode)
	}
}

func TestStartProcessFailedError(t *testing.T) {
	stderr := &mockWriter{}
	os.Args = []string{"schmp", "-m", "random"}
	exitCode := start(&mockWriter{}, stderr)
	expected := "could not process: mode \"random\" is not supported\n"
	if exitCode != 5 || stderr.data != expected {
		t.Log(stderr.data)
		t.Log(expected)
		t.Fatalf("Invalid exit code: %d", exitCode)
	}
}

func TestSuccessfulResultOnDiffStdout(t *testing.T) {
	os.Args = []string{"schmp", "-f", "../../testdata/json/json_1.json", "-f", "../../testdata/json/json_2.json"}
	exitCode := start(&mockWriter{}, &mockWriter{})
	if exitCode != 7 {
		t.Fatalf("Invalid exit code: %d", exitCode)
	}
}

func TestSuccessfulResultOnNoDiffStdout(t *testing.T) {
	stdout := &mockWriter{}
	os.Args = []string{"schmp", "-f", "../../testdata/json/json_1.json", "-f", "../../testdata/json/json_1.json"}
	exitCode := start(stdout, &mockWriter{})
	expected := "No Diff, all files are exactly the same!\n"
	if exitCode != 0 || expected != stdout.data {
		fmt.Println(stdout)
		t.Fatalf("Invalid exit code: %d", exitCode)
	}
}

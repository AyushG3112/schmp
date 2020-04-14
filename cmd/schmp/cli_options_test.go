package main

import (
	"os"
	"reflect"
	"testing"
)

func TestCLIOptionsValidations(t *testing.T) {
	t.Run("invalid outType", func(t *testing.T) {
		options := cliOptions{
			outType: "xml",
		}
		expected := []string{`out-type "xml" is not allowed`}

		validationErrors := options.validate()
		if !reflect.DeepEqual(validationErrors, expected) {
			t.Fatalf("expected error not found:\nExpected: %s\nActual: %s", validationErrors, expected)
		}
	})

	t.Run("outType set no outfile", func(t *testing.T) {
		options := cliOptions{
			outType: "json",
		}
		expected := []string{"`out-file` is required if `out-type` is not stdout"}
		validationErrors := options.validate()
		if !reflect.DeepEqual(validationErrors, expected) {
			t.Fatalf("expected error not found:\nExpected: %s\nActual: %s", validationErrors, expected)
		}
	})

	t.Run("outFile is optional if outType is stdout", func(t *testing.T) {
		options := cliOptions{
			outType: "stdout",
		}
		expected := []string{}
		validationErrors := options.validate()
		if !reflect.DeepEqual(validationErrors, expected) {
			t.Fatalf("expected error not found:\nExpected: %s\nActual: %s", validationErrors, expected)
		}
	})
}

func TestCLIOptionsToProcessingOptions(t *testing.T) {
	options := cliOptions{
		outType: "stdout",
		mode:    "yaml",
		inFiles: []string{
			"../../testdata/yaml/yaml_1.yaml",
			"../../testdata/yaml/yaml_2.yaml",
		},
	}

	processingOptions, err := options.toProcessingOptions()
	if err != nil {
		t.Fatalf("failed to convert to ProcessingOptions: %s", err.Error())
	}
	if processingOptions.Mode != options.mode {
		t.Fatalf("Could not set mode")
	}
	if len(processingOptions.Sources) != len(options.inFiles) {
		t.Fatalf("sources length did not match")
	}
	for _, v := range processingOptions.Sources {
		f := v.(*os.File)
		err := f.Close()
		if err != nil {
			t.Fatalf("failed to close file: %s", err.Error())
		}
	}
}

func TestCLIOptionsToProcessingOptionsFailsIfFilePathInvalid(t *testing.T) {
	options := cliOptions{
		outType: "stdout",
		mode:    "yaml",
		inFiles: []string{
			"testdata/yaml/yaml_1.yaml",
			"testdata/yaml/yaml_2.yaml",
		},
	}

	_, err := options.toProcessingOptions()
	if err == nil {
		t.Fatalf("no error returned")
	}
}

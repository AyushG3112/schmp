package main

import (
	"reflect"
	"testing"

	"github.com/AyushG3112/schmp"
)

func TestProcessShouldReturnErrorWhenToProcessingOptionsFails(t *testing.T) {
	_, err := process(cliOptions{
		inFiles: []string{
			"random.json",
		},
	})
	if err == nil {
		t.Fatal("did not forward error")
	}
}

func TestProcessShouldReturnCorrectDiffForJSON(t *testing.T) {
	result, err := process(cliOptions{
		mode: "json",
		inFiles: []string{
			"../../testdata/json/json_1.json",
			"../../testdata/json/json_2.json",
		},
	})
	if err != nil {
		t.Fatalf("process failed: %s", err.Error())
	}

	expected := schmp.ComparisonOutput{
		Diff: map[string][]string{
			"quiz.arts": {
				"map[string]interface {}",
				"",
			},
			"quiz.maths.q2.options": {
				"[]interface {}",
				"map[string]interface {}",
			},
			"quiz.maths.q2.question": {
				"float64",
				"string",
			},
			"quiz.science": {
				"",
				"map[string]interface {}",
			},
			"quiz.sport.q1.answer": {
				"nil",
				"string",
			},
			"quiz.sport.q1.options": {
				"[]interface {}",
				"",
			},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

func TestProcessShouldReturnCorrectDiffForYAML(t *testing.T) {
	result, err := process(cliOptions{
		mode: "yaml",
		inFiles: []string{
			"../../testdata/yaml/yaml_1.yaml",
			"../../testdata/yaml/yaml_2.yaml",
		},
	})
	if err != nil {
		t.Fatalf("process failed: %s", err.Error())
	}

	expected := schmp.ComparisonOutput{
		Diff: map[string][]string{
			"quiz.arts": {
				"map[string]interface {}",
				"",
			},
			"quiz.maths.q2.options": {
				"[]interface {}",
				"map[string]interface {}",
			},
			"quiz.maths.q2.question": {
				"float64",
				"string",
			},
			"quiz.science": {
				"",
				"map[string]interface {}",
			},
			"quiz.sport.q1.answer": {
				"nil",
				"string",
			},
			"quiz.sport.q1.options": {
				"[]interface {}",
				"",
			},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

func TestProcessShouldReturnCorrectDiffForTOML(t *testing.T) {
	result, err := process(cliOptions{
		mode: "toml",
		inFiles: []string{
			"../../testdata/toml/toml_1.toml",
			"../../testdata/toml/toml_2.toml",
		},
	})
	if err != nil {
		t.Fatalf("process failed: %s", err.Error())
	}

	expected := schmp.ComparisonOutput{
		Diff: map[string][]string{
			"quiz.arts": {
				"map[string]interface {}",
				"",
			},
			"quiz.maths.q2.options": {
				"[]interface {}",
				"map[string]interface {}",
			},
			"quiz.maths.q2.question": {
				"float64",
				"string",
			},
			"quiz.science": {
				"",
				"map[string]interface {}",
			},
			"quiz.sport.q1.answer": {
				"",
				"string",
			},
			"quiz.sport.q1.options": {
				"[]interface {}",
				"",
			},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

// func process(options cliOptions) (schmp.ComparisonOutput, error) {
// 	processingOptions, err := options.toProcessingOptions()
// 	if err != nil {
// 		return schmp.ComparisonOutput{}, err
// 	}
// 	defer closeFileHandles(processingOptions.Sources)
// 	return schmp.Compare(processingOptions)
// }

// func closeFileHandles(files []io.Reader) {
// 	for _, r := range files {
// 		if v, ok := r.(io.ReadCloser); ok {
// 			v.Close()
// 		}
// 	}
// }

package schmp

import (
	"errors"
	"io"
	"os"
	"reflect"
	"testing"
)

type MockErrorReader struct {
}

func (m MockErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("failed to read")
}

func TestCompareJSON(t *testing.T) {
	f1, err := os.Open("testdata/json/json_1.json")
	if err != nil {
		t.Fatalf("error while opening file: %s", err.Error())
	} else {
		defer f1.Close()
	}
	f2, err := os.Open("testdata/json/json_2.json")
	if err != nil {
		t.Fatalf("error while opening file: %s", err.Error())
	} else {
		defer f2.Close()
	}
	result, err := Compare(ProcessingOptions{
		Mode:    "json",
		Sources: []io.Reader{f1, f2},
	})
	if err != nil {
		t.Fatalf("failed to compare: %s", err.Error())
	}
	expected := map[string][]string{
		"quiz.arts": {
			"map[string]interface {}",
			"",
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
		"quiz.maths.q2.options": {
			"[]interface {}",
			"map[string]interface {}",
		},
		"quiz.maths.q2.question": {
			"float64",
			"string",
		},
	}
	if !reflect.DeepEqual(expected, result.Diff) {
		t.Fatalf("expected and actual did not match")
	}
}

func TestCompareTOML(t *testing.T) {
	f1, err := os.Open("testdata/toml/toml_1.toml")
	if err != nil {
		t.Fatalf("error while opening file: %s", err.Error())
	} else {
		defer f1.Close()
	}
	f2, err := os.Open("testdata/toml/toml_2.toml")
	if err != nil {
		t.Fatalf("error while opening file: %s", err.Error())
	} else {
		defer f2.Close()
	}
	result, err := Compare(ProcessingOptions{
		Mode:    "toml",
		Sources: []io.Reader{f1, f2},
	})
	if err != nil {
		t.Fatalf("failed to compare: %s", err.Error())
	}
	expected := map[string][]string{
		"quiz.arts": {
			"map[string]interface {}",
			"",
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
		"quiz.maths.q2.options": {
			"[]interface {}",
			"map[string]interface {}",
		},
		"quiz.maths.q2.question": {
			"float64",
			"string",
		},
	}
	if !reflect.DeepEqual(expected, result.Diff) {
		t.Fatalf("expected and actual did not match")
	}
}

func TestCompareYAML(t *testing.T) {
	f1, err := os.Open("testdata/yaml/yaml_1.yaml")
	if err != nil {
		t.Fatalf("error while opening file: %s", err.Error())
	} else {
		defer f1.Close()
	}
	f2, err := os.Open("testdata/yaml/yaml_2.yaml")
	if err != nil {
		t.Fatalf("error while opening file: %s", err.Error())
	} else {
		defer f2.Close()
	}
	result, err := Compare(ProcessingOptions{
		Mode:    "yaml",
		Sources: []io.Reader{f1, f2},
	})
	if err != nil {
		t.Fatalf("failed to compare: %s", err.Error())
	}
	expected := map[string][]string{
		"quiz.arts": {
			"map[string]interface {}",
			"",
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
		"quiz.maths.q2.options": {
			"[]interface {}",
			"map[string]interface {}",
		},
		"quiz.maths.q2.question": {
			"float64",
			"string",
		},
	}
	if !reflect.DeepEqual(expected, result.Diff) {
		t.Fatalf("expected and actual did not match")
	}
}

func TestShouldReturnErrorIfParserReturnsError(t *testing.T) {
	_, err := Compare(ProcessingOptions{
		Mode:    "yaml",
		Sources: []io.Reader{MockErrorReader{}, MockErrorReader{}},
	})
	if err == nil {
		t.Fatalf("did not fail")
	}
}

func TestShouldReturnErrorOnInvalidMode(t *testing.T) {
	f1, err := os.Open("testdata/json/json_1.json")
	if err != nil {
		t.Fatalf("error while opening file: %s", err.Error())
	} else {
		defer f1.Close()
	}
	f2, err := os.Open("testdata/json/json_2.json")
	if err != nil {
		t.Fatalf("error while opening file: %s", err.Error())
	} else {
		defer f2.Close()
	}
	_, err = Compare(ProcessingOptions{
		Mode:    "invalid",
		Sources: []io.Reader{f1, f2},
	})
	if err == nil {
		t.Fatalf("did not fail")
	}
}

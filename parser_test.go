package schmp

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestGetParser(t *testing.T) {
	t.Run("get json parser", func(t *testing.T) {
		parser, err := getParser("json")
		if err != nil {
			t.Fatalf("error while getting parser: %s", err.Error())
		}
		if parser == nil {
			t.Fatalf("parser is nil")
		}
	})

	t.Run("get toml parser", func(t *testing.T) {
		parser, err := getParser("yaml")
		if err != nil {
			t.Fatalf("error while getting parser: %s", err.Error())
		}
		if parser == nil {
			t.Fatalf("parser is nil")
		}
	})

	t.Run("get yaml parser", func(t *testing.T) {
		parser, err := getParser("toml")
		if err != nil {
			t.Fatalf("error while getting parser: %s", err.Error())
		}
		if parser == nil {
			t.Fatalf("parser is nil")
		}
	})

	t.Run("get invalid parser", func(t *testing.T) {
		_, err := getParser("invalid")
		if err == nil {
			t.Fatalf("error did mot occure while getting invalid parser")
		}
	})
}

func TestJSONParser(t *testing.T) {
	parser, err := getParser("json")
	if err != nil {
		t.Fatalf("error while getting parser: %s", err.Error())
	}
	if parser == nil {
		t.Fatalf("parser is nil")
	}
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
	maps, err := parser(ProcessingOptions{
		Mode:    "json",
		Sources: []io.Reader{f1, f2},
	})
	if err != nil {
		t.Fatalf("error while parsing file: %s", err.Error())
	}

	expected := []map[string]interface{}{
		{
			"quiz": map[string]interface{}{
				"arts": map[string]interface{}{
					"question": nil,
				},
				"sport": map[string]interface{}{
					"q1": map[string]interface{}{
						"question": "Which one is correct team name in NBA?",
						"options": []interface{}{
							"New York Bulls",
							"Los Angeles Kings",
							"Golden State Warriros",
							"Huston Rocket",
						},
						"answer": nil,
					},
				},
				"maths": map[string]interface{}{
					"q1": map[string]interface{}{
						"question": "5 + 7 = ?",
						"options": []interface{}{
							"10",
							"11",
							"12",
							"13",
						},
						"answer": "12",
					},
					"q2": map[string]interface{}{
						"question": 100.0,
						"options": []interface{}{
							"1",
							"2",
							"3",
							"4",
						},
						"answer": "4",
					},
				},
			},
		},
		{
			"quiz": map[string]interface{}{
				"science": map[string]interface{}{
					"question": nil,
				},
				"sport": map[string]interface{}{
					"q1": map[string]interface{}{
						"question": "Which one is correct team name in NBA?",
						"answer":   "Huston Rocket",
					},
				},
				"maths": map[string]interface{}{
					"q1": map[string]interface{}{
						"question": "5 + 7 = ?",
						"options": []interface{}{
							"10",
							"11",
							"12",
							"13",
						},
						"answer": "12",
					},
					"q2": map[string]interface{}{
						"question": "12 - 8 = ?",
						"options":  map[string]interface{}{},
						"answer":   "4",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(expected, maps) {
		t.Fatalf("expected and actual did not match")
	}
}

func TestYAMLParser(t *testing.T) {
	parser, err := getParser("yaml")
	if err != nil {
		t.Fatalf("error while getting parser: %s", err.Error())
	}
	if parser == nil {
		t.Fatalf("parser is nil")
	}
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
	maps, err := parser(ProcessingOptions{
		Mode:    "yaml",
		Sources: []io.Reader{f1, f2},
	})
	if err != nil {
		t.Fatalf("error while parsing file: %s", err.Error())
	}

	expected := []map[string]interface{}{
		{
			"quiz": map[string]interface{}{
				"arts": map[string]interface{}{
					"question": nil,
				},
				"sport": map[string]interface{}{
					"q1": map[string]interface{}{
						"question": "Which one is correct team name in NBA?",
						"options": []interface{}{
							"New York Bulls",
							"Los Angeles Kings",
							"Golden State Warriros",
							"Huston Rocket",
						},
						"answer": nil,
					},
				},
				"maths": map[string]interface{}{
					"q1": map[string]interface{}{
						"question": "5 + 7 = ?",
						"options": []interface{}{
							"10",
							"11",
							"12",
							"13",
						},
						"answer": "12",
					},
					"q2": map[string]interface{}{
						"question": 100.0,
						"options": []interface{}{
							"1",
							"2",
							"3",
							"4",
						},
						"answer": "4",
					},
				},
			},
		},
		{
			"quiz": map[string]interface{}{
				"science": map[string]interface{}{
					"question": nil,
				},
				"sport": map[string]interface{}{
					"q1": map[string]interface{}{
						"question": "Which one is correct team name in NBA?",
						"answer":   "Huston Rocket",
					},
				},
				"maths": map[string]interface{}{
					"q1": map[string]interface{}{
						"question": "5 + 7 = ?",
						"options": []interface{}{
							"10",
							"11",
							"12",
							"13",
						},
						"answer": "12",
					},
					"q2": map[string]interface{}{
						"question": "12 - 8 = ?",
						"options":  map[string]interface{}{},
						"answer":   "4",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(expected, maps) {
		t.Fatalf("expected and actual did not match")
	}
}

func TestTOMLParser(t *testing.T) {
	parser, err := getParser("toml")
	if err != nil {
		t.Fatalf("error while getting parser: %s", err.Error())
	}
	if parser == nil {
		t.Fatalf("parser is nil")
	}
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
	maps, err := parser(ProcessingOptions{
		Mode:    "toml",
		Sources: []io.Reader{f1, f2},
	})
	if err != nil {
		t.Fatalf("error while parsing file: %s", err.Error())
	}

	expected := []map[string]interface{}{
		{
			"quiz": map[string]interface{}{
				"arts": map[string]interface{}{},
				"sport": map[string]interface{}{
					"q1": map[string]interface{}{
						"question": "Which one is correct team name in NBA?",
						"options": []interface{}{
							"New York Bulls",
							"Los Angeles Kings",
							"Golden State Warriros",
							"Huston Rocket",
						},
					},
				},
				"maths": map[string]interface{}{
					"q1": map[string]interface{}{
						"question": "5 + 7 = ?",
						"options": []interface{}{
							"10",
							"11",
							"12",
							"13",
						},
						"answer": "12",
					},
					"q2": map[string]interface{}{
						"question": 100.0,
						"options": []interface{}{
							"1",
							"2",
							"3",
							"4",
						},
						"answer": "4",
					},
				},
			},
		},
		{
			"quiz": map[string]interface{}{
				"science": map[string]interface{}{},
				"sport": map[string]interface{}{
					"q1": map[string]interface{}{
						"question": "Which one is correct team name in NBA?",
						"answer":   "Huston Rocket",
					},
				},
				"maths": map[string]interface{}{
					"q1": map[string]interface{}{
						"question": "5 + 7 = ?",
						"options": []interface{}{
							"10",
							"11",
							"12",
							"13",
						},
						"answer": "12",
					},
					"q2": map[string]interface{}{
						"question": "12 - 8 = ?",
						"options":  map[string]interface{}{},
						"answer":   "4",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(expected, maps) {
		t.Fatalf("expected and actual did not match")
	}
}

func TestParserShouldReturnErrorWhenParsingFails(t *testing.T) {
	parser, err := getParser("yaml")
	if err != nil {
		t.Fatalf("error while getting parser: %s", err.Error())
	}
	if parser == nil {
		t.Fatalf("parser is nil")
	}
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
	_, err = parser(ProcessingOptions{
		Mode:    "yaml",
		Sources: []io.Reader{f1, f2},
	})
	if err == nil {
		t.Fatalf("no error occured during invalid parsing")
	}

}

func TestParserShouldReturnErrorWhenReadingFails(t *testing.T) {

	parser, err := getParser("yaml")
	if err != nil {
		t.Fatalf("error while getting parser: %s", err.Error())
	}
	if parser == nil {
		t.Fatalf("parser is nil")
	}
	_, err = parser(ProcessingOptions{
		Mode:    "yaml",
		Sources: []io.Reader{MockErrorReader{}, MockErrorReader{}},
	})
	if err == nil {
		t.Fatalf("no error occured during invalid reading")
	}

}

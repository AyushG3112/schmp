package schmp

import (
	"reflect"
	"testing"
)

func TestJSONCompareHasDiff(t *testing.T) {
	m1 := map[string]interface{}{
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
					"question": 100,
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
	}

	m2 := map[string]interface{}{
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
			"<nil>",
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
			"int",
			"string",
		},
	}
	actual, err := compare([]map[string]interface{}{m1, m2})
	if err != nil {
		t.Fatalf("failed to compare: %s", err.Error())
	}
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Fatalf("Actual and expected did not match.\n Actual: \n%v\n\nExpected: \n%v", actual, expected)
	}
}

func TestJSONCompareNoDiff(t *testing.T) {
	m1 := map[string]interface{}{
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
					"question": 100,
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
	}

	m2 := map[string]interface{}{
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
					"question": 100,
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
	}

	expected := map[string][]string{}
	actual, err := compare([]map[string]interface{}{m1, m2})
	if err != nil {
		t.Fatalf("failed to compare: %s", err.Error())
	}
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Fatalf("Actual and expected did not match.\n Actual: \n%v\n\nExpected: \n%v", actual, expected)
	}
}

func TestJSONCompareHasDiffWorksWithNilMaps(t *testing.T) {
	var m1 map[string]interface{}
	m2 := map[string]interface{}{
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
	}

	expected := map[string][]string{
		"quiz": {
			"",
			"map[string]interface {}",
		},
	}
	actual, err := compare([]map[string]interface{}{m1, m2})
	if err != nil {
		t.Fatalf("failed to compare: %s", err.Error())
	}
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Fatalf("Actual and expected did not match.\n Actual: \n%v\n\nExpected: \n%v", actual, expected)
	}
}

func TestJSONCompareHasDiffWorksWithAllNilMaps(t *testing.T) {
	var m1 map[string]interface{}
	var m2 map[string]interface{}
	var m3 map[string]interface{}

	expected := map[string][]string{}
	actual, err := compare([]map[string]interface{}{m1, m2, m3})
	if err != nil {
		t.Fatalf("failed to compare: %s", err.Error())
	}
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Fatalf("Actual and expected did not match.\n Actual: \n%v\n\nExpected: \n%v", actual, expected)
	}
}

func assertCompareEqual(t *testing.T, maps []map[string]interface{}, expected map[string][]string) {
	t.Helper()
	actual, err := compare(maps)
	if err != nil {
		t.Fatalf("failed to compare: %s", err.Error())
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Actual and expected did not match.\n Actual: \n%v\n\nExpected: \n%v", actual, expected)
	}
}

// Array elements are compared by schema at the synthetic "<path>[]" key.
func TestCompareArrayElementSchemaDiff(t *testing.T) {
	m1 := map[string]interface{}{"tags": []interface{}{"x", "y"}}
	m2 := map[string]interface{}{"tags": []interface{}{1, 2, 3}}
	assertCompareEqual(t, []map[string]interface{}{m1, m2}, map[string][]string{
		"tags[]": {"string", "int"},
	})
}

// Differing lengths alone are not a diff: the element schema is identical.
func TestCompareArrayLengthInsensitive(t *testing.T) {
	m1 := map[string]interface{}{"tags": []interface{}{1, 2, 3}}
	m2 := map[string]interface{}{"tags": []interface{}{1}}
	assertCompareEqual(t, []map[string]interface{}{m1, m2}, map[string][]string{})
}

// Objects inside arrays are traversed, so per-field drift is caught.
func TestCompareArrayOfObjectsDiff(t *testing.T) {
	m1 := map[string]interface{}{"users": []interface{}{
		map[string]interface{}{"id": 1},
	}}
	m2 := map[string]interface{}{"users": []interface{}{
		map[string]interface{}{"id": "x"},
		map[string]interface{}{"id": "y"},
	}}
	assertCompareEqual(t, []map[string]interface{}{m1, m2}, map[string][]string{
		"users[].id": {"int", "string"},
	})
}

// A heterogeneous array collapses to the set of its element types.
func TestCompareHeterogeneousArrayCollapsesToSet(t *testing.T) {
	m1 := map[string]interface{}{"x": []interface{}{1, "a"}}
	m2 := map[string]interface{}{"x": []interface{}{1}}
	assertCompareEqual(t, []map[string]interface{}{m1, m2}, map[string][]string{
		"x[]": {"int | string", "int"},
	})
}

// When an array's own type differs (array vs scalar), the element key is
// suppressed rather than reported on top of the parent.
func TestCompareArrayVsScalarSuppressesElements(t *testing.T) {
	m1 := map[string]interface{}{"a": []interface{}{1, 2}}
	m2 := map[string]interface{}{"a": 5}
	assertCompareEqual(t, []map[string]interface{}{m1, m2}, map[string][]string{
		"a": {"[]interface {}", "int"},
	})
}

// Regression for the empty-string-key suppression bug: a differing "" key must
// not hide unrelated top-level diffs.
func TestCompareEmptyStringKeyDoesNotSuppressSiblings(t *testing.T) {
	m1 := map[string]interface{}{"": 1, "port": 8080}
	m2 := map[string]interface{}{"": "a", "port": "8080"}
	assertCompareEqual(t, []map[string]interface{}{m1, m2}, map[string][]string{
		"":     {"int", "string"},
		"port": {"int", "string"},
	})
}

// A map value the decoders never produce surfaces as an error instead of a
// silently wrong (partial) result.
func TestCompareUnsupportedMapTypeReturnsError(t *testing.T) {
	m1 := map[string]interface{}{"cfg": map[string]int{"a": 1}}
	m2 := map[string]interface{}{"cfg": map[string]interface{}{"a": "x"}}
	if _, err := compare([]map[string]interface{}{m1, m2}); err == nil {
		t.Fatal("expected an error for an unsupported map type, got nil")
	}
}

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
			"int",
			"string",
		},
	}
	actual, err := compare([]map[string]interface{}{m1, m2}, ProcessingOptions{}, "", make(map[string][]string))
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
	actual, err := compare([]map[string]interface{}{m1, m2}, ProcessingOptions{}, "", make(map[string][]string))
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
	actual, err := compare([]map[string]interface{}{m1, m2}, ProcessingOptions{}, "", make(map[string][]string))
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
	actual, err := compare([]map[string]interface{}{m1, m2, m3}, ProcessingOptions{}, "", make(map[string][]string))
	if err != nil {
		t.Fatalf("failed to compare: %s", err.Error())
	}
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Fatalf("Actual and expected did not match.\n Actual: \n%v\n\nExpected: \n%v", actual, expected)
	}
}

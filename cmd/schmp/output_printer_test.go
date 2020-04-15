package main

import (
	"reflect"
	"testing"

	"github.com/AyushG3112/schmp"
)

func failIfStringsNotEqual(s1, s2 string, t *testing.T) {
	if s1 != s2 {
		t.Fatalf("strings did not match: '%s' and '%s'", s1, s2)
	}
}

func TestGetOutputType(t *testing.T) {

	t.Run("map should return object", func(t *testing.T) {
		failIfStringsNotEqual(getOutputType("map[string]interface {}"), "object", t)
	})

	t.Run("slice should return array", func(t *testing.T) {
		failIfStringsNotEqual(getOutputType("[]map[string]interface {}"), "array", t)
	})

	t.Run("int should return integer", func(t *testing.T) {
		failIfStringsNotEqual(getOutputType("int32"), "integer", t)
		failIfStringsNotEqual(getOutputType("int64"), "integer", t)
	})

	t.Run("float should return float", func(t *testing.T) {
		failIfStringsNotEqual(getOutputType("float32"), "float", t)
		failIfStringsNotEqual(getOutputType("float64"), "float", t)
	})

	t.Run("nil should return null", func(t *testing.T) {
		failIfStringsNotEqual(getOutputType("nil"), "null", t)
	})

	t.Run("string should return string", func(t *testing.T) {
		failIfStringsNotEqual(getOutputType("string"), "string", t)
	})

	t.Run("anything else should return itself", func(t *testing.T) {
		failIfStringsNotEqual(getOutputType("random"), "random", t)
	})
}

func TestToOutputFormat(t *testing.T) {
	actual := outputList{}
	actual.fromComparisonDiff(map[string][]string{
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
	}, cliOptions{
		inFiles: []string{"a.json", "b.json"},
	})

	expected := outputList{
		outputs: []outputFormat{
			{
				Field: "quiz.arts",
				Types: []typeData{
					{
						File: "a.json",
						Type: "object",
					},
					{
						File: "b.json",
						Type: "",
					},
				},
			},

			{
				Field: "quiz.maths.q2.options",
				Types: []typeData{
					{
						File: "a.json",
						Type: "array",
					},
					{
						File: "b.json",
						Type: "object",
					},
				},
			},
			{
				Field: "quiz.maths.q2.question",
				Types: []typeData{
					{
						File: "a.json",
						Type: "float",
					},
					{
						File: "b.json",
						Type: "string",
					},
				},
			},
			{
				Field: "quiz.science",
				Types: []typeData{
					{
						File: "a.json",
						Type: "",
					},
					{
						File: "b.json",
						Type: "object",
					},
				},
			},
			{
				Field: "quiz.sport.q1.answer",
				Types: []typeData{
					{
						File: "a.json",
						Type: "null",
					},
					{
						File: "b.json",
						Type: "string",
					},
				},
			},
			{
				Field: "quiz.sport.q1.options",
				Types: []typeData{
					{
						File: "a.json",
						Type: "array",
					},
					{
						File: "b.json",
						Type: "",
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Log(expected)
		t.Log(actual)
		t.Fatal("expected and actual did not match")
	}
}

func TestOutputFormatWriteStdout(t *testing.T) {
	opFormat := outputFormat{
		Field: "test",
		Types: []typeData{
			{
				File: "a.json",
				Type: "string",
			},
			{
				File: "b.json",
				Type: "integer",
			},
		},
	}

	m := &mockWriter{}
	err := opFormat.printStdout(m)
	if err != nil {
		t.Fatal(err)
	}
	expected := "test:\n\ta.json: string\n\tb.json: integer\n"
	if expected != m.data {
		t.Log(expected)
		t.Log(m.data)
		t.FailNow()
	}
}

func TestOutputFormatWriteStdoutFailed(t *testing.T) {
	opFormat := outputFormat{
		Field: "test",
		Types: []typeData{
			{
				File: "a.json",
				Type: "string",
			},
			{
				File: "b.json",
				Type: "integer",
			},
		},
	}

	m := &mockWriterError{msg: "TestOutputFormatWriteStdoutFailed"}
	err := opFormat.printStdout(m)
	if err == nil {
		t.Fatal("err is nil")
	}
	expected := "mock error: TestOutputFormatWriteStdoutFailed"
	if expected != err.Error() {
		t.Log(expected)
		t.Log(err.Error())
		t.FailNow()
	}
}

func TestOutputListWriteStdoutNoDiff(t *testing.T) {
	opList := outputList{outputs: []outputFormat{}}

	m := &mockWriter{}
	err := opList.printStdout(m)
	if err != nil {
		t.Fatal(err)
	}
	expected := "No Diff, all files are exactly the same!\n"
	if expected != m.data {
		t.Log(expected)
		t.Log(m.data)
		t.FailNow()
	}
}

func TestOutputListWriteStdoutFailed(t *testing.T) {
	opList := outputList{outputs: []outputFormat{}}

	m := &mockWriterError{msg: "TestOutputListWriteStdoutFailed"}
	err := opList.printStdout(m)
	if err == nil {
		t.Fatal("err is nil")
	}
	expected := "mock error: TestOutputListWriteStdoutFailed"
	if expected != err.Error() {
		t.Log(expected)
		t.Log(err.Error())
		t.FailNow()
	}
}

func TestOutputListWriteStdoutSomeDiff(t *testing.T) {
	opList := outputList{
		outputs: []outputFormat{
			{
				Field: "test.1",
				Types: []typeData{
					{
						File: "a.json",
						Type: "string",
					},
					{
						File: "b.json",
						Type: "integer",
					},
				},
			},
			{
				Field: "test.2",
				Types: []typeData{
					{
						File: "a.json",
						Type: "",
					},
					{
						File: "b.json",
						Type: "object",
					},
				},
			},
			{
				Field: "test.3",
				Types: []typeData{
					{
						File: "a.json",
						Type: "null",
					},
					{
						File: "b.json",
						Type: "",
					},
				},
			},
		},
	}

	m := &mockWriter{}
	err := opList.printStdout(m)
	if err != nil {
		t.Fatal(err)
	}
	expected := "3 differences found:\n\ntest.1:\n\ta.json: string\n\tb.json: integer\n\ntest.2:\n\ta.json: \n\tb.json: object\n\ntest.3:\n\ta.json: null\n\tb.json: \n"
	if expected != m.data {
		t.Log(expected)
		t.Log(m.data)
		t.FailNow()
	}
}

func TestOutputListWriteJSONNoDiff(t *testing.T) {
	opList := outputList{outputs: []outputFormat{}}

	m := &mockWriter{}
	opList.writeJSON(m, "")
	expected := "No Diff, all files are exactly the same!\n"
	if expected != m.data {
		t.Log(expected)
		t.Log(m.data)
		t.FailNow()
	}
}

// func TestOutputListWriteJSONSomeDiff(t *testing.T) {
// 	opList := outputList{
// 		outputs: []outputFormat{
// 			{
// 				Field: "test.1",
// 				Types: []typeData{
// 					{
// 						File: "a.json",
// 						Type: "string",
// 					},
// 					{
// 						File: "b.json",
// 						Type: "integer",
// 					},
// 				},
// 			},
// 			{
// 				Field: "test.2",
// 				Types: []typeData{
// 					{
// 						File: "a.json",
// 						Type: "",
// 					},
// 					{
// 						File: "b.json",
// 						Type: "object",
// 					},
// 				},
// 			},
// 			{
// 				Field: "test.3",
// 				Types: []typeData{
// 					{
// 						File: "a.json",
// 						Type: "null",
// 					},
// 					{
// 						File: "b.json",
// 						Type: "",
// 					},
// 				},
// 			},
// 		},
// 	}

// 	m := &mockWriter{}
// 	err := opList.writeJSON(m, "")
// 	expected := ""
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if expected != m.data {
// 		t.Log(expected)
// 		t.Log(m.data)
// 		t.FailNow()
// 	}
// }

func TestPrintDiffOutputWithSomeDiff(t *testing.T) {
	comparisonResult := schmp.ComparisonOutput{
		Diff: map[string][]string{
			"test.1": {
				"string",
				"integer",
			},
			"test.2": {
				"",
				"object",
			},
			"test.3": {
				"null",
				"",
			},
		},
	}

	m := &mockWriter{}
	printDiffOutput(comparisonResult, cliOptions{
		outType: "stdout",
		inFiles: []string{
			"a.json",
			"b.json",
		},
	}, m)
	expected := "3 differences found:\n\ntest.1:\n\ta.json: string\n\tb.json: integer\n\ntest.2:\n\ta.json: \n\tb.json: object\n\ntest.3:\n\ta.json: null\n\tb.json: \n"
	if expected != m.data {
		t.Log(expected)
		t.Log(m.data)
		t.FailNow()
	}
}

func TestPrintDiffOutputNoDiff(t *testing.T) {
	comparisonResult := schmp.ComparisonOutput{
		Diff: map[string][]string{},
	}

	m := &mockWriter{}
	printDiffOutput(comparisonResult, cliOptions{
		outType: "stdout",
	}, m)
	expected := "No Diff, all files are exactly the same!\n"
	if expected != m.data {
		t.Log(expected)
		t.Log(m.data)
		t.FailNow()
	}
}

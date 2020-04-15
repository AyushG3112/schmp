package main

import (
	"testing"

	"github.com/pkg/errors"
)

type mockWriter struct {
	data string
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.data = m.data + string(p)
	return len(p), nil
}

type mockWriterError struct {
	msg  string
	data string
}

func (m *mockWriterError) Write(p []byte) (n int, err error) {
	m.data = string(p)

	return 0, errors.New("mock error: " + m.msg)
}

type mockBytesWriter struct {
	data []byte
}

func (m *mockBytesWriter) Write(p []byte) (n int, err error) {
	if m.data == nil {
		m.data = []byte{}
	}
	m.data = append(m.data, p...)
	return len(p), nil
}

func TestPrintValidationErrors(t *testing.T) {
	mockWriter := &mockWriter{}
	err := printValidationErrors([]string{"validation error 1", "validation error 2"}, mockWriter)
	if err != nil {
		t.Fatal(err)
	}
	expected := "Could not process due to following errors: \n - validation error 1\n - validation error 2\n"
	if expected != mockWriter.data {
		t.FailNow()
	}
}

func TestPrintValidationErrorsDoesNotPrintAnythingIfNoErrors(t *testing.T) {
	m := &mockWriter{}
	err := printValidationErrors([]string{}, m)
	if err != nil {
		t.Fatal(err)
	}
	expected := ""
	if expected != m.data {
		t.FailNow()
	}
}

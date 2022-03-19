package testutils

import "testing"

func ExpectSlice[T comparable](t *testing.T, expected []T, actual []T) {
	t.Helper()

	if len(expected) != len(actual) {
		t.Errorf("Expected slice %v, got %v", expected, actual)
		return
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Expected slice %v, got %v", expected, actual)
			return
		}
	}
}

func ExpectMap[T comparable, V comparable](t *testing.T, expected map[T]V, actual map[T]V) {
	t.Helper()

	if len(expected) != len(actual) {
		t.Errorf("Expected map %v, got %v", expected, actual)
		return
	}
	for key := range expected {
		if expected[key] != actual[key] {
			t.Errorf("Expected map %v, got %v", expected, actual)
			return
		}
	}
}

func ExpectError(t *testing.T, err error, expected string) {
	t.Helper()

	if err == nil {
		t.Errorf("Expected error, got nil")
		return
	}
	actual := err.Error()
	if actual != expected {
		t.Errorf("Error message is %s, expected %s", actual, expected)
	}
}

func ExpectNilError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

func ExpectPanic(t *testing.T) {
	t.Helper()

	if r := recover(); r == nil {
		t.Errorf("The code did not panic")
	}
}

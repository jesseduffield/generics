package list

import "testing"

func expectSlice(t *testing.T, expected []int, actual []int) {
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

func expectError(t *testing.T, err error, expected string) {
	if err == nil {
		t.Errorf("Expected error, got nil")
		return
	}
	actual := err.Error()
	if actual != expected {
		t.Errorf("Error message is %s, expected %s", actual, expected)
	}
}

func expectNilError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

func cloneSlice(slice []int) []int {
	clone := make([]int, len(slice))
	copy(clone, slice)
	return clone
}

package hashmap

import "testing"

// duplicating this for now to spare the need for having a separate test utils package
func expectSlice[T comparable](t *testing.T, expected []T, actual []T) {
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

// duplicating this for now to spare the need for having a separate test utils package
func expectMap[T comparable, V comparable](t *testing.T, expected map[T]V, actual map[T]V) {
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

func TestKeys(t *testing.T) {
	tests := []struct {
		hashMap  map[string]int
		expected []string
	}{
		{map[string]int{}, []string{}},
		{map[string]int{"a": 1}, []string{"a"}},
	}
	for _, test := range tests {
		expectSlice(t, test.expected, Keys(test.hashMap))
	}
}

func TestValues(t *testing.T) {
	tests := []struct {
		hashMap  map[string]int
		expected []int
	}{
		{map[string]int{}, []int{}},
		{map[string]int{"a": 1}, []int{1}},
	}
	for _, test := range tests {
		expectSlice(t, test.expected, Values(test.hashMap))
	}
}

func TestTransformKeys(t *testing.T) {
	tests := []struct {
		hashMap   map[int]string
		transform func(int) int64
		expected  map[int64]string
	}{
		{
			hashMap:   map[int]string{},
			transform: func(i int) int64 { return (int64)(i) },
			expected:  map[int64]string{},
		},
		{
			hashMap:   map[int]string{1: "a"},
			transform: func(i int) int64 { return 2 * (int64)(i) },
			expected:  map[int64]string{2: "a"},
		},
	}
	for _, test := range tests {
		expectMap(t, test.expected, TransformKeys(test.hashMap, test.transform))
	}
}

func TestTransformValues(t *testing.T) {
	tests := []struct {
		hashMap   map[string]int
		transform func(int) int64
		expected  map[string]int64
	}{
		{
			hashMap:   map[string]int{},
			transform: func(i int) int64 { return (int64)(i) },
			expected:  map[string]int64{},
		},
		{
			hashMap:   map[string]int{"a": 1},
			transform: func(i int) int64 { return 2 * (int64)(i) },
			expected:  map[string]int64{"a": 2},
		},
	}
	for _, test := range tests {
		expectMap(t, test.expected, TransformValues(test.hashMap, test.transform))
	}
}

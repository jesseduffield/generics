package list

import "testing"

func TestEqual(t *testing.T) {
	tests := []struct {
		first    []int
		second   []int
		expected bool
	}{
		{[]int{}, []int{}, true},
		{[]int{1}, []int{1}, true},
		{[]int{}, []int{1}, false},
		{[]int{1, 2}, []int{1, 2}, true},
		{[]int{1, 2}, []int{2, 1}, false},
		{[]int{1, 2, 3}, []int{1, 2}, false},
	}
	for _, test := range tests {
		first := NewComparableFromSlice(test.first)
		second := NewComparableFromSlice(test.second)
		if first.Equal(second) != test.expected {
			t.Errorf("Equal(%v, %v) = %v, expected %v",
				test.first, test.second, first.Equal(second), test.expected,
			)
		}
	}
}

func TestCompact(t *testing.T) {
	tests := []struct {
		slice    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{1, 2, 1}, []int{1, 2, 1}},
		{[]int{1, 1, 1}, []int{1}},
	}
	for _, test := range tests {
		list := NewComparableFromSlice(test.slice)
		expectSlice(t, test.expected, list.Compact().ToSlice())
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		slice    []int
		value    int
		expected int
	}{
		{[]int{}, 1, -1},
		{[]int{1}, 1, 0},
		{[]int{1, 1}, 1, 0},
		{[]int{1, 1}, 2, -1},
		{[]int{1, 2, 3}, 2, 1},
		{[]int{1, 2, 3}, 3, 2},
	}
	for _, test := range tests {
		list := NewComparableFromSlice(test.slice)

		if list.IndexOf(test.value) != test.expected {
			t.Errorf("Index(%v, %v) = %v, expected %v", test.slice, test.value, list.IndexOf(test.value), test.expected)
		}
	}
}

func TestIndexFunc(t *testing.T) {
	tests := []struct {
		slice    []int
		f        func(value int) bool
		expected int
	}{
		{[]int{}, func(value int) bool { return true }, -1},
		{[]int{1}, func(value int) bool { return true }, 0},
		{[]int{1, 1}, func(value int) bool { return true }, 0},
		{[]int{1, 1}, func(value int) bool { return false }, -1},
		{[]int{1, 2, 3}, func(value int) bool { return value == 2 }, 1},
		{[]int{1, 2, 3}, func(value int) bool { return value == 3 }, 2},
	}
	for _, test := range tests {
		list := NewComparableFromSlice(test.slice)

		if list.IndexFunc(test.f) != test.expected {
			t.Errorf("IndexFunc(%v, func) = %v, expected %v", test.slice, list.IndexFunc(test.f), test.expected)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		slice    []int
		value    int
		expected bool
	}{
		{[]int{}, 1, false},
		{[]int{1}, 1, true},
		{[]int{1, 1}, 1, true},
		{[]int{1, 1}, 2, false},
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 3, true},
	}
	for _, test := range tests {
		list := NewComparableFromSlice(test.slice)

		if list.Contains(test.value) != test.expected {
			t.Errorf("Contains(%v, %v) = %v, expected %v", test.slice, test.value, list.Contains(test.value), test.expected)
		}
	}
}

func TestContainsFunc(t *testing.T) {
	tests := []struct {
		slice    []int
		f        func(value int) bool
		expected bool
	}{
		{[]int{}, func(value int) bool { return true }, false},
		{[]int{1}, func(value int) bool { return true }, true},
		{[]int{1, 1}, func(value int) bool { return true }, true},
		{[]int{1, 1}, func(value int) bool { return false }, false},
		{[]int{1, 2, 3}, func(value int) bool { return value == 2 }, true},
		{[]int{1, 2, 3}, func(value int) bool { return value == 3 }, true},
	}
	for _, test := range tests {
		list := NewComparableFromSlice(test.slice)

		if list.ContainsFunc(test.f) != test.expected {
			t.Errorf("ContainsFunc(%v, func) = %v, expected %v", test.slice, list.ContainsFunc(test.f), test.expected)
		}
	}
}

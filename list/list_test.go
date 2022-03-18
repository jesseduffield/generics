package list

import (
	"testing"
)

func TestPush(t *testing.T) {
	list := NewList[int]()
	list.Push(1)
	list.Push(2)
	slice := list.ToSlice()
	expectSlice(t, []int{1, 2}, slice)
}

func TestInsert(t *testing.T) {
	tests := []struct {
		startSlice []int
		index      int
		value      int
		endSlice   []int
	}{
		{[]int{}, 0, 1, []int{1}},
		{[]int{1}, 0, 2, []int{2, 1}},
		{[]int{1, 2}, 1, 3, []int{1, 3, 2}},
		{[]int{1, 2}, 2, 3, []int{1, 2, 3}},
	}
	for _, test := range tests {
		list := NewListFromSlice(test.startSlice)
		err := list.Insert(test.index, test.value)
		expectNilError(t, err)
		expectSlice(t, test.endSlice, list.ToSlice())
	}

	errorTests := []struct {
		startSlice         []int
		index              int
		value              int
		expectedErrMessage string
	}{
		{[]int{}, 1, 1, "Cannot insert: index 1 out of bounds for list of length 0"},
		{[]int{}, 2, 1, "Cannot insert: index 2 out of bounds for list of length 0"},
		{[]int{1}, 2, 1, "Cannot insert: index 2 out of bounds for list of length 1"},
		{[]int{1}, -1, 1, "Cannot insert: index -1 is negative"},
	}
	for _, test := range errorTests {
		t.Run("", func(t *testing.T) {
			originalSlice := cloneSlice(test.startSlice)
			list := NewListFromSlice(test.startSlice)
			err := list.Insert(test.index, test.value)
			expectError(t, err, test.expectedErrMessage)
			expectSlice(t, originalSlice, list.ToSlice())
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		startSlice []int
		index      int
		endSlice   []int
	}{
		{[]int{1}, 0, []int{}},
		{[]int{1, 2}, 0, []int{2}},
		{[]int{1, 2}, 1, []int{1}},
	}
	for _, test := range tests {
		list := NewListFromSlice(test.startSlice)
		err := list.Remove(test.index)
		expectSlice(t, test.endSlice, list.ToSlice())
		expectNilError(t, err)
	}

	errorTests := []struct {
		startSlice         []int
		index              int
		expectedErrMessage string
	}{
		{[]int{}, 0, "Cannot remove: index 0 out of bounds for list of length 0"},
		{[]int{}, 1, "Cannot remove: index 1 out of bounds for list of length 0"},
		{[]int{1}, 1, "Cannot remove: index 1 out of bounds for list of length 1"},
		{[]int{1}, -1, "Cannot remove: index -1 is negative"},
	}
	for _, test := range errorTests {
		t.Run("", func(t *testing.T) {
			originalSlice := cloneSlice(test.startSlice)
			list := NewListFromSlice(test.startSlice)
			err := list.Remove(test.index)
			expectError(t, err, test.expectedErrMessage)
			expectSlice(t, originalSlice, list.ToSlice())
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		startSlice []int
		endSlice   []int
	}{
		{[]int{1}, []int{}},
		{[]int{1, 2}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2}},
	}
	for _, test := range tests {
		list := NewListFromSlice(test.startSlice)
		list.Pop()
		expectSlice(t, test.endSlice, list.ToSlice())
	}

	errorTests := []struct {
		startSlice         []int
		expectedErrMessage string
	}{
		{[]int{}, "Cannot pop: list is empty"},
	}
	for _, test := range errorTests {
		t.Run("", func(t *testing.T) {
			originalSlice := cloneSlice(test.startSlice)
			list := NewListFromSlice(test.startSlice)
			err := list.Pop()
			expectError(t, err, test.expectedErrMessage)
			expectSlice(t, originalSlice, list.ToSlice())
		})
	}
}

func TestFilter(t *testing.T) {
	even := func(value int) bool { return value%2 == 0 }
	tests := []struct {
		startSlice []int
		testFunc   func(value int) bool
		endSlice   []int
	}{
		{[]int{1}, even, []int{}},
		{[]int{1, 2}, even, []int{2}},
		{[]int{1, 2, 3}, even, []int{2}},
		{[]int{1, 2, 3, 4}, even, []int{2, 4}},
	}
	for _, test := range tests {
		originalSlice := cloneSlice(test.startSlice)
		list := NewListFromSlice(test.startSlice)
		expectSlice(t, test.endSlice, list.Filter(test.testFunc).ToSlice())
		expectSlice(t, originalSlice, list.ToSlice())
	}
}

func TestFilterInPlace(t *testing.T) {
	even := func(value int) bool { return value%2 == 0 }
	tests := []struct {
		startSlice []int
		testFunc   func(value int) bool
		endSlice   []int
	}{
		{[]int{1}, even, []int{}},
		{[]int{1, 2}, even, []int{2}},
		{[]int{1, 2, 3}, even, []int{2}},
		{[]int{1, 2, 3, 4}, even, []int{2, 4}},
	}

	for _, test := range tests {
		list := NewListFromSlice(test.startSlice)
		list.FilterInPlace(test.testFunc)
		expectSlice(t, test.endSlice, list.ToSlice())
	}
}

func TestMap(t *testing.T) {
	double := func(value int) int { return value * 2 }
	tests := []struct {
		startSlice []int
		mapFunc    func(value int) int
		endSlice   []int
	}{
		{[]int{}, double, []int{}},
		{[]int{1}, double, []int{2}},
		{[]int{1, 2}, double, []int{2, 4}},
		{[]int{1, 2, 3}, double, []int{2, 4, 6}},
		{[]int{1, 2, 3, 4}, double, []int{2, 4, 6, 8}},
	}

	for _, test := range tests {
		originalSlice := cloneSlice(test.startSlice)
		list := NewListFromSlice(test.startSlice)
		expectSlice(t, test.endSlice, list.Map(test.mapFunc).ToSlice())
		expectSlice(t, originalSlice, list.ToSlice())
	}
}

func TestMapInPlace(t *testing.T) {
	double := func(value int) int { return value * 2 }
	tests := []struct {
		startSlice []int
		mapFunc    func(value int) int
		endSlice   []int
	}{
		{[]int{}, double, []int{}},
		{[]int{1}, double, []int{2}},
		{[]int{1, 2}, double, []int{2, 4}},
		{[]int{1, 2, 3}, double, []int{2, 4, 6}},
		{[]int{1, 2, 3, 4}, double, []int{2, 4, 6, 8}},
	}
	for _, test := range tests {
		list := NewListFromSlice(test.startSlice)
		list.MapInPlace(test.mapFunc)
		expectSlice(t, test.endSlice, list.ToSlice())
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		startSlice []int
		other      *List[int]
		endSlice   []int
	}{
		{[]int{}, NewListFromSlice([]int{}), []int{}},
		{[]int{1}, NewListFromSlice([]int{2}), []int{1, 2}},
		{[]int{1, 2}, NewListFromSlice([]int{3, 4}), []int{1, 2, 3, 4}},
	}
	for _, test := range tests {
		list := NewListFromSlice(test.startSlice)
		list.Concat(test.other)
		expectSlice(t, test.endSlice, list.ToSlice())
	}
}

func TestConcatSlice(t *testing.T) {
	tests := []struct {
		startSlice []int
		other      []int
		endSlice   []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{2}, []int{1, 2}},
		{[]int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
	}
	for _, test := range tests {
		list := NewListFromSlice(test.startSlice)
		list.ConcatSlice(test.other)
		expectSlice(t, test.endSlice, list.ToSlice())
	}
}

func TestSome(t *testing.T) {
	even := func(value int) bool { return value%2 == 0 }
	tests := []struct {
		startSlice []int
		testFunc   func(value int) bool
		expected   bool
	}{
		{[]int{}, even, false},
		{[]int{1}, even, false},
		{[]int{2}, even, true},
		{[]int{1, 2}, even, true},
	}
	for _, test := range tests {
		list := NewListFromSlice(test.startSlice)
		if list.Some(test.testFunc) != test.expected {
			t.Errorf("Some(%v) = %v, expected %v",
				test.startSlice, list.Some(test.testFunc), test.expected,
			)
		}
	}
}

func TestEvery(t *testing.T) {
	even := func(value int) bool { return value%2 == 0 }
	tests := []struct {
		startSlice []int
		testFunc   func(value int) bool
		expected   bool
	}{
		{[]int{}, even, true},
		{[]int{1}, even, false},
		{[]int{2}, even, true},
		{[]int{1, 2}, even, false},
		{[]int{2, 2}, even, true},
	}
	for _, test := range tests {
		list := NewListFromSlice(test.startSlice)
		if list.Every(test.testFunc) != test.expected {
			t.Errorf("Every(%v) = %v, expected %v",
				test.startSlice, list.Some(test.testFunc), test.expected,
			)
		}
	}
}

func TestClone(t *testing.T) {
	list := NewListFromSlice([]int{1, 2, 3})
	clone := list.Clone()
	expectSlice(t, []int{1, 2, 3}, clone.ToSlice())

	// ensure that the clone has its own slice
	clone.Insert(0, 1)
	expectSlice(t, []int{1, 1, 2, 3}, clone.ToSlice())
	expectSlice(t, []int{1, 2, 3}, list.ToSlice())
}

func TestFilterMap(t *testing.T) {
	list := NewListFromSlice([]int{1, 2, 3, 4})
	result := list.
		Filter(func(value int) bool { return value%2 == 0 }).
		Map(func(value int) int { return value * 2 })

	expectSlice(t, []int{4, 8}, result.ToSlice())
}

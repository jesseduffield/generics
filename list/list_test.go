package list

import (
	"testing"
)

func TestPush(t *testing.T) {
	list := New[int]()
	list.Push(1)
	list.Push(2)
	slice := list.ToSlice()
	expectSlice(t, []int{1, 2}, slice)
}

func TestInsert(t *testing.T) {
	tests := []struct {
		startSlice []int
		index      int
		values     []int
		endSlice   []int
	}{
		{[]int{}, 0, []int{1}, []int{1}},
		{[]int{1}, 0, []int{2}, []int{2, 1}},
		{[]int{1, 2}, 1, []int{3}, []int{1, 3, 2}},
		{[]int{1, 2}, 2, []int{3}, []int{1, 2, 3}},
		{[]int{1, 2}, 2, []int{3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 2}, 1, []int{3, 4}, []int{1, 3, 4, 2}},
	}
	for _, test := range tests {
		list := NewFromSlice(test.startSlice)
		list.Insert(test.index, test.values...)
		expectSlice(t, test.endSlice, list.ToSlice())
	}

	panicTests := []struct {
		startSlice []int
		index      int
		value      int
	}{
		{[]int{}, 1, 1},
		{[]int{}, 2, 1},
		{[]int{1}, 2, 1},
		{[]int{1}, -1, 1},
	}
	for _, test := range panicTests {
		t.Run("", func(t *testing.T) {
			defer expectPanic(t)
			list := NewFromSlice(test.startSlice)
			list.Insert(test.index, test.value)
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
		list := NewFromSlice(test.startSlice)
		list.Remove(test.index)
		expectSlice(t, test.endSlice, list.ToSlice())
	}

	panicTests := []struct {
		startSlice []int
		index      int
	}{
		{[]int{}, 0},
		{[]int{}, 1},
		{[]int{1}, 1},
		{[]int{1}, -1},
	}
	for _, test := range panicTests {
		t.Run("", func(t *testing.T) {
			defer expectPanic(t)
			list := NewFromSlice(test.startSlice)
			list.Remove(test.index)
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
		list := NewFromSlice(test.startSlice)
		list.Pop()
		expectSlice(t, test.endSlice, list.ToSlice())
	}

	panicTests := []struct {
		startSlice []int
	}{
		{[]int{}},
	}
	for _, test := range panicTests {
		t.Run("", func(t *testing.T) {
			defer expectPanic(t)
			list := NewFromSlice(test.startSlice)
			list.Pop()
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
		list := NewFromSlice(test.startSlice)
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
		list := NewFromSlice(test.startSlice)
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
		list := NewFromSlice(test.startSlice)
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
		list := NewFromSlice(test.startSlice)
		list.MapInPlace(test.mapFunc)
		expectSlice(t, test.endSlice, list.ToSlice())
	}
}

func TestAppend(t *testing.T) {
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
		list := NewFromSlice(test.startSlice)
		list.Append(test.other...)
		expectSlice(t, test.endSlice, list.ToSlice())
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		startSlice []int
		other      []int
		endSlice   []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{2}, []int{2, 1}},
		{[]int{1, 2}, []int{3, 4}, []int{3, 4, 1, 2}},
	}
	for _, test := range tests {
		list := NewFromSlice(test.startSlice)
		list.Prepend(test.other...)
		expectSlice(t, test.endSlice, list.ToSlice())
	}
}

func TestConcat(t *testing.T) {
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
		originalSlice := cloneSlice(test.startSlice)
		list := NewFromSlice(test.startSlice)
		result := list.Concat(test.other...)
		expectSlice(t, test.endSlice, result.ToSlice())
		expectSlice(t, originalSlice, list.ToSlice())
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
		list := NewFromSlice(test.startSlice)
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
		list := NewFromSlice(test.startSlice)
		if list.Every(test.testFunc) != test.expected {
			t.Errorf("Every(%v) = %v, expected %v",
				test.startSlice, list.Every(test.testFunc), test.expected,
			)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		startSlice []int
		expected   []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}
	for _, test := range tests {
		list := NewFromSlice(test.startSlice)
		expectSlice(t, test.expected, list.Reverse().ToSlice())
	}
}

func TestReverseInPlace(t *testing.T) {
	tests := []struct {
		startSlice []int
		expected   []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}
	for _, test := range tests {
		list := NewFromSlice(test.startSlice)
		list.ReverseInPlace()
		expectSlice(t, test.expected, list.ToSlice())
	}
}

func TestClone(t *testing.T) {
	list := NewFromSlice([]int{1, 2, 3})
	clone := list.Clone()
	expectSlice(t, []int{1, 2, 3}, clone.ToSlice())

	// ensure that the clone has its own slice
	clone.Insert(0, 1)
	expectSlice(t, []int{1, 1, 2, 3}, clone.ToSlice())
	expectSlice(t, []int{1, 2, 3}, list.ToSlice())
}

func TestFilterMap(t *testing.T) {
	list := NewFromSlice([]int{1, 2, 3, 4})
	result := list.
		Filter(func(value int) bool { return value%2 == 0 }).
		Map(func(value int) int { return value * 2 })

	expectSlice(t, []int{4, 8}, result.ToSlice())
}

func TestIsEmpty(t *testing.T) {
	list := NewFromSlice([]int{})
	if !list.IsEmpty() {
		t.Errorf("IsEmpty() = %v, expected %v", list.IsEmpty(), true)
	}

	list = NewFromSlice([]int{1})
	if list.IsEmpty() {
		t.Errorf("IsEmpty() = %v, expected %v", list.IsEmpty(), false)
	}
}

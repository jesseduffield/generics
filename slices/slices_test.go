package slices

import (
	"strconv"
	"testing"

	"github.com/jesseduffield/generics/internal/testutils"
	"golang.org/x/exp/slices"
)

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
		result := Insert(test.startSlice, test.index, test.values...)
		testutils.ExpectSlice(t, test.endSlice, result)
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
		func() {
			defer testutils.ExpectPanic(t)
			Insert(test.startSlice, test.index, test.value)
		}()
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
		result := Remove(test.startSlice, test.index)
		testutils.ExpectSlice(t, test.endSlice, result)
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
		func() {
			defer testutils.ExpectPanic(t)
			Remove(test.startSlice, test.index)
		}()
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		startSlice []int
		value      int
		endSlice   []int
	}{
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 2, []int{1}},
		{[]int{1, 2, 3}, 3, []int{1, 2}},
	}
	for _, test := range tests {
		value, slice := Pop(test.startSlice)
		if value != test.value {
			t.Errorf("expected %d, got %d", test.value, value)
		}
		testutils.ExpectSlice(t, test.endSlice, slice)
	}

	panicTests := []struct {
		startSlice []int
	}{
		{[]int{}},
	}
	for _, test := range panicTests {
		func() {
			defer testutils.ExpectPanic(t)
			Pop(test.startSlice)
		}()
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
		testSlice := slices.Clone(test.startSlice)
		testutils.ExpectSlice(t, test.endSlice, Filter(testSlice, test.testFunc))
		testutils.ExpectSlice(t, testSlice, test.startSlice)
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
		result := FilterInPlace(test.startSlice, test.testFunc)
		testutils.ExpectSlice(t, test.endSlice, result)
	}
}

func TestFilterMap(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	result := FilterMap(slice,
		func(value int) (string, bool) {
			if value%2 != 0 {
				return "", false
			}

			return strconv.Itoa(value * 2), true
		},
	)

	testutils.ExpectSlice(t, []string{"4", "8"}, result)
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
		testSlice := slices.Clone(test.startSlice)
		testutils.ExpectSlice(t, test.endSlice, Map(testSlice, test.mapFunc))
		testutils.ExpectSlice(t, test.startSlice, testSlice)
	}
}

func TestMapWithIndex(t *testing.T) {
	double := func(value int, i int) int { return value*2 + i }
	tests := []struct {
		startSlice []int
		mapFunc    func(value int, i int) int
		endSlice   []int
	}{
		{[]int{}, double, []int{}},
		{[]int{1}, double, []int{2}},
		{[]int{1, 2}, double, []int{2, 5}},
		{[]int{1, 2, 3}, double, []int{2, 5, 8}},
		{[]int{1, 2, 3, 4}, double, []int{2, 5, 8, 11}},
	}

	for _, test := range tests {
		testSlice := slices.Clone(test.startSlice)
		testutils.ExpectSlice(t, test.endSlice, MapWithIndex(testSlice, test.mapFunc))
		testutils.ExpectSlice(t, test.startSlice, testSlice)
	}
}

func TestFilterWithIndex(t *testing.T) {
	double := func(value int, i int) bool { return (value*2+i)%2 == 0 }
	tests := []struct {
		startSlice []int
		mapFunc    func(value int, i int) bool
		endSlice   []int
	}{
		{[]int{}, double, []int{}},
		{[]int{1}, double, []int{1}},
		{[]int{1, 2}, double, []int{1}},
		{[]int{1, 2, 3}, double, []int{1, 3}},
		{[]int{1, 2, 3, 4}, double, []int{1, 3}},
	}

	for _, test := range tests {
		testSlice := slices.Clone(test.startSlice)
		testutils.ExpectSlice(t, test.endSlice, FilterWithIndex(testSlice, test.mapFunc))
		testutils.ExpectSlice(t, test.startSlice, testSlice)
	}
}

func TestFlatMap(t *testing.T) {
	f := func(value int) []int { return []int{value * 2, value * 4} }
	tests := []struct {
		startSlice []int
		mapFunc    func(value int) []int
		endSlice   []int
	}{
		{[]int{}, f, []int{}},
		{[]int{1}, f, []int{2, 4}},
		{[]int{1, 2}, f, []int{2, 4, 4, 8}},
		{[]int{1, 2, 3}, f, []int{2, 4, 4, 8, 6, 12}},
		{[]int{1, 2, 3, 4}, f, []int{2, 4, 4, 8, 6, 12, 8, 16}},
	}

	for _, test := range tests {
		testSlice := slices.Clone(test.startSlice)
		testutils.ExpectSlice(t, test.endSlice, FlatMap(testSlice, test.mapFunc))
		testutils.ExpectSlice(t, test.startSlice, testSlice)
	}
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		startSlice [][]int
		endSlice   []int
	}{
		{[][]int{}, []int{}},
		{[][]int{{1}}, []int{1}},
		{[][]int{{1, 2}, {3}}, []int{1, 2, 3}},
		{[][]int{{1, 2}, {}, {3, 4}}, []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		testutils.ExpectSlice(t, test.endSlice, Flatten(test.startSlice))
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
		testSlice := slices.Clone(test.startSlice)
		MapInPlace(testSlice, test.mapFunc)
		testutils.ExpectSlice(t, test.endSlice, testSlice)
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		slice         []int
		values        []int
		expectedSlice []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{1}, []int{1}},
		{[]int{1}, []int{2}, []int{2, 1}},
		{[]int{1, 2}, []int{3}, []int{3, 1, 2}},
		{[]int{1, 2}, []int{3, 4}, []int{3, 4, 1, 2}},
	}
	for _, test := range tests {
		slice := Prepend(test.slice, test.values...)
		testutils.ExpectSlice(t, test.expectedSlice, slice)
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		startSlice []int
		values     []int
		endSlice   []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{2}, []int{1, 2}},
		{[]int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
	}
	for _, test := range tests {
		testSlice := slices.Clone(test.startSlice)
		result := Concat(testSlice, test.values...)
		testutils.ExpectSlice(t, test.endSlice, result)
		testutils.ExpectSlice(t, test.startSlice, testSlice)
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
		if Some(test.startSlice, test.testFunc) != test.expected {
			t.Errorf("Some(%v) = %v, expected %v",
				test.startSlice, Some(test.startSlice, test.testFunc), test.expected,
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
		if Every(test.startSlice, test.testFunc) != test.expected {
			t.Errorf("Every(%v) = %v, expected %v",
				test.startSlice, Every(test.startSlice, test.testFunc), test.expected,
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
		testSlice := slices.Clone(test.startSlice)
		result := Reverse(testSlice)
		testutils.ExpectSlice(t, test.expected, result)
		testutils.ExpectSlice(t, test.startSlice, testSlice)
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
		testSlice := slices.Clone(test.startSlice)
		ReverseInPlace(testSlice)
		testutils.ExpectSlice(t, test.expected, testSlice)
	}
}

func TestClone(t *testing.T) {
	slice := []int{1, 2, 3}
	clone := Clone(slice)
	testutils.ExpectSlice(t, []int{1, 2, 3}, clone)

	// ensure that the clone has its own slice
	clone = Prepend(clone, 0)
	testutils.ExpectSlice(t, []int{0, 1, 2, 3}, clone)
	testutils.ExpectSlice(t, []int{1, 2, 3}, slice)
}

func TestMove(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	slice = Move(slice, 2, 1)
	testutils.ExpectSlice(t, []int{1, 3, 2, 4}, slice)
}

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
		if Equal(test.first, test.second) != test.expected {
			t.Errorf("Equal(%v, %v) = %v, expected %v",
				test.first, test.second, Equal(test.first, test.second), test.expected,
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
		result := Compact(test.slice)
		testutils.ExpectSlice(t, test.expected, result)
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
		if Index(test.slice, test.value) != test.expected {
			t.Errorf("Index(%v, %v) = %v, expected %v", test.slice, test.value, Index(test.slice, test.value), test.expected)
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
		if IndexFunc(test.slice, test.f) != test.expected {
			t.Errorf("IndexFunc(%v, func) = %v, expected %v", test.slice, IndexFunc(test.slice, test.f), test.expected)
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
		if Contains(test.slice, test.value) != test.expected {
			t.Errorf("Contains(%v, %v) = %v, expected %v", test.slice, test.value, Contains(test.slice, test.value), test.expected)
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
		if ContainsFunc(test.slice, test.f) != test.expected {
			t.Errorf("ContainsFunc(%v, func) = %v, expected %v", test.slice, ContainsFunc(test.slice, test.f), test.expected)
		}
	}
}

func TestShift(t *testing.T) {
	tests := []struct {
		slice         []int
		expectedValue int
		expectedSlice []int
	}{
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{2}},
	}
	for _, test := range tests {
		value, slice := Shift(test.slice)
		if value != test.expectedValue {
			t.Errorf("Shift(%v) = %v, expected %v", test.slice, value, test.expectedValue)
		}
		testutils.ExpectSlice(t, test.expectedSlice, slice)
	}
}

func TestPartition(t *testing.T) {
	even := func(value int) bool { return value%2 == 0 }
	tests := []struct {
		startSlice    []int
		testFunc      func(value int) bool
		endSliceLeft  []int
		endSliceRight []int
	}{
		{[]int{1}, even, []int{}, []int{1}},
		{[]int{1, 2}, even, []int{2}, []int{1}},
		{[]int{1, 2, 3}, even, []int{2}, []int{1, 3}},
		{[]int{1, 2, 3, 4}, even, []int{2, 4}, []int{1, 3}},
	}
	for _, test := range tests {
		testSlice := slices.Clone(test.startSlice)
		left, right := Partition(testSlice, test.testFunc)
		testutils.ExpectSlice(t, test.endSliceLeft, left)
		testutils.ExpectSlice(t, test.endSliceRight, right)
		testutils.ExpectSlice(t, testSlice, test.startSlice)
	}
}

func TestMaxBy(t *testing.T) {
	tests := []struct {
		slice    []int
		f        func(value int) int
		expected int
	}{
		{[]int{}, func(value int) int { return value }, 0},
		{[]int{1}, func(value int) int { return value }, 1},
		{[]int{-1}, func(value int) int { return value }, -1},
		{[]int{1, 2}, func(value int) int { return value }, 2},
		{[]int{3, 1, 2, 3}, func(value int) int { return value }, 3},
		{[]int{1, 2, 2, 3, 4}, func(value int) int { return value }, 4},
	}
	for _, test := range tests {
		if MaxBy(test.slice, test.f) != test.expected {
			t.Errorf("MaxBy(%v, func) = %v, expected %v", test.slice, MaxBy(test.slice, test.f), test.expected)
		}
	}
}

func TestMinBy(t *testing.T) {
	tests := []struct {
		slice    []int
		f        func(value int) int
		expected int
	}{
		{[]int{}, func(value int) int { return value }, 0},
		{[]int{1}, func(value int) int { return value }, 1},
		{[]int{-1}, func(value int) int { return value }, -1},
		{[]int{3, 1, 2}, func(value int) int { return value }, 1},
		{[]int{1, 2, 3}, func(value int) int { return value }, 1},
		{[]int{1, 2, 3, 4}, func(value int) int { return value }, 1},
	}
	for _, test := range tests {
		if MinBy(test.slice, test.f) != test.expected {
			t.Errorf("MinBy(%v, func) = %v, expected %v", test.slice, MinBy(test.slice, test.f), test.expected)
		}
	}
}

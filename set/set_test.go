package set

import "testing"

func TestAddIncludes(t *testing.T) {
	set := New[int]()
	set.Add(1)
	if !set.Includes(1) {
		t.Errorf("Add(1) failed: Includes(1) returned false")
	}

	if set.Includes(0) {
		t.Errorf("Add(1) failed: Includes(0) returned true")
	}
}

func TestAddSliceIncludes(t *testing.T) {
	set := New[int]()
	set.AddSlice([]int{1, 2})
	if !set.Includes(1) {
		t.Errorf("AddSlice failed: Includes(1) returned false")
	}

	if !set.Includes(2) {
		t.Errorf("AddSlice failed: Includes(2) returned false")
	}

	if set.Includes(3) {
		t.Errorf("AddSlice failed: Includes(3) returned true")
	}
}

func TestRemoveIncludes(t *testing.T) {
	set := NewFromSlice([]int{1, 2})
	set.Remove(1)
	if set.Includes(1) {
		t.Errorf("Remove failed: Includes(1) returned true")
	}

	if !set.Includes(2) {
		t.Errorf("Remove failed: Includes(2) returned false")
	}
}

func TestRemoveSliceIncludes(t *testing.T) {
	set := NewFromSlice([]int{1, 2, 3})
	set.RemoveSlice([]int{1, 2})
	if set.Includes(1) {
		t.Errorf("Remove failed: Includes(1) returned true")
	}

	if set.Includes(2) {
		t.Errorf("Remove failed: Includes(2) returned true")
	}

	if !set.Includes(3) {
		t.Errorf("Remove failed: Includes(3) returned false")
	}
}

func TestNewFromSlice(t *testing.T) {
	set := NewFromSlice([]int{1, 2})
	if !set.Includes(1) {
		t.Errorf("NewFromSlice failed: Includes(1) returned false")
	}

	if !set.Includes(2) {
		t.Errorf("NewFromSlice failed: Includes(2) returned false")
	}

	if set.Includes(3) {
		t.Errorf("NewFromSlice failed: Includes(3) returned true")
	}
}

func TestToSlice(t *testing.T) {
	set := New[int]()
	set.Add(1)
	set.Add(2)

	slice := set.ToSlice()
	if len(slice) != 2 {
		t.Errorf("ToSlice failed: len(slice) = %d, expected 2", len(slice))
	}

	if !((slice[0] == 1 && slice[1] == 2) || (slice[0] == 2 && slice[1] == 1)) {
		t.Errorf("ToSlice failed: expected 1 and 2 in slice in any order, got %v", slice)
	}
}

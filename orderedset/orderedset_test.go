package orderedset

import (
	"testing"

	"github.com/jesseduffield/generics/internal/testutils"
)

func TestAddIncludes(t *testing.T) {
	os := New[int]()
	os.Add(1)
	if !os.Includes(1) {
		t.Errorf("Add(1) failed: Includes(1) returned false")
	}

	if os.Includes(0) {
		t.Errorf("Add(1) failed: Includes(0) returned true")
	}
}

func TestAddSliceIncludes(t *testing.T) {
	os := New[int]()
	os.Add(1, 2)
	if !os.Includes(1) {
		t.Errorf("AddSlice failed: Includes(1) returned false")
	}

	if !os.Includes(2) {
		t.Errorf("AddSlice failed: Includes(2) returned false")
	}

	if os.Includes(3) {
		t.Errorf("AddSlice failed: Includes(3) returned true")
	}
}

func TestRemoveIncludes(t *testing.T) {
	os := NewFromSlice([]int{1, 2})
	os.Remove(1)
	if os.Includes(1) {
		t.Errorf("Remove failed: Includes(1) returned true")
	}

	if !os.Includes(2) {
		t.Errorf("Remove failed: Includes(2) returned false")
	}
}

func TestRemoveSliceIncludes(t *testing.T) {
	os := NewFromSlice([]int{1, 2, 3})
	os.RemoveSlice([]int{1, 2})
	if os.Includes(1) {
		t.Errorf("Remove failed: Includes(1) returned true")
	}

	if os.Includes(2) {
		t.Errorf("Remove failed: Includes(2) returned true")
	}

	if !os.Includes(3) {
		t.Errorf("Remove failed: Includes(3) returned false")
	}
}

func TestNewFromSlice(t *testing.T) {
	os := NewFromSlice([]int{1, 2})
	if !os.Includes(1) {
		t.Errorf("NewFromSlice failed: Includes(1) returned false")
	}

	if !os.Includes(2) {
		t.Errorf("NewFromSlice failed: Includes(2) returned false")
	}

	if os.Includes(3) {
		t.Errorf("NewFromSlice failed: Includes(3) returned true")
	}
}

func TestLen(t *testing.T) {
	set := NewFromSlice([]int{})
	if set.Len() != 0 {
		t.Errorf("Len() = %v, expected %v", set.Len(), 0)
	}

	set = NewFromSlice([]int{1})
	if set.Len() != 1 {
		t.Errorf("Len() = %v, expected %v", set.Len(), 1)
	}
}

func TestToSlice(t *testing.T) {
	set := New[int]()
	set.Add(1)
	set.Add(3)
	set.Add(2)

	slice := set.ToSliceFromOldest()
	testutils.ExpectSlice(t, []int{1, 3, 2}, slice)

	slice = set.ToSliceFromNewest()
	testutils.ExpectSlice(t, []int{2, 3, 1}, slice)
}

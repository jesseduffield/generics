package list

import (
	"github.com/jesseduffield/generics/slices"
)

type ComparableList[T comparable] struct {
	*List[T]
}

func NewComparable[T comparable]() *ComparableList[T] {
	return &ComparableList[T]{List: New[T]()}
}

func NewComparableFromSlice[T comparable](slice []T) *ComparableList[T] {
	return &ComparableList[T]{List: NewFromSlice(slice)}
}

func (l *ComparableList[T]) Equal(other *ComparableList[T]) bool {
	return l.EqualSlice(other.ToSlice())
}

func (l *ComparableList[T]) EqualSlice(other []T) bool {
	return slices.Equal(l.ToSlice(), other)
}

func (l *ComparableList[T]) Compact() {
	l.slice = slices.Compact(l.slice)
}

func (l *ComparableList[T]) Index(needle T) int {
	return slices.Index(l.slice, needle)
}

func (l *ComparableList[T]) Contains(needle T) bool {
	return slices.Contains(l.slice, needle)
}

// Sorts in-place
func (l *ComparableList[T]) SortFunc(test func(a T, b T) bool) {
	slices.SortFunc(l.slice, test)
}

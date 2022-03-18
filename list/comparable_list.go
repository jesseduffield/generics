package list

import "sort"

type ComparableList[T comparable] struct {
	*List[T]
}

func NewComparableList[T comparable]() *ComparableList[T] {
	return &ComparableList[T]{List: NewList[T]()}
}

func NewComparableListFromSlice[T comparable](slice []T) *ComparableList[T] {
	return &ComparableList[T]{List: NewListFromSlice(slice)}
}

func (l *ComparableList[T]) Equal(other *ComparableList[T]) bool {
	if len(l.slice) != len(other.slice) {
		return false
	}

	for i := range l.slice {
		if l.slice[i] != other.slice[i] {
			return false
		}
	}

	return true
}

func (l *ComparableList[T]) Compact() *ComparableList[T] {
	if len(l.slice) == 0 {
		return l
	}

	index := 1
	last := l.slice[0]
	for _, value := range l.slice[1:] {
		if value != last {
			l.slice[index] = value
			index++
			last = value
		}
	}

	l.slice = l.slice[:index]

	return l
}

func (l *ComparableList[T]) IndexOf(needle T) int {
	for index, value := range l.slice {
		if needle == value {
			return index
		}
	}
	return -1
}

func (l *ComparableList[T]) Contains(needle T) bool {
	return l.IndexOf(needle) != -1
}

func (l *ComparableList[T]) SortInPlace(test func(a T, b T) bool) {
	sort.Slice(l.slice, func(i, j int) bool {
		return test(l.slice[i], l.slice[j])
	})
}

func (l *ComparableList[T]) Sort(test func(a T, b T) bool) *ComparableList[T] {
	result := make([]T, len(l.slice))
	copy(result, l.slice)
	sort.Slice(result, func(i, j int) bool {
		return test(result[i], result[j])
	})
	return NewComparableListFromSlice(result)
}

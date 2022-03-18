package list

import (
	"fmt"
)

type List[T any] struct {
	slice []T
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func NewListFromSlice[T any](slice []T) *List[T] {
	return &List[T]{slice: slice}
}

func (l *List[T]) ToSlice() []T {
	return l.slice
}

// Mutative methods

func (l *List[T]) Push(v T) {
	l.slice = append(l.slice, v)
}

func (l *List[T]) Pop() error {
	if len(l.slice) == 0 {
		return fmt.Errorf("Cannot pop: list is empty")
	}
	l.slice = l.slice[0 : len(l.slice)-1]

	return nil
}

func (l *List[T]) MustPop() {
	if err := l.Pop(); err != nil {
		panic(err)
	}
}

func (l *List[T]) Concat(other *List[T]) {
	l.ConcatSlice(other.ToSlice())
}

func (l *List[T]) ConcatSlice(other []T) {
	l.slice = append(l.slice, other...)
}

func (l *List[T]) Insert(index int, value T) error {
	if index > len(l.slice) {
		return fmt.Errorf(
			"Cannot insert: index %d out of bounds for list of length %d", index, len(l.slice),
		)
	}

	if index < 0 {
		return fmt.Errorf(
			"Cannot insert: index %d is negative", index,
		)
	}

	l.slice = append(l.slice[:index], append([]T{value}, l.slice[index:]...)...)

	return nil
}

func (l *List[T]) MustInsert(index int, value T) {
	if err := l.Insert(index, value); err != nil {
		panic(err)
	}
}

func (l *List[T]) Remove(index int) error {
	if index > len(l.slice)-1 {
		return fmt.Errorf(
			"Cannot remove: index %d out of bounds for list of length %d", index, len(l.slice),
		)
	}

	if index < 0 {
		return fmt.Errorf(
			"Cannot remove: index %d is negative", index,
		)
	}

	l.slice = append(l.slice[:index], l.slice[index+1:]...)

	return nil
}

func (l *List[T]) MustRemove(index int) {
	if err := l.Remove(index); err != nil {
		panic(err)
	}
}

func (l *List[T]) Delete(from int, to int) {
	l.slice = append(l.slice[:from], l.slice[to:]...)
}

func (l *List[T]) FilterInPlace(test func(value T) bool) {
	newLength := 0
	for _, element := range l.slice {
		if test(element) {
			l.slice[newLength] = element
			newLength++
		}
	}

	l.slice = l.slice[:newLength]
}

func (l *List[T]) MapInPlace(f func(value T) T) {
	for i := range l.slice {
		l.slice[i] = f(l.slice[i])
	}
}

func (l *List[T]) ReverseInPlace() {
	for i, j := 0, len(l.slice)-1; i < j; i, j = i+1, j-1 {
		l.slice[i], l.slice[j] = l.slice[j], l.slice[i]
	}
}

// Non-mutative methods

func (l *List[T]) Filter(test func(value T) bool) *List[T] {
	result := make([]T, 0)
	for _, element := range l.slice {
		if test(element) {
			result = append(result, element)
		}
	}

	return NewListFromSlice(result)
}

func (l *List[T]) Map(f func(value T) T) *List[T] {
	result := make([]T, 0, len(l.slice))
	for _, element := range l.slice {
		result = append(result, f(element))
	}

	return NewListFromSlice(result)
}

func (l *List[T]) Clone() *List[T] {
	newSlice := make([]T, 0, len(l.slice))
	for _, value := range l.slice {
		newSlice = append(newSlice, value)
	}

	return NewListFromSlice(newSlice)
}

func (l *List[T]) Some(test func(value T) bool) bool {
	for _, value := range l.slice {
		if test(value) {
			return true
		}
	}

	return false
}

func (l *List[T]) Every(test func(value T) bool) bool {
	for _, value := range l.slice {
		if !test(value) {
			return false
		}
	}

	return true
}

func (l *List[T]) IndexFunc(f func(T) bool) int {
	for index, value := range l.slice {
		if f(value) {
			return index
		}
	}
	return -1
}

func (l *List[T]) ContainsFunc(f func(T) bool) bool {
	return l.IndexFunc(f) != -1
}

func (l *List[T]) Reverse() *List[T] {
	result := make([]T, len(l.slice))
	for i, j := 0, len(l.slice)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = l.slice[j], l.slice[i]
	}
	return NewListFromSlice(result)
}

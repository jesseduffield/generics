package list

import "github.com/jesseduffield/generics/slices"

// List is a struct which wraps a slice and provides convenience methods for it.
// Unfortunately due to some limitations in Go's type system, certain methods
// are not available e.g. Map.

type List[T any] struct {
	slice []T
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func NewFromSlice[T any](slice []T) *List[T] {
	return &List[T]{slice: slice}
}

func (l *List[T]) ToSlice() []T {
	return l.slice
}

// Mutative methods

func (l *List[T]) Push(v T) {
	l.slice = append(l.slice, v)
}

func (l *List[T]) Pop() T {
	var value T
	value, l.slice = slices.Pop(l.slice)
	return value
}

func (l *List[T]) Insert(index int, values ...T) {
	l.slice = slices.Insert(l.slice, index, values...)
}

func (l *List[T]) Append(values ...T) {
	l.slice = append(l.slice, values...)
}

func (l *List[T]) Prepend(values ...T) {
	l.slice = append(values, l.slice...)
}

func (l *List[T]) Remove(index int) {
	l.Delete(index, index+1)
}

func (l *List[T]) Delete(from int, to int) {
	l.slice = slices.Delete(l.slice, from, to)
}

func (l *List[T]) FilterInPlace(test func(value T) bool) {
	l.slice = slices.FilterInPlace(l.slice, test)
}

func (l *List[T]) MapInPlace(f func(value T) T) {
	slices.MapInPlace(l.slice, f)
}

func (l *List[T]) ReverseInPlace() {
	slices.ReverseInPlace(l.slice)
}

// Non-mutative methods

// Similar to Append but we leave the original slice untouched and return a new list
func (l *List[T]) Concat(values ...T) *List[T] {
	return NewFromSlice(slices.Concat(l.slice, values...))
}

func (l *List[T]) Filter(test func(value T) bool) *List[T] {
	return NewFromSlice(slices.Filter(l.slice, test))
}

// Unfortunately this does not support mapping from one type to another
// because Go does not yet (and may never) support methods defining their own
// type parameters. For that functionality you'll need to use the standalone
// Map function instead
func (l *List[T]) Map(f func(value T) T) *List[T] {
	return NewFromSlice(slices.Map(l.slice, f))
}

func (l *List[T]) Clone() *List[T] {
	return NewFromSlice(slices.Clone(l.slice))
}

func (l *List[T]) Some(test func(value T) bool) bool {
	return slices.Some(l.slice, test)
}

func (l *List[T]) Every(test func(value T) bool) bool {
	return slices.Every(l.slice, test)
}

func (l *List[T]) IndexFunc(f func(T) bool) int {
	return slices.IndexFunc(l.slice, f)
}

func (l *List[T]) ContainsFunc(f func(T) bool) bool {
	return slices.ContainsFunc(l.slice, f)
}

func (l *List[T]) Reverse() *List[T] {
	return NewFromSlice(slices.Reverse(l.slice))
}

func (l *List[T]) IsEmpty() bool {
	return len(l.slice) == 0
}

func (l *List[T]) Len() int {
	return len(l.slice)
}

func (l *List[T]) Get(index int) T {
	return l.slice[index]
}

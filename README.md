# Generics

This is a repo for some helper methods/structs that involve generics (added in Go 1.18).

## slices package

This package contains all the functions in the official slices [package](https://pkg.go.dev/golang.org/x/exp/slices#Insert) but adds extra functions as well, resulting in a superset of the official API. Any official functions are just forwarded to the official implementations. This allows you to use this package wherever you would otherwise use the official slices package.

As the official slices package evolves, so too will this package. If a function is added to the official package that does basically the same thing as a function from this package, we'll replace our function for the official function.

Here are the newly added functions.

```go
func Some[T any](slice []T, test func(T) bool) bool
func Every[T any](slice []T, test func(T) bool) bool
func Map[T any, V any](slice []T, f func(T) V) []V
func MapInPlace[T any](slice []T, f func(T) T)
func Filter[T any](slice []T, test func(T) bool) []T
func FilterInPlace[T any](slice []T, test func(T) bool) []T
func FilterMap[T any, E any](slice []T, test func(T) (bool, E)) []E
func FilterThenMap[T any, E any](slice []T, test func(T) bool, mapFn func(T) E) []E
func FlatMap[T any, V any](slice []T, f func(T) []V) []V
func Flatten[T any](slice [][]T) []T
func Reverse[T any](slice []T) []T
func ReverseInPlace[T any](slice []T)
func Prepend[T any](slice []T, values ...T) []T
func Remove[T any](slice []T, index int) []T
func Move[T any](slice []T, fromIndex int, toIndex int) []T
func Concat[T any](slice []T, values ...T) []T
func ContainsFunc[T any](slice []T, f func(T) bool) bool
func Pop[T any](slice []T) (T, []T)
func Shift[T any](slice []T) (T, []T)
```

## list package

This package provides a List struct which wraps a slice and gives you access to all the above functions, with a couple exceptions. For example, there's no Map method because go does not support type parameters on struct methods.

## set package

This package provides a Set struct with the following methods:

```go
Add(values ...T)
AddSlice(slice []T)
Remove(value T)
RemoveSlice(slice []T)
Includes(value T) bool
ToSlice() []T
```

## maps package

Provides some helper methods for maps:

```go
func Keys[Key comparable, Value any](m map[Key]Value) []Key
func Values[Key comparable, Value any](m map[Key]Value) []Value
func TransformValues[Key comparable, Value any, NewValue any
func TransformKeys[Key comparable, Value any, NewKey comparable](m map[Key]Value, fn func(Key) NewKey) map[NewKey]Value
func MapToSlice[Key comparable, Value any, Mapped any](m map[Key]Value, f func(Key, Value) Mapped) []Mapped
func Filter[Key comparable, Value any](m map[Key]Value, f func(Key, Value) bool) map[Key]Value
```

## Alternatives

Check out https://github.com/samber/lo for some more generic helper functions

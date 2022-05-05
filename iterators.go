// Package iterator contains an implementation of the map, filter, reduce pattern for Go.
package iterator

// Iterable is a generic interface for all iterables.
type Iterable[T any] interface {
	// Next returns the first or next value of T and true if a value is available.
	// If no more values are available or an error has occurred then a zero value of T and false is returned.
	Next() (T, bool)
	// Error returns nil after Next returned false when the iteration has completed successfully, otherwise
	// an error is returned.
	Error() error
}

// SliceIterator is a generic struct implementing an iterator that iterates over slices.
type SliceIterator[T any] struct {
	// idx has the position in the slice
	idx int
	// values contains the slice to iterate
	values []T
	// reverse contains a bool to tell the code to iterate the slice in reverse when this value is true
	reverse bool
}

// Next returns the first or next value of T and true if a value is available.
// If no more values are available or an error has occurred then a zero value of T and false is returned.
func (iter *SliceIterator[T]) Next() (T, bool) {
	if iter.idx < len(iter.values) {
		iter.idx++
	}
	if iter.idx == len(iter.values) {
		var t T
		return t, false
	}
	if iter.reverse {
		return iter.values[len(iter.values)-1-iter.idx], true
	}
	return iter.values[iter.idx], true
}

// Error returns nil after Next returned false when the iteration has completed successfully, otherwise
// an error is returned. The SliceIterator never returns an error.
func (iter *SliceIterator[T]) Error() error {
	return nil
}

// FromSlice creates a SliceIterator that iterates the provided slice.
func FromSlice[T any](values []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		idx:     -1,
		values:  values,
		reverse: false,
	}
}

// FromReverseSlice creates a SliceIterator that iterates the provided slice in reverse.
func FromReverseSlice[T any](values []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		idx:     -1,
		values:  values,
		reverse: true,
	}
}

// Algorithms
// Foreach

// ForEachFunc is the closure type that needs to be provided to ForEach
type ForEachFunc[T any] func(T)

// ForEach accepts an Iterable and calls the provided ForEachFunc closure with each value.
// An error is returned when an error during iteration has occurred.
func ForEach[T any](iter Iterable[T], f ForEachFunc[T]) error {
	for v, b := iter.Next(); b; v, b = iter.Next() {
		f(v)
	}
	return iter.Error()
}

// Map

// MapFunc is the closure type that needs to be provided to Map to perform the mapping operation with.
type MapFunc[T any, R any] func(T) R

// MapIterator is a struct the implements an Iterable that performs the map operation.
type MapIterator[T any, R any] struct {
	// srcItr is the Iterable this iterator pulls the original values from.
	srcItr Iterable[T]
	// mapFunc is the closure that performs the map operation.
	mapFunc MapFunc[T, R]
}

// Next returns the first or next value of T and true if a value is available.
// Each value is transformed with the provided MapFunc closure.
// If no more values are available or an error has occurred then a zero value of T and false is returned.
func (iter *MapIterator[T, R]) Next() (R, bool) {
	v, b := iter.srcItr.Next()
	if !b {
		var r R
		return r, false
	}
	return iter.mapFunc(v), true
}

// Error returns nil after Next returned false when the iteration has completed successfully, otherwise
// an error is returned.
func (iter *MapIterator[T, R]) Error() error {
	return iter.srcItr.Error()
}

// Map accepts an Iterable and MapFunc closure and creates a MapIterator that
// will perform the map operation on the values of the provided Iterable and
// returns the transformed values when iterated.
func Map[T any, R any](iter Iterable[T], f MapFunc[T, R]) *MapIterator[T, R] {
	return &MapIterator[T, R]{
		srcItr:  iter,
		mapFunc: f,
	}
}

// Filter

// PredicateFunc is the closure type that needs to be provided to Filter to perform the filter operation with.
// If the predicate returns true the value will be returned, otherwise it will be filtered.
type PredicateFunc[T any] func(T) bool

// FilterIterator is a struct the implements an Iterable that performs the filter operation.
type FilterIterator[T any] struct {
	// srcItr is the Iterable this iterator pulls the original values from.
	srcItr Iterable[T]
	// PredicateFunc is the closure that determines is the value needs to be filtered or not.
	predicate PredicateFunc[T]
}

// Next returns the first or next value of T and true if a value is available.
// Each value is checked against the provided PredicateFunc closure. When false is returned the value will be filtered.
// If no more values are available or an error has occurred then a zero value of T and false is returned.
func (iter *FilterIterator[T]) Next() (T, bool) {
	for v, b := iter.srcItr.Next(); b; v, b = iter.srcItr.Next() {
		if iter.predicate(v) {
			return v, true
		}
	}
	var t T
	return t, false
}

// Error returns nil after Next returned false when the iteration has completed successfully, otherwise
// an error is returned.
func (iter *FilterIterator[T]) Error() error {
	return iter.srcItr.Error()
}

// Filter accepts an Iterable and PredicateFunc closure and creates a FilterIterator that
// will perform the filter operation on the values of the provided Iterable and
// returns the filtered values when iterated.
func Filter[T any](iter Iterable[T], predicate PredicateFunc[T]) *FilterIterator[T] {
	return &FilterIterator[T]{
		srcItr:    iter,
		predicate: predicate,
	}
}

// Reduce

// ReduceFunc is the closure type that needs to be provided to Reduce to perform the reduce operation with.
type ReduceFunc[T any, R any] func(R, T) R

// Reduce accepts an Iterable, init value and ReduceFunc and reduces the values of the iterator to a single value by
// calling the ReduceFunc closure.
func Reduce[T any, R any](iter Iterable[T], init R, reducer ReduceFunc[T, R]) (R, error) {
	for v, b := iter.Next(); b; v, b = iter.Next() {
		init = reducer(init, v)
	}
	return init, iter.Error()
}

// ToSlice

// ToSlice renders the Iterable to a slice.
func ToSlice[T any](iter Iterable[T]) ([]T, error) {
	var result []T

	for v, b := iter.Next(); b; v, b = iter.Next() {
		result = append(result, v)
	}

	return result, iter.Error()
}

// Generators

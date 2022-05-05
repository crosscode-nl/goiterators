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

// GeneratorFunc is a closure that receives the count and repeat values and returns a generated value.
type GeneratorFunc[T any] func(c, r uint64) T

// GeneratingIterator is an iterator that will iterate repeat times and return values generated by
// the GeneratorFunc closure.
type GeneratingIterator[T any] struct {
	// count contains the number of executed iterations
	count uint64
	// repeat contains the number of iterations
	repeat uint64
	// generator contains a closure that returns a generated value
	generator GeneratorFunc[T]
}

// Next returns the first or next value of T and true if a value is available.
// If no more values are available or an error has occurred then a zero value of T and false is returned.
func (g *GeneratingIterator[T]) Next() (T, bool) {
	var t T
	var r bool
	if g.count < g.repeat {
		t = g.generator(g.count, g.repeat)
		r = true
		g.count++
	}
	return t, r
}

// Error returns nil after Next returned false when the iteration has completed successfully, otherwise
// an error is returned. The GeneratingIterator never returns an error.
func (g *GeneratingIterator[T]) Error() error {
	return nil
}

// Generate accepts a repeat count and a GeneratorFunc closure and returns a GeneratingIterator that repeats
// the given repeat times and returns values returned by the GeneratorFunc closure.
func Generate[T any](r uint64, gf GeneratorFunc[T]) *GeneratingIterator[T] {
	return &GeneratingIterator[T]{
		count:     0,
		repeat:    r,
		generator: gf,
	}
}

// The SignedIntegers interface defines all valid numerics to be used in the generic NumberGenerator
type SignedIntegers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// RepeatingIntegerGenerator accepts en initial value, a repeat value and a step value.
// The initial value is increased after each iteration step with the step value.
func RepeatingIntegerGenerator[T SignedIntegers](i T, r uint64, s T) *GeneratingIterator[T] {

	next := func(c uint64, r uint64) T {
		return i + (s * T(c))
	}

	return Generate(r, next)
}

// StepSequence accepts signed integer for start and end values. It will return GeneratingIterator
// that returns a sequence of values from start (inclusive) to end (inclusive).
// The sequence will increase or decrease the value returned with each iteration step with step.
// StepSequence will correct the sign of step for generating a sequence from start to end.
func StepSequence[T SignedIntegers](start T, end T, step T) *GeneratingIterator[T] {
	absStep := step
	if start > end {
		if step > 0 {
			step *= -1
		} else {
			absStep *= -1
		}
		return RepeatingIntegerGenerator(start, uint64((start-end)/absStep)+1, step)
	}
	if step < 0 {
		step *= -1
		absStep *= -1
	}
	return RepeatingIntegerGenerator(start, uint64((end-start)/absStep)+1, step)
}

// Sequence accepts signed integer for start and end values. It will return GeneratingIterator
// that returns a sequence of values from start (inclusive) to end (inclusive).
// The sequence will increase or decrease the value returned with each iteration step with 1.
// StepSequence will correct the sign of step for generating a sequence from start to end.
func Sequence[T SignedIntegers](start T, end T) *GeneratingIterator[T] {
	return StepSequence(start, end, 1)
}

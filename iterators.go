package iterator

type Iterable[T any] interface {
	Next() (T, bool)
	Error() error
}

type SliceIterator[T any] struct {
	idx     int
	values  []T
	error   error
	reverse bool
}

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

func (iter *SliceIterator[T]) Error() error {
	return iter.error
}

func FromSlice[T any](values []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		idx:     -1,
		values:  values,
		error:   nil,
		reverse: false,
	}
}

func FromReverseSlice[T any](values []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		idx:     -1,
		values:  values,
		error:   nil,
		reverse: true,
	}
}

// Algorithms
// Foreach

type ForEachFunc[T any] func(T)

func ForEach[T any](iter Iterable[T], f ForEachFunc[T]) error {
	for v, b := iter.Next(); b; v, b = iter.Next() {
		f(v)
	}
	return iter.Error()
}

// Map

type MapFunc[T any, R any] func(T) R

type MapIterator[T any, R any] struct {
	srcItr  Iterable[T]
	mapFunc MapFunc[T, R]
}

func (iter *MapIterator[T, R]) Next() (R, bool) {
	v, b := iter.srcItr.Next()
	if !b {
		var r R
		return r, false
	}
	return iter.mapFunc(v), true
}

func (iter *MapIterator[T, R]) Error() error {
	return iter.srcItr.Error()
}

func Map[T any, R any](iter Iterable[T], f MapFunc[T, R]) *MapIterator[T, R] {
	return &MapIterator[T, R]{
		srcItr:  iter,
		mapFunc: f,
	}
}

// Filter

type PredicateFunc[T any] func(T) bool

type FilterIterator[T any] struct {
	srcItr    Iterable[T]
	predicate PredicateFunc[T]
}

func (iter *FilterIterator[T]) Next() (T, bool) {
	for v, b := iter.srcItr.Next(); b; v, b = iter.srcItr.Next() {
		if iter.predicate(v) {
			return v, true
		}
	}
	var t T
	return t, false
}

func (iter *FilterIterator[T]) Error() error {
	return iter.srcItr.Error()
}

func Filter[T any](iter Iterable[T], predicate PredicateFunc[T]) *FilterIterator[T] {
	return &FilterIterator[T]{
		srcItr:    iter,
		predicate: predicate,
	}
}

// Reduce

type ReduceFunc[T any, R any] func(R, T) R

func Reduce[T any, R any](iter Iterable[T], init R, reducer ReduceFunc[T, R]) (R, error) {
	for v, b := iter.Next(); b; v, b = iter.Next() {
		init = reducer(init, v)
	}
	return init, iter.Error()
}

// ToSlice

func ToSlice[T any](iter Iterable[T]) ([]T, error) {
	var result []T

	for v, b := iter.Next(); b; v, b = iter.Next() {
		result = append(result, v)
	}

	return result, iter.Error()
}

// Generators

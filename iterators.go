package iterator

type Iterable[T any] interface {
	Next() bool
	Get() T
	Error() error
}

type SliceIterator[T any] struct {
	idx     int
	values  []T
	error   error
	reverse bool
}

func (iter *SliceIterator[T]) Next() bool {
	if iter.idx < len(iter.values) {
		iter.idx++
	}
	if iter.idx == len(iter.values) {
		return false
	}
	return true
}

func (iter *SliceIterator[T]) Get() T {
	if iter.idx == len(iter.values) || iter.idx < 0 {
		var result T
		return result
	}
	if iter.reverse {
		return iter.values[len(iter.values)-1-iter.idx]
	}
	return iter.values[iter.idx]
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
	for iter.Next() {
		f(iter.Get())
	}
	return iter.Error()
}

// Map

type MapFunc[T any, R any] func(T) R

type MapIterator[T any, R any] struct {
	srcItr  Iterable[T]
	mapFunc MapFunc[T, R]
}

func (iter *MapIterator[T, R]) Next() bool {
	return iter.srcItr.Next()
}

func (iter *MapIterator[T, R]) Get() R {
	return iter.mapFunc(iter.srcItr.Get())
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

func (iter *FilterIterator[T]) Next() bool {
	for iter.srcItr.Next() {
		if iter.predicate(iter.srcItr.Get()) {
			return true
		}
	}
	return false
}

func (iter *FilterIterator[T]) Get() T {
	return iter.srcItr.Get()
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
	for iter.Next() {
		init = reducer(init, iter.Get())
	}
	return init, iter.Error()
}

// ToSlice

func ToSlice[T any](iter Iterable[T]) ([]T, error) {
	var result []T
	for iter.Next() {
		result = append(result, iter.Get())
	}
	return result, iter.Error()
}

// Generators

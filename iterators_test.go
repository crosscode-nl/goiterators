package iterator

import (
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"iterator/testhelpers"
	"reflect"
	"testing"
)

type testFixture struct {
	slice     []int
	result    Iterable[int]
	predicate PredicateFunc[int]
}

var t testFixture

func aSliceIteratorIsReturned() error {
	if t.result == nil {
		return errors.New("expected: non nil result got: nil")
	}
	si := t.result.(*SliceIterator[int])
	if si == nil {
		return errors.New("expected: non nil si got: nil")
	}
	return nil
}

func nextReturnsTrueTimesAndThenReturnsFalse(num int) error {
	for ; num > 0; num-- {
		if t.result.Next() != true {
			return errors.New("expected: true got: false")
		}
	}
	if t.result.Next() != false {
		return errors.New("expected: false got: true")
	}
	return nil
}

func getAfterNextShouldReturn(listofints *godog.Table) error {
	values, err := testhelpers.TableToSliceOfInts(listofints)
	if err != nil {
		return err
	}
	for _, expected := range values {
		if t.result.Next() != true {
			return errors.New("expected: true got: false")
		}
		received := t.result.Get()
		if received != expected {
			return fmt.Errorf("expected: %d got: %d", expected, received)
		}
	}
	return nil
}

func aSliceIteratorIsReturnedWithIdxContaining(arg1 int) error {
	si := t.result.(*SliceIterator[int])
	if arg1 != si.idx {
		return fmt.Errorf("expected: %v got: %v", arg1, si.idx)
	}
	return nil
}

func aSliceIteratorIsReturnedWithErrorContainingNil() error {
	si := t.result.(*SliceIterator[int])
	if nil != si.error {
		return fmt.Errorf("expected: %v got: %v", nil, si.error)
	}
	return nil
}

func aSliceIteratorIsReturnedWithReverseContainingFalse() error {
	si := t.result.(*SliceIterator[int])
	if false != si.reverse {
		return fmt.Errorf("expected: %v got: %v", false, si.reverse)
	}
	return nil
}

func aSliceIteratorIsReturnedWithReverseContainingTrue() error {
	si := t.result.(*SliceIterator[int])
	if true != si.reverse {
		return fmt.Errorf("expected: %v got: %v", true, si.reverse)
	}
	return nil
}

func aSliceIteratorIsReturnedWithValuesContaining(listofints *godog.Table) error {
	si := t.result.(*SliceIterator[int])
	s, err := testhelpers.TableToSliceOfInts(listofints)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(s, si.values) {
		return fmt.Errorf("expected: %v got: %v", s, si.values)
	}
	return nil
}

func aSliceWithTheFollowingValues(listofints *godog.Table) (err error) {
	t.slice, err = testhelpers.TableToSliceOfInts(listofints)
	return
}

func fromSliceIsCalled() {
	t.result = FromSlice(t.slice)
}

func fromReverseSliceIsCalled() {
	t.result = FromReverseSlice(t.slice)
}

func aFilterIteratorIsReturned() error {
	if t.result == nil {
		return errors.New("expected: non nil result got: nil")
	}
	fi := t.result.(*FilterIterator[int])
	if fi == nil {
		return errors.New("expected: non nil fi got: nil")
	}
	return nil
}

func aPredicateThatOnlySelectsOddNumbers() {
	t.predicate = func(a int) bool {
		result := (a % 2) != 0
		fmt.Printf("%v %v", a, result)
		return result
	}
}

func anIterableWithTheFollowingValues(listofints *godog.Table) error {
	s, err := testhelpers.TableToSliceOfInts(listofints)
	if err != nil {
		return err
	}
	t.result = FromSlice(s)
	return nil
}

func filterIsCalled() {
	t.result = Filter(t.result, t.predicate)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	t = testFixture{}

	ctx.Step(`^a SliceIterator is returned$`, aSliceIteratorIsReturned)
	ctx.Step(`^a slice with the following values:$`, aSliceWithTheFollowingValues)
	ctx.Step(`^FromSlice is called$`, fromSliceIsCalled)
	ctx.Step(`^FromReverseSlice is called$`, fromReverseSliceIsCalled)
	ctx.Step(`^Next\(\) returns true (\d+) times and then returns false$`, nextReturnsTrueTimesAndThenReturnsFalse)
	ctx.Step(`^Get\(\) after Next\(\) should return:$`, getAfterNextShouldReturn)
	ctx.Step(`^a SliceIterator is returned with \.error containing nil$`, aSliceIteratorIsReturnedWithErrorContainingNil)
	ctx.Step(`^a SliceIterator is returned with \.idx containing (-\d+)$`, aSliceIteratorIsReturnedWithIdxContaining)
	ctx.Step(`^a SliceIterator is returned with \.reverse containing false$`, aSliceIteratorIsReturnedWithReverseContainingFalse)
	ctx.Step(`^a SliceIterator is returned with \.values containing:$`, aSliceIteratorIsReturnedWithValuesContaining)
	ctx.Step(`^a SliceIterator is returned with \.reverse containing true$`, aSliceIteratorIsReturnedWithReverseContainingTrue)
	ctx.Step(`^a FilterIterator is returned$`, aFilterIteratorIsReturned)
	ctx.Step(`^a predicate that only selects odd numbers$`, aPredicateThatOnlySelectsOddNumbers)
	ctx.Step(`^an Iterable with the following values:$`, anIterableWithTheFollowingValues)
	ctx.Step(`^Filter is called$`, filterIsCalled)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

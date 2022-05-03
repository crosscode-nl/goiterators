package iterator

import (
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"iterator/testhelpers"
	"reflect"
	"testing"
)

var slice []int
var result Iterable[int]

func aSliceIteratorIsReturned() error {
	if result == nil {
		return errors.New("expected: non nil result got: nil")
	}
	si := result.(*SliceIterator[int])
	if si == nil {
		return errors.New("expected: non nil si got: nil")
	}
	return nil
}

func nextReturnsTrueTimesAndThenReturnsFalse(num int) error {
	for ; num > 0; num-- {
		if result.Next() != true {
			return errors.New("expected: true got: false")
		}
	}
	if result.Next() != false {
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
		if result.Next() != true {
			return errors.New("expected: true got: false")
		}
		received := result.Get()
		if received != expected {
			return fmt.Errorf("expected: %d got: %d", expected, received)
		}
	}
	return nil
}

func aSliceIteratorIsReturnedWithIdxContaining(arg1 int) error {
	si := result.(*SliceIterator[int])
	if arg1 != si.idx {
		return fmt.Errorf("expected: %v got: %v", arg1, si.idx)
	}
	return nil
}

func aSliceIteratorIsReturnedWithErrorContainingNil() error {
	si := result.(*SliceIterator[int])
	if nil != si.error {
		return fmt.Errorf("expected: %v got: %v", nil, si.error)
	}
	return nil
}

func aSliceIteratorIsReturnedWithReverseContainingFalse() error {
	si := result.(*SliceIterator[int])
	if false != si.reverse {
		return fmt.Errorf("expected: %v got: %v", false, si.reverse)
	}
	return nil
}

func aSliceIteratorIsReturnedWithValuesContaining(listofints *godog.Table) error {
	si := result.(*SliceIterator[int])
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
	slice, err = testhelpers.TableToSliceOfInts(listofints)
	return
}

func fromSliceIsCalled() error {
	result = FromSlice(slice)
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	slice = []int{}

	ctx.Step(`^a SliceIterator is returned$`, aSliceIteratorIsReturned)
	ctx.Step(`^a slice with the following values:$`, aSliceWithTheFollowingValues)
	ctx.Step(`^FromSlice is called$`, fromSliceIsCalled)
	ctx.Step(`^Next\(\) returns true (\d+) times and then returns false$`, nextReturnsTrueTimesAndThenReturnsFalse)
	ctx.Step(`^Get\(\) after Next\(\) should return:$`, getAfterNextShouldReturn)
	ctx.Step(`^a SliceIterator is returned with \.error containing nil$`, aSliceIteratorIsReturnedWithErrorContainingNil)
	ctx.Step(`^a SliceIterator is returned with \.idx containing (-\d+)$`, aSliceIteratorIsReturnedWithIdxContaining)
	ctx.Step(`^a SliceIterator is returned with \.reverse containing false$`, aSliceIteratorIsReturnedWithReverseContainingFalse)
	ctx.Step(`^a SliceIterator is returned with \.values containing:$`, aSliceIteratorIsReturnedWithValuesContaining)

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

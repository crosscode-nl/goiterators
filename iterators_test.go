package iterator

import (
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"strconv"
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
	for _, row := range listofints.Rows {
		expected, err := strconv.Atoi(row.Cells[0].Value)
		if err != nil {
			return err
		}
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

func aSliceWithTheFollowingValues(listofints *godog.Table) error {
	for _, row := range listofints.Rows {
		i, err := strconv.Atoi(row.Cells[0].Value)
		if err != nil {
			return err
		}
		slice = append(slice, i)
	}
	return nil
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

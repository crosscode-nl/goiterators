package iterator

import (
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type testFixture struct {
	slice                   []int
	resultingIntIterator    Iterable[int]
	resultingStringIterator Iterable[string]
	predicate               PredicateFunc[int]
	mapper                  MapFunc[int, string]
	resultingSlice          []int
	reducer                 ReduceFunc[int, int]
	initialReduceValue      int
	counter                 ForEachFunc[int]
	count                   int
	sum                     int
}

var t testFixture

func toSliceOfInts(table *godog.Table) (result []int, err error) {
	var value int
	for _, row := range table.Rows {
		value, err = strconv.Atoi(row.Cells[0].Value)
		if err != nil {
			return
		}
		result = append(result, value)
	}
	return
}

func toSliceOfStrings(table *godog.Table) (result []string) {
	var value string
	for _, row := range table.Rows {
		value = row.Cells[0].Value
		result = append(result, value)
	}
	return
}

func nextReturnsTrueTimesAndThenReturnsFalse(num int) error {
	for ; num > 0; num-- {
		if t.resultingIntIterator.Next() != true {
			return errors.New("expected: true got: false")
		}
	}
	if t.resultingIntIterator.Next() != false {
		return errors.New("expected: false got: true")
	}
	return nil
}

func getAfterNextShouldReturn(listofints *godog.Table) error {
	values, err := toSliceOfInts(listofints)
	if err != nil {
		return err
	}
	for _, expected := range values {
		if t.resultingIntIterator.Next() != true {
			return errors.New("expected: true got: false")
		}
		received := t.resultingIntIterator.Get()
		if received != expected {
			return fmt.Errorf("expected: %d got: %d", expected, received)
		}
	}
	return nil
}

func aSliceIteratorIsReturnedWithIdxContaining(arg1 int) error {
	si := t.resultingIntIterator.(*SliceIterator[int])
	if arg1 != si.idx {
		return fmt.Errorf("expected: %v got: %v", arg1, si.idx)
	}
	return nil
}

func aSliceIteratorIsReturnedWithErrorContainingNil() error {
	si := t.resultingIntIterator.(*SliceIterator[int])
	if nil != si.error {
		return fmt.Errorf("expected: %v got: %v", nil, si.error)
	}
	return nil
}

func aSliceIteratorIsReturnedWithReverseContainingFalse() error {
	si := t.resultingIntIterator.(*SliceIterator[int])
	if false != si.reverse {
		return fmt.Errorf("expected: %v got: %v", false, si.reverse)
	}
	return nil
}

func aSliceIteratorIsReturnedWithReverseContainingTrue() error {
	si := t.resultingIntIterator.(*SliceIterator[int])
	if true != si.reverse {
		return fmt.Errorf("expected: %v got: %v", true, si.reverse)
	}
	return nil
}

func aSliceIteratorIsReturnedWithValuesContaining(listofints *godog.Table) error {
	si := t.resultingIntIterator.(*SliceIterator[int])
	s, err := toSliceOfInts(listofints)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(s, si.values) {
		return fmt.Errorf("expected: %v got: %v", s, si.values)
	}
	return nil
}

func aSliceWithTheFollowingValues(listofints *godog.Table) (err error) {
	t.slice, err = toSliceOfInts(listofints)
	return
}

func fromSliceIsCalled() {
	t.resultingIntIterator = FromSlice(t.slice)
}

func fromReverseSliceIsCalled() {
	t.resultingIntIterator = FromReverseSlice(t.slice)
}

func aPredicateThatOnlySelectsOddNumbers() {
	t.predicate = func(a int) bool {
		result := (a % 2) != 0
		fmt.Printf("%v %v", a, result)
		return result
	}
}

func anIterableWithTheFollowingValues(listofints *godog.Table) error {
	s, err := toSliceOfInts(listofints)
	if err != nil {
		return err
	}
	t.resultingIntIterator = FromSlice(s)
	return nil
}

func filterIsCalled() {
	t.resultingIntIterator = Filter(t.resultingIntIterator, t.predicate)
}

func aMapFunctionThatMultiplesTheValuesAndConvertsTheIntToAStringPrefixedWithTest() {
	t.mapper = func(i int) string {
		return "test" + strconv.Itoa(i*2)
	}
}

func getAfterNextShouldReturnTheFollowingValuesAsStrings(listofints *godog.Table) error {
	values := toSliceOfStrings(listofints)
	for _, expected := range values {
		if t.resultingStringIterator.Next() != true {
			return errors.New("expected: true got: false")
		}
		received := t.resultingStringIterator.Get()
		if received != expected {
			return fmt.Errorf("expected: %v got: %v", expected, received)
		}
	}
	return nil
}

func mapIsCalled() {
	t.resultingStringIterator = Map(t.resultingIntIterator, t.mapper)
}

func aSliceIsReturnedWithTheFollowingValues(listofints *godog.Table) error {
	s, err := toSliceOfInts(listofints)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(s, t.resultingSlice) {
		return fmt.Errorf("expected: %v got: %v", s, t.resultingSlice)
	}
	return nil
}

func toSliceIsCalled() (err error) {
	t.resultingSlice, err = ToSlice(t.resultingIntIterator)
	return
}

func aReduceFunctionThatSumsAllValues() {
	t.reducer = func(a, b int) int {
		return a + b
	}
}

func initialValueOf(init int) {
	t.initialReduceValue = init
}

func reduceIsCalled() (err error) {
	t.sum, err = Reduce(t.resultingIntIterator, t.initialReduceValue, t.reducer)
	return
}

func theReturnedSumIs(expected int) error {
	if t.sum != expected {
		return fmt.Errorf("expected: %v got: %v", expected, t.sum)
	}
	return nil
}

func aForeachFunctionThatSumsAndCountsTheCalls() {
	t.counter = func(i int) {
		t.count++
		t.sum += i
	}
}

func foreachIsCalled() error {
	return ForEach(t.resultingIntIterator, t.counter)
}

func theReturnedCountIs(expected int) error {
	if t.count != expected {
		return fmt.Errorf("expected: %v got: %v", expected, t.sum)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	t = testFixture{}

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
	ctx.Step(`^a predicate that only selects odd numbers$`, aPredicateThatOnlySelectsOddNumbers)
	ctx.Step(`^an Iterable with the following values:$`, anIterableWithTheFollowingValues)
	ctx.Step(`^Filter is called$`, filterIsCalled)
	ctx.Step(`^a map function that multiples the values and converts the int to a string, prefixed with test$`, aMapFunctionThatMultiplesTheValuesAndConvertsTheIntToAStringPrefixedWithTest)
	ctx.Step(`^Get\(\) after Next\(\) should return the following values as strings:$`, getAfterNextShouldReturnTheFollowingValuesAsStrings)
	ctx.Step(`^Map is called$`, mapIsCalled)
	ctx.Step(`^a slice is returned with the following values:$`, aSliceIsReturnedWithTheFollowingValues)
	ctx.Step(`^ToSlice is called$`, toSliceIsCalled)
	ctx.Step(`^a reduce function that sums all values$`, aReduceFunctionThatSumsAllValues)
	ctx.Step(`^initial value of (\d+)$`, initialValueOf)
	ctx.Step(`^Reduce is called$`, reduceIsCalled)
	ctx.Step(`^Foreach is called$`, foreachIsCalled)

	ctx.Step(`^a foreach function that sums and counts the calls$`, aForeachFunctionThatSumsAndCountsTheCalls)
	ctx.Step(`^The returned count is (\d+)$`, theReturnedCountIs)
	ctx.Step(`^The returned sum is (\d+)$`, theReturnedSumIs)

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

// Benchmarks

func BenchmarkFilter(b *testing.B) {

	var s []int

	for n := 0; n < 1000; n++ {
		s = append(s, n)
	}

	odd := func(v int) bool {
		return (v % 2) != 0
	}

	benchFunc := func() []int {
		si := FromSlice(s)
		fi := Filter[int](si, odd)
		ns, _ := ToSlice[int](fi)
		return ns
	}

	for n := 0; n < b.N; n++ {
		benchFunc()
	}
}

func BenchmarkFilterInIdiomaticGo(b *testing.B) {

	var s []int

	for n := 0; n < 1000; n++ {
		s = append(s, n)
	}

	benchFunc := func() []int {
		var ns []int
		for _, v := range s {
			if (v % 2) != 0 {
				ns = append(ns, v)
			}
		}
		return ns
	}

	for n := 0; n < b.N; n++ {
		benchFunc()
	}
}

func BenchmarkFilterMap(b *testing.B) {

	var s []int

	for n := 0; n < 1000; n++ {
		s = append(s, n)
	}

	odd := func(v int) bool {
		return (v % 2) != 0
	}

	benchFunc := func() []string {
		si := FromSlice(s)
		fi := Filter[int](si, odd)
		mi := Map[int, string](fi, strconv.Itoa)
		ns, _ := ToSlice[string](mi)
		return ns
	}

	for n := 0; n < b.N; n++ {
		benchFunc()
	}
}

func BenchmarkFilterMapInIdiomaticGo(b *testing.B) {

	var s []int

	for n := 0; n < 1000; n++ {
		s = append(s, n)
	}

	benchFunc := func() []string {
		var ns []string
		for _, v := range s {
			if (v % 2) != 0 {
				ns = append(ns, strconv.Itoa(v))
			}
		}
		return ns
	}

	for n := 0; n < b.N; n++ {
		benchFunc()
	}
}

func BenchmarkFilterMapReduce(b *testing.B) {

	var s []int

	for n := 0; n < 1000; n++ {
		s = append(s, n)
	}

	odd := func(v int) bool {
		return (v % 2) != 0
	}

	join := func(builder *strings.Builder, value string) *strings.Builder {
		if builder.Len() > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(value)
		return builder
	}

	benchFunc := func() string {
		si := FromSlice(s)
		fi := Filter[int](si, odd)
		mi := Map[int, string](fi, strconv.Itoa)
		sb, _ := Reduce[string, *strings.Builder](mi, &strings.Builder{}, join)
		return sb.String()
	}

	for n := 0; n < b.N; n++ {
		benchFunc()
	}
}

func BenchmarkFilterMapReduceInIdiomaticGo(b *testing.B) {

	var s []int

	for n := 0; n < 1000; n++ {
		s = append(s, n)
	}

	benchFunc := func() string {
		builder := strings.Builder{}
		for _, v := range s {
			if (v % 2) != 0 {
				if builder.Len() > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(strconv.Itoa(v))
			}
		}
		return builder.String()
	}

	for n := 0; n < b.N; n++ {
		benchFunc()
	}
}

func filterIntSlice(in []int, predicate func(int) bool) (output []int) {
	for _, v := range in {
		if predicate(v) {
			output = append(output, v)
		}
	}
	return
}

func mapIntSliceToStringSlice(in []int, mapper func(int) string) (output []string) {
	for _, v := range in {
		output = append(output, mapper(v))
	}
	return
}

func reduceStringSliceToString(in []string, init *strings.Builder, reducer func(*strings.Builder, string) *strings.Builder) (output string) {
	for _, v := range in {
		init = reducer(init, v)
	}
	return init.String()
}

func BenchmarkFilterMapDIY(b *testing.B) {

	var s []int

	for n := 0; n < 1000; n++ {
		s = append(s, n)
	}

	odd := func(v int) bool {
		return (v % 2) != 0
	}

	join := func(builder *strings.Builder, value string) *strings.Builder {
		if builder.Len() > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(value)
		return builder
	}

	benchFunc := func() string {
		return reduceStringSliceToString(mapIntSliceToStringSlice(filterIntSlice(s, odd), strconv.Itoa), &strings.Builder{}, join)
	}

	for n := 0; n < b.N; n++ {
		benchFunc()
	}
}

func BenchmarkFilterMapDIY2(b *testing.B) {

	var s []int

	for n := 0; n < 1000; n++ {
		s = append(s, n)
	}

	odd := func(v int) bool {
		return (v % 2) != 0
	}

	join := func(builder *strings.Builder, value string) *strings.Builder {
		if builder.Len() > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(value)
		return builder
	}

	benchFunc := func() string {
		sb := strings.Builder{}
		for _, v := range s {
			if odd(v) {
				join(&sb, strconv.Itoa(v))
			}
		}
		return sb.String()
	}

	for n := 0; n < b.N; n++ {
		benchFunc()
	}
}

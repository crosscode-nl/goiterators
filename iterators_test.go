package iterator

import (
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"reflect"
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

func aSliceWithTheFollowingValuesByList(listofints *godog.Table) error {
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
	ctx.Step(`^a slice with the following values by list$`, aSliceWithTheFollowingValuesByList)
	ctx.Step(`^FromSlice is called$`, fromSliceIsCalled)
	ctx.Step(`^Next\(\) returns true (\d+) times and then returns false$`, nextReturnsTrueTimesAndThenReturnsFalse)
	ctx.Step(`^Get\(\) after Next\(\) should return:$`, getAfterNextShouldReturn)

}

/*
func TestFilter(t *testing.T) {
	type args struct {
		iter      Iterable
		predicate PredicateFunc
	}
	tests := []struct {
		name string
		args args
		want Iterable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.iter, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterIterator_Error(t *testing.T) {
	type fields struct {
		srcItr    Iterable
		predicate PredicateFunc
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &FilterIterator{
				srcItr:    tt.fields.srcItr,
				predicate: tt.fields.predicate,
			}
			if err := iter.Error(); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFilterIterator_Get(t *testing.T) {
	type fields struct {
		srcItr    Iterable
		predicate PredicateFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &FilterIterator{
				srcItr:    tt.fields.srcItr,
				predicate: tt.fields.predicate,
			}
			if got := iter.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterIterator_Next(t *testing.T) {
	type fields struct {
		srcItr    Iterable
		predicate PredicateFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &FilterIterator{
				srcItr:    tt.fields.srcItr,
				predicate: tt.fields.predicate,
			}
			if got := iter.Next(); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	type args struct {
		iter Iterable
		f    ForEachFunc
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ForEach(tt.args.iter, tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFromReverseSlice(t *testing.T) {
	type args struct {
		values []T
	}
	tests := []struct {
		name string
		args args
		want Iterable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromReverseSlice(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromReverseSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}



func TestMap(t *testing.T) {
	type args struct {
		iter Iterable
		f    MapFunc
	}
	tests := []struct {
		name string
		args args
		want Iterable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.iter, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIterator_Error(t *testing.T) {
	type fields struct {
		srcItr  Iterable
		mapFunc MapFunc
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &MapIterator{
				srcItr:  tt.fields.srcItr,
				mapFunc: tt.fields.mapFunc,
			}
			if err := iter.Error(); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMapIterator_Get(t *testing.T) {
	type fields struct {
		srcItr  Iterable
		mapFunc MapFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   R
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &MapIterator{
				srcItr:  tt.fields.srcItr,
				mapFunc: tt.fields.mapFunc,
			}
			if got := iter.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIterator_Next(t *testing.T) {
	type fields struct {
		srcItr  Iterable
		mapFunc MapFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &MapIterator{
				srcItr:  tt.fields.srcItr,
				mapFunc: tt.fields.mapFunc,
			}
			if got := iter.Next(); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		iter    Iterable
		init    R
		reducer ReduceFunc
	}
	tests := []struct {
		name    string
		args    args
		want    R
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reduce(tt.args.iter, tt.args.init, tt.args.reducer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reduce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceIterator_Error(t *testing.T) {
	type fields struct {
		idx     int
		values  []T
		error   error
		reverse bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &SliceIterator{
				idx:     tt.fields.idx,
				values:  tt.fields.values,
				error:   tt.fields.error,
				reverse: tt.fields.reverse,
			}
			if err := iter.Error(); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSliceIterator_Get(t *testing.T) {
	type fields struct {
		idx     int
		values  []T
		error   error
		reverse bool
	}
	tests := []struct {
		name   string
		fields fields
		want   T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &SliceIterator{
				idx:     tt.fields.idx,
				values:  tt.fields.values,
				error:   tt.fields.error,
				reverse: tt.fields.reverse,
			}
			if got := iter.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceIterator_Next(t *testing.T) {
	type fields struct {
		idx     int
		values  []T
		error   error
		reverse bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &SliceIterator{
				idx:     tt.fields.idx,
				values:  tt.fields.values,
				error:   tt.fields.error,
				reverse: tt.fields.reverse,
			}
			if got := iter.Next(); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	type args struct {
		iter Iterable
	}
	tests := []struct {
		name    string
		args    args
		want    []T
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSlice(tt.args.iter)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
*/

func TestFromSlice(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want Iterable[int]
	}{
		{name: "Simple", args: args{values: []int{1, 2, 3, 4}}, want: &SliceIterator[int]{idx: -1, values: []int{1, 2, 3, 4}, error: nil, reverse: false}}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromSlice(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
} /*

func TestSliceIterator_Next(t *testing.T) {
	type fields struct {
		idx     int
		values  []int
		error   error
		reverse bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &SliceIterator{
				idx:     tt.fields.idx,
				values:  tt.fields.values,
				error:   tt.fields.error,
				reverse: tt.fields.reverse,
			}
			if got := iter.Next(); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/

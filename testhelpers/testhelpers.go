package testhelpers

import (
	"github.com/cucumber/godog"
	"strconv"
)

func TableToSliceOfInts(table *godog.Table) (result []int, err error) {
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

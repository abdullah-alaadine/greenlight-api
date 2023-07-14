package core

import "fmt"

type Table struct {
	numRows int
	pages   []Page
}

type TableResult int

const (
	TableFull TableResult = iota
	ExecuteSucceeded
	ExecuteFailed
)

func (t *Table) GetTableSize() string {
	return fmt.Sprintf("Table Size: %d", t.numRows)
}

package core

import (
	"fmt"

	"github.com/marianogappa/sqlparser/query"
)

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

func (t *Table) ExecuteInsert(q query.Query) TableResult {
	if t.numRows < MaxRows {
		row := RowFromQuery(q)
		pageNum := t.numRows / PageSize
		if len(t.pages) <= pageNum {
			t.pages = append(t.pages, Page{})
		}
		page := &t.pages[pageNum]
		if page.PageSize() < PageSize {
			page.AddRow(row)
			t.numRows++
			return ExecuteSucceeded
		}
		return ExecuteFailed
	}
	return TableFull
}

package main

import "testing"

func TestSplit(t *testing.T) {
	tables := []struct {
		min int
		max int
		character string
		minOut int
		maxOut int
	}{
		{0,127,"F",0,63},
		{0,63,"B",32,63},
		{32,63,"F",32,47},
		{32,47,"B",40,47},
		{40,47,"B",44,47},
		{44,47,"F",44,45},
		{44,45,"F",44,44},

		{0,7,"R",4,7},
		{4,7,"L",4,5},
		{4,5,"R",5,5},
	}

	for _, table := range tables {

		minValue, maxValue := Split(table.min, table.max, table.character)

		if minValue != table.minOut || maxValue != table.maxOut {
			t.Errorf(`Split went wrong, 
			given: %d, %d %s, 
			expected: %d %d, 
			got %d %d`,
			table.min, table.max, table.character,
			table.minOut, table.maxOut,
			minValue, maxValue)
		}
	}
}


func TestProcessInputLine(t *testing.T) {
	tables := []struct {
		input string
		expectedRow int
		expectedColumn int
		expectedId int
	}{
		{"BFFFBBFRRR",70,7,567},
		{"FFFBBBFRRR",14,7,119},
		{"BBFFBBFRLL",102,4,820},
	}

	for _, table := range tables {
		row, column, id := ProcessInputLine(table.input)

		if row != table.expectedRow ||
		column != table.expectedColumn ||
		id != table.expectedId {
			t.Errorf(`ProcessInputLine went wrong
			given: %s
			expected: row:%d column:%d id:%d
			got: row:%d column:%d id:%d`,
			table.input,
			table.expectedRow, table.expectedColumn, table.expectedId,
			row, column, id)
		}
	}
}

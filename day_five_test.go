package main

import "testing"

func TestSplitFront(t *testing.T) {
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

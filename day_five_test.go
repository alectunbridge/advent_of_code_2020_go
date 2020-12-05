package main

import "testing"

func TestSplitFront(t *testing.T) {
	minRowNum := 0 
	maxRowNum := 127
	character := "F"
	expectedMinValue := 0
	expectedMaxValue := 63
	
	minValue, maxValue := Split(minRowNum, maxRowNum,character)

	if minValue != expectedMinValue || maxValue != expectedMaxValue {
		t.Errorf(`Split went wrong, 
		given: %d, %d %s, 
		expected: %d %d, 
		got %d %d`,
		minRowNum, maxRowNum, character,
		expectedMinValue, expectedMaxValue,
		minValue, maxValue)
	}
}

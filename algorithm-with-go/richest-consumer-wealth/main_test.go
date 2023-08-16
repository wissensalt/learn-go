package main

import "testing"

func TestGetRichestCostumer(t *testing.T) {
	input := [][]int{{1, 2, 3}, {10, 11, 12}, {4, 5, 6}, {7, 8, 9}}
	actual := GetRichestCostumer(input)
	expected := 33
	if expected != actual {
		t.Errorf("Expected %v, But Actual %v\n", expected, actual)
	}
}

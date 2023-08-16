package main

import (
	"reflect"
	"testing"
)

func TestSumOneDArray(t *testing.T) {
	input := []int{4, 1, 5, 9, 2}
	actual := SumOneDArray(input)
	expected := []int{4, 5, 10, 19, 21}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v But Actual %v\n", expected, actual)
	}
}

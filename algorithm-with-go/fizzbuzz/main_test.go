package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	input := 15
	actual := FizzBuzz(input)
	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, But Actual %v\n", expected, actual)
	}
}

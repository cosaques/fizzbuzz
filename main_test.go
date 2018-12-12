package main

import "testing"

func TestErrorIfLimitLessThan1(t *testing.T) {
	if _, err := FizzBuzz("str1", "str2", 3, 5, 0); err == nil {
		t.Error("Should return error if limit is less than zero")
	}
}

func TestErrorIfIntIsZero(t *testing.T) {
	if _, err := FizzBuzz("str1", "str2", 0, 5, 100); err == nil {
		t.Error("Should return error if one of ints is zero")
	}

	if _, err := FizzBuzz("str1", "str2", 3, 0, 100); err == nil {
		t.Error("Should return error if one of ints is zero")
	}
}

func TestSimpleFizzBuzz(t *testing.T) {
	actual, _ := FizzBuzz("fizz", "buzz", 3, 5, 15)
	expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}

	if len(actual) != len(expected) {
		t.Fatalf("Expected list of length %d, but get length of %d", len(expected), len(actual))
	}

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("Expected \"%s\" on position %d, but get \"%s\"", expected[i], i, v)
		}
	}
}

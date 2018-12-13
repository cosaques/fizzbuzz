package tools

import (
	"errors"
	"strconv"
)

// FizzBuzz build a list of numbers from 1 to limit, where
// all multiples of int1 are replaced by string1,
// all multiples of int2 are replaced by string2,
// all multiples of int1 and int2 are replaced by string1string2
func FizzBuzz(string1 string, string2 string, int1 int, int2 int, limit int) ([]string, error) {
	if int1 == 0 || int2 == 0 {
		return nil, errors.New("int1, int2 should not be zero")
	}

	if limit < 1 {
		return nil, errors.New("limit should be greater than 1")
	}

	var result []string
	for i := 1; i <= limit; i++ {
		var item string
		if i%int1 == 0 {
			item += string1
		}
		if i%int2 == 0 {
			item += string2
		}
		if item == "" {
			item = strconv.Itoa(i)
		}

		result = append(result, item)
	}

	return result, nil
}

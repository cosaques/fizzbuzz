package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fizzbuzz, _ := FizzBuzz("fizz", "buzz", 3, 5, 15)
	fmt.Println(fizzbuzz)
}

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
	string1string2 := string1 + string2
	for i := 1; i <= limit; i++ {
		if i%int1 == 0 && i%int2 == 0 {
			result = append(result, string1string2)
		} else if i%int1 == 0 {
			result = append(result, string1)
		} else if i%int2 == 0 {
			result = append(result, string2)
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result, nil
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var int1, int2, limit int
	var string1, string2 string
	if err := getParams(r, &string1, &string2, &int1, &int2, &limit); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fizzBuzz, err := FizzBuzz(string1, string2, int1, int2, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(fizzBuzz)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getParams(r *http.Request, string1 *string, string2 *string, int1 *int, int2 *int, limit *int) error {
	var err error

	if *string1, err = getStrParam(r, "string1"); err != nil {
		return err
	}

	if *string2, err = getStrParam(r, "string2"); err != nil {
		return err
	}

	if *int1, err = getIntParam(r, "int1"); err != nil {
		return err
	}

	if *int2, err = getIntParam(r, "int2"); err != nil {
		return err
	}

	if *limit, err = getIntParam(r, "limit"); err != nil {
		return err
	}

	return nil
}

func getStrParam(r *http.Request, key string) (string, error) {
	if params, ok := r.URL.Query()[key]; ok && len(params) > 0 {
		return params[0], nil
	}
	return "", fmt.Errorf("Parameter \"%s\" is not set", key)
}

func getIntParam(r *http.Request, key string) (int, error) {
	if params, ok := r.URL.Query()[key]; ok && len(params) > 0 {
		v, err := strconv.Atoi(params[0])
		if err != nil {
			return 0, err
		}
		return v, nil
	}
	return 0, fmt.Errorf("Parameter \"%s\" is not set", key)
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

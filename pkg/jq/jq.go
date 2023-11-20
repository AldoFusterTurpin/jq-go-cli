package jq

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// GetIndexFromCmd receives a command like .[0] and returns the index from it
// and an error if any.
func ValidateIdentityOperator(identityOperator string) (bool, error) {
	patern := `^\.$`
	return regexp.Match(patern, []byte(identityOperator))
}

// ValidateAndGetIndexFromArrayExp receives an array expression
// like ".[0]"" and returns the index from it and an error if any.
func ValidateAndGetIndexFromArrayExp(arrayExpression string) (int, error) {
	matched, err := validateArrayExpr(arrayExpression)
	if err != nil {
		return -1, err
	}

	if !matched {
		return -1, fmt.Errorf("invalid array index")
	}

	return getIndexFromCmd(arrayExpression)
}

func validateArrayExpr(arrayExp string) (bool, error) {
	patern := `^\.\[\d+\]$`
	return regexp.Match(patern, []byte(arrayExp))
}

func getIndexFromCmd(arrayExpression string) (int, error) {
	s := strings.ReplaceAll(arrayExpression, ".", "")
	s = strings.ReplaceAll(s, "[", "")
	s = strings.ReplaceAll(s, "]", "")

	indexInt, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}
	return indexInt, nil
}

func GetIthElementFromArray(f interface{}, i int) (interface{}, error) {
	if i < 0 {
		return nil, fmt.Errorf("error, index can not be negative")
	}
	s, ok := f.([]interface{})
	if !ok {
		return nil, fmt.Errorf("error, can not get ith element as input is not an array")
	}
	length := len(s)
	if i >= length {
		return nil, fmt.Errorf("error, index out of bounds. Array has %v elements, but you specified element %v", length, i)
	}

	return s[i], nil
}

func ArrayIndexOperator(firstArg string, jsonArray interface{}) (interface{}, error){
  index, err := ValidateAndGetIndexFromArrayExp(firstArg)
	if err != nil {
		return nil, err
	}

	elementFromArray, err := GetIthElementFromArray(jsonArray, index)
	if err != nil {
		return nil, err
	}
  return elementFromArray, nil
}

func PrettifyAndPrint(f interface{}) {
	prettyJson, err := MarshalAndPrettify(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(prettyJson))
}

func MarshalAndPrettify(jsonElement interface{}) ([]byte, error) {
	prefix := ""
	indent := "    "
	prettyJsonBytes, err := json.MarshalIndent(jsonElement, prefix, indent)
	return prettyJsonBytes, err
}

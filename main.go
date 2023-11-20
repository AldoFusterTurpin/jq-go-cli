package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/AldoFusterTurpin/jq-go-cli/pkg/jq"
)

func main() {
	nArgs := len(os.Args)
	if nArgs > 2 {
		log.Fatalf("unsupported number of args: %v\n", nArgs-1)
	}

	jsonElement, err := readAllAndUnmarshalFrom(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	hasNoOptions := nArgs == 1
	if hasNoOptions {
		jq.PrettifyAndPrint(jsonElement)
		return
	}

	firstArg := os.Args[1]
	isIdentityOperator, err := jq.ValidateIdentityOperator(firstArg)
	if err != nil {
		log.Fatal(err)
	}

	if isIdentityOperator {
		jq.PrettifyAndPrint(jsonElement)
		return
	}

	elementFromArray, err := jq.ArrayIndexOperator(firstArg, jsonElement)
	if err != nil {
		log.Fatal(err)
	}

	prettyElement, err := jq.MarshalAndPrettify(elementFromArray)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(prettyElement))
}

func readAllAndUnmarshalFrom(r io.Reader) (interface{}, error) {
	jsonBytes, err := io.ReadAll(r)
	if err != nil && err != io.EOF {
		return nil, err
	}

	var jsonElement interface{}
	err = json.Unmarshal(jsonBytes, &jsonElement)
	if err != nil {
		return nil, err
	}
	return jsonElement, nil
}

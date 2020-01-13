package main

import (
	"encoding/json"
	"fmt"
	languages "github.com/utkarsh-pro/code-executor/Containers"
	"os"
)

// JSON struct describes the json payload
type JSON struct {
	Language string            `json:"language"`
	Command  string            `json:"command"`
	Files    []EvaluationFiles `json:"files"`
	Stdin    string            `json:"stdin"`
}

// EvaluationFiles hold the structure for the files that should be evaluated
type EvaluationFiles struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// Result struct
type Result struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Error  string `json:"error"`
}

func main() {
	decodedJSON := &JSON{}
	err := json.NewDecoder(os.Stdin).Decode(decodedJSON)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Type: %T\nContent: %+v\nAddress: %p\n", *decodedJSON, *decodedJSON, decodedJSON)
	_, ok := languages.Languages[decodedJSON.Language]
	if ok {

	}
}

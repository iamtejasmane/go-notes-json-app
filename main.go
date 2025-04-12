package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(title, content)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("enter note title:")
	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("enter note content")
	if err != nil {
		return "", "", err
	}

	return title, content, nil

}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var value string
	fmt.Println(&value)

	if value == "" {
		return "", errors.New("the fileds can not be empty")
	}

	return value, nil

}

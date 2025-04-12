package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/iamtejasmane/go-notes-json-app/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(userNote)
	userNote.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("enter note title: ")
	content := getUserInput("enter note content: ")

	return title, content

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/iamtejasmane/go-notes-json-app/note"
	"github.com/iamtejasmane/go-notes-json-app/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := getTodoData()

	t, err := todo.New(text)
	if err != nil {
		fmt.Print(err)
		return
	}
	t.Display()

	err = saverData(t)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("ToDo File save successfully.")

	// fmt.Println(userNote)
	userNote.Display()
	err = saverData(userNote)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func getTodoData() string {
	text := getUserInput("enter todo text: ")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("enter note title: ")
	content := getUserInput("enter note content: ")

	return title, content

}

func saverData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Println("File save successfully.")
	return nil
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

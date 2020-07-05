package main

import (
	"fmt"
	"github.com/yunussandikci/shareterm/cli/client"
	"io/ioutil"
	"os"
)

type App struct {
	input *os.File
	client.Client
}

func NewApp(file *os.File) App {
	return App{
		input:  file,
		Client: client.NewClient(),
	}
}

func (a App) execute() string {
	debugInput := os.Getenv("DEBUG_INPUT")
	var readInput string
	if debugInput == "" {
		 readInput = string(a.readInput())
	} else {
		readInput = debugInput
	}
	pasteCreateResponse := a.Client.CreatePaste(readInput)
	return fmt.Sprintf("%s",pasteCreateResponse.URL)
}

func (a App) readInput() []byte {
	bytes, readErr := ioutil.ReadAll(os.Stdin)
	if readErr != nil {
		panic(readErr)
	}
	return bytes
}

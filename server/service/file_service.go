package service

import (
	"fmt"
	"github.com/sethvargo/go-password/password"
	"io/ioutil"
	"os"
)

const PasteFolder = "data"

type FileService interface {
	WriteFile(fileName string, bytes []byte) error
	ReadFile(fileName string) ([]byte, error)
	GenerateFileName() (string, error)
}

type fileService struct{}

func NewFileService() FileService {
	return fileService{}
}

func (f fileService) WriteFile(fileName string, bytes []byte) error {
	filePath := getFilePath(fileName)
	writeErr := ioutil.WriteFile(filePath, bytes, 0644)
	if writeErr != nil {
		return writeErr
	}
	return nil
}

func (f fileService) ReadFile(fileName string) ([]byte, error) {
	filePath := getFilePath(fileName)
	bytes, readErr := ioutil.ReadFile(filePath)
	if readErr != nil {
		return nil, readErr
	}
	return bytes, nil
}

func (f fileService) GenerateFileName() (string, error) {
	for {
		generatedName, generateErr :=
			password.Generate(12, 6, 0, true, false)
		if generateErr != nil {
			return "", generateErr
		}
		generatedFilePath := getFilePath(generatedName)
		if !fileExists(generatedFilePath) {
			return generatedName, nil
		}
	}
}

func fileExists(filename string) bool {
	fileInfo, fileError := os.Stat(filename)
	if os.IsNotExist(fileError) {
		return false
	}
	return !fileInfo.IsDir()
}

func getFilePath(fileName string) string {
	return fmt.Sprintf("%s/%s.paste", getOrCreateFolder(PasteFolder), fileName)
}

func getOrCreateFolder(folder string) string {
	currentPath, currentPathErr := os.Getwd()
	if currentPathErr != nil {
		panic(currentPathErr)
	}
	folderPath := fmt.Sprintf("%s/%s", currentPath, folder)
	_, folderErr := os.Stat(folderPath)
	if folderErr != nil {
		currentPathErr := os.Mkdir(folderPath, os.ModePerm)
		if currentPathErr != nil {
			panic(currentPathErr)
		}
	}
	return folderPath
}

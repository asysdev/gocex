package utils

import (
	"log"
	"os"
)

func CreateFile(fileName string) *os.File {

	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return nil
	}

	return file
}

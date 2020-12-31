package router

import (
	"log"
	"os"
	"bufio"
)

func open(fileName string) *os.File {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func Read(fileName string) *OSGIRouter {
	
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileName := "textfile.txt"

	data := []byte("Hello, Workd!")

	err := ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatalf("Can't write to file: %s", err)
	}

	readData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Cannot read from file: %s", err)
	}
	fmt.Printf("Data: %s\n", string(readData))

	newFileName := "renamed_textfile.txt"
	err = os.Rename(fileName, newFileName)
	if err != nil {
		log.Fatalf("Cannot rename file: %s", err)
	}

	err = os.Remove(newFileName)
	if err != nil {
		log.Fatalf("Cannot remove file: %s", err)
	}
}

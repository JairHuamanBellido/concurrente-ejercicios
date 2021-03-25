package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func write(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func main() {

	// defer es una función para la limpieza de recursos.
	if err := write("readme.txt", "This is a readme file"); err != nil {
		log.Fatal("failed to write file:", err)
	}

	// Si existen múltiples sentencias de defer, actua como FIFO
	defer fmt.Println("End 3")
	defer fmt.Println("End 2")
	defer fmt.Println("End 1")
}

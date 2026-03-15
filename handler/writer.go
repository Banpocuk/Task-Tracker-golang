package handler

import (
	"fmt"
	"os"
)

func write(input_write string) {
	file, err := os.Create("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := []byte(input_write)

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
	//1. Read the commands
	//2. Write here
	fmt.Println("Writing done")
}

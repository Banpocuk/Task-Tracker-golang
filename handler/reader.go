package handler

import (
	"fmt"
	"os"
)


func reader()  {
    file, err := os.Create("data.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    data := []byte("hello")

    _, err = file.Write(data)
    if err != nil {
        panic(err)
    }
	//1. Read the commands
	//2. Give insturtion to the writer
}
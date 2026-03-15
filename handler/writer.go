package handler

import (
    "fmt"
    "os"
)

func writer()  {
    file, err := os.Open("data.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    buf := make([]byte, 10)

    n, err := file.Read(buf)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(buf[:n]))
	//1. Answer the user
	//2.Write the answers to the file
}
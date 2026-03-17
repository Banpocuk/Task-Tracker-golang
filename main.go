package main

import (
	C "Task-Tracker-golang/handler"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input string

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write 'help' to see commands")
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		C.Commands_holder(input)
	}
}

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
	fmt.Println("Welcome to my task-tracker \nWrite 'help' to see commands")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	C.Commands_holder(input)
}

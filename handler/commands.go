package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var read *bufio.Reader = bufio.NewReader(os.Stdin)

func Commands_holder(input string) {
	switch input {
	case "help":
		fmt.Println("1.exit \n2.write \n3.read \n4.delete")
	case "exit":
		fmt.Println("exiting...")
		return
	case "write":
		input := input_holder("write")
		writer(input)
	case "read":
		input := input_holder("read")
		reader(input)
	case "delete":
		input := input_holder("delete")
		// delete func
	default:
		fmt.Println("Invalid command")
	}

	//All the commands will be stored here
}

func input_holder(command string) string {
	fmt.Print("What to ", command, ": ")
	input, _ := read.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

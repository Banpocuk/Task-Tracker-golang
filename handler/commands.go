package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var read *bufio.Reader = bufio.NewReader(os.Stdin)

func Commands_holder(input string) {
	switch [0]input {
	case "help":
		fmt.Println("1.exit \n2.write \n3.read \n4.delete")
	case "exit":
		fmt.Println("exiting...")
		return
	case "write":
		writer(input[1])
	case "read":
		input := input_holder("read")
		reader(input)
	case "delete":
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

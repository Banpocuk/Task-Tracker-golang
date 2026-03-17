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
		fmt.Println("1.write \n2.delete \n3.list 4.list by id 5.list by status \n6.update \n7.exit") // Add more commands here
	case "exit":
		fmt.Println("exiting...")
		os.Exit(0)
	case "write":
		input := input_holder("write?")
		status := input_holder("status? todo/in-progress/done")
		writer(input, status)
	case "list":
		reader()
	case "list by id":
		id := input_holder("list by id?")
		reader_by_id(id)
	case "list by status":
		status := input_holder("list by status?")
		reader_by_status(status)
	case "update by id":
		id := input_holder("update by id?")
		update_by_id(id)
	case "update by status":
		status := input_holder("update by status?")
		update_by_status(status)
	case "delete":
		input := input_holder("delete?")
		delete_task(input)
	default:
		fmt.Println("Invalid command 👾")
	}

	//All the commands will be stored here
}

func input_holder(command string) string {
	fmt.Print("What to ", command, " 🤔 : ")
	input, _ := read.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

package handler

import (
	"fmt"
	"os"
	"strings"
)

var ts, yn, task, status string
var update_id int

func default_updater(input string, index int) {

	fmt.Println("What to update? task/status")
	fmt.Scan(&ts)
	switch ts {
	case "task":
		task = input_holder("new task")
		update_id = 1
	case "status":
		status = input_holder("new status")
		update_id = 2
	}

	fmt.Println("Are you sure? 🤔 y/n")
	fmt.Scan(&yn)
	switch yn {
	case "y":
		data, _ := os.ReadFile("data.txt")
		lines := strings.Split(string(data), "\n")

		for _, line := range lines {
			parts := strings.Split(line, "|")
			if parts[index] == input { // looking for task with status/id
				parts[update_id] = task

				fmt.Println("Task updated ✅: ", strings.Join(parts, "|"))
			}
		}
	case "n":
		fmt.Println("Update cancelled ❌")
	default:
		fmt.Println("Invalid input ❌ Please enter 'y' or 'n'.")
	}
}

func update_by_id(id string) {
	default_updater(id, 0)
}

func update_by_status(status string) {
	default_updater(status, 2)
}

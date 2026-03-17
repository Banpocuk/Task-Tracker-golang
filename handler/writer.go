package handler

import (
	"fmt"
	"os"
	"strings"
)

func writer(input_write string, status string) {
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	id := get_id()
	entry := fmt.Sprintf("%d|%s|%s\n", id, input_write, status)
	file.WriteString(entry)
	fmt.Println("Task added with ID ✅:", id)
}

func get_id() int {
	// AI code. (down here)
	data, err := os.ReadFile("data.txt")
	if err != nil || strings.TrimSpace(string(data)) == "" {
		return 1
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return len(lines) + 1
}

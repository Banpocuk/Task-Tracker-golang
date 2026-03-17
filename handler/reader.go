package handler

import (
	"fmt"
	"os"
	"strings"
)

func reader() {
	// Читаем весь файл
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 128)

	n, err := file.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf[:n]))

	// 1. Вернуть все задачи
	// 2. По айди вернуть задачи
	// 3. Вернуть задачи по статусу
	/* 4. update задачи
	4.1 по айди
	4.2 по статусу
	*/
}

func reader_by_id(id string) {
	default_id_status_reader(id, 0)
}

func reader_by_status(status string) {
	default_id_status_reader(status, 2)
}

func default_id_status_reader(input string, index int) string {
	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		parts := strings.Split(line, "|")
		if parts[index] == input { // looking for task with status
			fmt.Println("Task found ✅: ", line)

			return line
		}
	}
	return ""
}

/*
	fmt.Println(parts[0]) // id
	fmt.Println(parts[1]) // task
	fmt.Println(parts[2]) // status
*/

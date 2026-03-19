package handler

import (
	"fmt"
	"os"
	"strings"
)

func delete_task(input_id string) {
	// 1. Читаем весь файл
	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(string(data), "\n")

	// 2. Фильтруем — убираем строку которую надо удалить
	var new_lines []string
	for _, line := range lines {
		parts := strings.Split(line, "|")
		if parts[0] != input_id {
			new_lines = append(new_lines, line)
			fmt.Println("Delete " + parts[0] + " " + parts[1] + " " + parts[2])
		}
	}

	// 3. Записываем обратно
	result := strings.Join(new_lines, "\n")
	os.WriteFile("data.txt", []byte(result), 0644)

	// 1. delete by id
	// 2. delete by status
	// 3. delete all
	// 4. deleting history
	// 5. undo delete
}

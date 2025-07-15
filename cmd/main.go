package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/wolfalex23/rental-tracker/internal/data"
	"github.com/wolfalex23/rental-tracker/internal/ui"
)

func menuLoop() {
	for {
		fmt.Println(`
Меню:
1. Вывести список филиалов
2. Вывести филиал по id
3. Добавить филиал
4. Удалить филиал
5. Изменить филиал
6. Выход
`)
		fmt.Print("Ваш выбор: ")
		choice, err := readChoice()
		if err != nil {
			fmt.Println("Ошибка при чтении выбора:", err)
			continue
		}

		switch choice {
		case 1:
			_ = ui.ListCmd.Execute()
		case 2:
			_ = ui.ListBranchCmd.Execute()
		case 3:
			_ = ui.AddCmd.Execute()
		case 4:
			ui.DeleteCmd.Execute()
		case 5:
			ui.UpdateCmd.Execute()
		case 6:
			fmt.Println("Работа программы завершена.")
			return
		default:
			fmt.Println("Неверный выбор. Выберите пункт из меню.")
		}
	}
}

// Чтение выбора пользователя
func readChoice() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, err
	}
	return choice, nil
}

func main() {
	dbPath := "branches.db"

	err := data.Init(dbPath)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)

	}

	defer data.Close()
	menuLoop()
	ui.Execute()

}

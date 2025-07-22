package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wolfalex23/rental-tracker/internal/handlers"
)

func MenuLoop() {
	for {
		fmt.Print(`
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
			handlers.ListHandler()
		case 2:
			handlers.ListOneHandler()
		case 3:
			handlers.AddHandler()
		case 4:
			handlers.DeleteHandler()
		case 5:
			handlers.UpdateHandler()
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

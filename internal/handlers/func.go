package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Вспомогательная функция для удобства чтения строк

func promptUser(prompt string) (string, bool) {
	for {
		fmt.Print(prompt + " (или введите 'esc' для выхода; нажмите Enter,чтобы пропустить ввод): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if strings.EqualFold(input, "esc") {
			return "", false // Вернуть пустую строку или специальный маркер, означающий отмену
		}

		return input, true
	}
}

// Чтение плавающих чисел

func readPositiveFloat(prompt string) (float64, bool) {
	for {
		strVal, ok := promptUser(prompt)
		if !ok {
			return 0, false // Вернётся при нажатии ESC, сигнализируя об отмене
		} else if strVal == "" && ok {
			strVal = "0"
		}

		val, err := strconv.ParseFloat(strVal, 64)
		if err != nil || val < 0 {
			fmt.Printf("Ошибка при вводе значения '%s'. Попробуйте снова.\n", prompt)
			continue
		}

		return val, true // Значение введено корректно
	}
}

func readPositiveInt(prompt string) (int, bool) {
	for {
		strVal, ok := promptUser(prompt)
		if !ok {
			return 0, false // Вернётся при нажатии ESC, сигнализируя об отмене
		} else if strVal == "" && ok {
			strVal = "0"
		}

		val, err := strconv.Atoi(strVal)
		if err != nil || val < 0 {
			fmt.Printf("Ошибка при вводе значения '%s'. Попробуйте снова.\n", prompt)
			continue
		}
		return val, true
	}
}

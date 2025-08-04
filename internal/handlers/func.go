package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Вспомогательная функция для удобства чтения строк

func promptUser(prompt string, showInstruction bool) (string, bool) {
	for {
		var fullPrompt string
		if showInstruction {
			fullPrompt = prompt + " (или введите 'esc' для выхода; нажмите Enter,чтобы пропустить ввод): "
		} else {
			fullPrompt = prompt + ": "
		}

		fmt.Print(fullPrompt)
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

func readPositiveFloat(prompt string, showInstruction bool) (float64, bool) {
	for {
		strVal, ok := promptUser(prompt, showInstruction)
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

func readPositiveInt(prompt string, showInstruction bool) (int, bool) {
	for {
		strVal, ok := promptUser(prompt, showInstruction)
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

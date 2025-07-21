package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Вспомогательная функция для удобства чтения строк

func promptUser(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// Чтение плавающих чисел

func readPositiveFloat(prompt string) float64 {
	for {
		strVal := promptUser(prompt)
		val, err := strconv.ParseFloat(strVal, 64)
		if err != nil || val <= 0 {
			fmt.Printf("Ошибка при вводе значения '%s', попробуйте ещё раз.\\n", prompt[:len(prompt)-2])
			continue
		}
		return val
	}
}

func readPositiveInt(prompt string) int {
	for {
		strVal := promptUser(prompt)
		val, err := strconv.Atoi(strVal)
		if err != nil || val <= 0 {
			fmt.Printf("Ошибка при вводе значения '%s', попробуйте ещё раз.\\n", prompt[:len(prompt)-2])
			continue
		}
		return val
	}
}

package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/wolfalex23/rental-tracker/internal/data"
	"github.com/wolfalex23/rental-tracker/internal/model"
)

// Вспомогательная функция для удобства чтения строк
func promptUser(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

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

func AddHandler() {
	department := promptUser("Название Филиала: ")
	address := promptUser("Адрес филиала: ")
	contract := promptUser("Номер договора: ")

	aria := readPositiveFloat("Площадь м2: ")
	meterInYear := readPositiveFloat("Стоимость м2 в год: ")
	totalInYear := readPositiveFloat("Итого в год: ")

	branch := model.Branch{
		Department:  department,
		Address:     address,
		Contract:    contract,
		Aria:        aria,
		MeterInYear: meterInYear,
		TotalInYear: totalInYear,
	}

	err := data.AddBranch(&branch)
	if err != nil {
		fmt.Printf("Ошибка при добавлении филиала: %v\n", err)
		return
	}
	fmt.Println("Филиал успешно добавлен.")
}

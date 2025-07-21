package handlers

import (
	"fmt"

	"github.com/wolfalex23/rental-tracker/internal/data"
	"github.com/wolfalex23/rental-tracker/internal/model"
)

// Вспомогательная функция для удобства чтения строк

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

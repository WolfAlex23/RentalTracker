package handlers

import (
	"fmt"

	"github.com/wolfalex23/rental-tracker/internal/data"
	"github.com/wolfalex23/rental-tracker/internal/model"
)

func AddHandler() {
	department, ok := promptUser("Название филиала")
	if !ok {
		return
	}

	address, ok := promptUser("Адрес филиала")
	if !ok {
		return
	}
	contract, ok := promptUser("Номер договора")
	if !ok {
		return
	}

	aria, ok := readPositiveFloat("Площадь м2")
	if !ok {
		return
	}
	meterInYear, ok := readPositiveFloat("Стоимость м2 в год")
	if !ok {
		return
	}
	totalInYear, ok := readPositiveFloat("Итого в год")
	if !ok {
		return
	}

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

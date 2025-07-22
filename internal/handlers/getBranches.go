package handlers

import (
	"fmt"

	"github.com/wolfalex23/rental-tracker/internal/data"
)

// Функция обработки вывода списка филиалов
func ListHandler() {
	branches, err := data.GetBranches()
	if err != nil {
		fmt.Printf("Ошибка при получении списка филиалов: %s\\n", err.Error())
		return
	}

	// Если филиалы найдены, выводим каждый филиал на экран
	for _, branch := range branches {
		fmt.Printf("ID: %d, Department: %s, Address: %s, Contract: %s, Aria: %.2f, MeterInYear: %.2f, TotalInYear: %.2f, UpdatedAt: %s\\n",
			branch.ID,
			branch.Department,
			branch.Address,
			branch.Contract,
			branch.Aria,
			branch.MeterInYear,
			branch.TotalInYear,
			branch.UpdatedAt)
	}

	if len(branches) == 0 {
		fmt.Println("Филиалы отсутствуют")
	}
}

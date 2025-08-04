package handlers

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"

	"github.com/wolfalex23/rental-tracker/internal/data"
	"github.com/wolfalex23/rental-tracker/internal/model"
)

func UpdateHandler() {

	id, ok := readPositiveInt("Номер филиала ", true)
	if !ok {
		return
	}

	branch, err := data.GetBranch(id)
	if err != nil {
		fmt.Printf("Ошибка при получении инфо филиала: %s\n", err.Error())
		return
	}
	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{
			Settings: tw.Settings{Separators: tw.Separators{BetweenRows: tw.On}},
		})))

	// Заголовки таблицы
	header := []string{"ID", "Department", "Address", "Contract", "Aria", "MeterInYear", "TotalInYear", "UpdatedAt"}
	table.Header(header)

	var prntBranch [][]string

	prntBranch = append(prntBranch, []string{
		fmt.Sprint(branch.ID),
		branch.Department,
		branch.Address,
		branch.Contract,
		fmt.Sprintf("%.2f", branch.Aria),
		fmt.Sprintf("%.2f", branch.MeterInYear),
		fmt.Sprintf("%.2f", branch.TotalInYear),
		branch.UpdatedAt.Format("2006-01-02 15:04:05"),
	})

	table.Bulk(prntBranch)

	// Рендерим таблицу
	table.Render()

	if branch == nil {
		fmt.Println("Филиал отсутствуют.")
	}

	department, ok := promptUser("Название филиала", true)
	if !ok {
		return
	} else if department == "" {
		department = branch.Department
	}

	address, ok := promptUser("Адрес филиала", true)
	if !ok {
		return
	} else if address == "" {
		address = branch.Address
	}

	contract, ok := promptUser("Номер договора", true)
	if !ok {
		return
	} else if contract == "" {
		contract = branch.Contract
	}

	aria, ok := readPositiveFloat("Площадь м2", true)
	if !ok {
		return
	} else if aria == 0 {
		aria = branch.Aria
	}

	meterInYear, ok := readPositiveFloat("Стоимость м2 в год", true)
	if !ok {
		return
	} else if meterInYear == 0 {
		meterInYear = branch.MeterInYear
	}

	totalInYear, ok := readPositiveFloat("Итого в год", true)
	if !ok {
		return
	} else if totalInYear == 0 {
		totalInYear = branch.TotalInYear
	}

	branch = &model.Branch{
		ID:          id,
		Department:  department,
		Address:     address,
		Contract:    contract,
		Aria:        aria,
		MeterInYear: meterInYear,
		TotalInYear: totalInYear,
	}

	err = data.UpdateBranch(branch)
	if err != nil {
		fmt.Printf("Ошибка при обновлении информации о филиале: %v\n", err)
		return
	}
	fmt.Println("Филиал успешно обновлен.")
}

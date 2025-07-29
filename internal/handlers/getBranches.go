package handlers

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"

	"github.com/wolfalex23/rental-tracker/internal/data"
)

// Функция обработки вывода списка филиалов
func ListHandler() {
	branches, err := data.GetBranches()
	if err != nil {
		fmt.Printf("Ошибка при получении списка филиалов: %s\\n", err.Error())
		return
	}

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{
			Settings: tw.Settings{Separators: tw.Separators{BetweenRows: tw.On}},
		})))

	// Заголовки таблицы
	header := []string{"ID", "Department", "Address", "Contract", "Aria", "MeterInYear", "TotalInYear", "UpdatedAt"}
	table.Header(header)

	var rows [][]string

	for _, branch := range branches {
		rows = append(rows, []string{
			fmt.Sprint(branch.ID),
			branch.Department,
			branch.Address,
			branch.Contract,
			fmt.Sprintf("%.2f", branch.Aria),
			fmt.Sprintf("%.2f", branch.MeterInYear),
			fmt.Sprintf("%.2f", branch.TotalInYear),
			branch.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// Добавляем данные в таблицу одним махом
	table.Bulk(rows)

	// Рендерим таблицу
	table.Render()

	if len(branches) == 0 {
		fmt.Println("Филиалы отсутствуют.")
	}
}

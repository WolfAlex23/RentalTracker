package ui

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wolfalex23/rental-tracker/internal/data"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Показывает список всех филиалов",
	Long:  "Эта команда выводит подробный список всех зарегистрированных филиалов.",
	RunE: func(cmd *cobra.Command, args []string) error {
		branches, err := data.GetBranches()
		if err != nil {
			return err
		}

		if len(branches) == 0 {
			fmt.Println("Нет ни одного филиала.")
			return nil
		}

		fmt.Println("Список филиалов:")
		for _, branch := range branches {
			cmd.Println(fmt.Sprintf(`
Информация о филиале:
ID: %v
Отдел: %s
Адрес: %s
Контракт: %s
Площадь: %.2f
Метры в год: %.2f м²
Всего в год: %.2f руб.
`,
				branch.ID,
				branch.Department,
				branch.Address,
				branch.Contract,
				branch.Aria,
				branch.MeterInYear,
				branch.TotalInYear,
			))
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
}

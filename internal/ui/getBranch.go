package ui

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wolfalex23/rental-tracker/internal/data"
)

var listBranchCmd = &cobra.Command{
	Use:   "listBranch",
	Short: "Показать информацию о филиале по ID",
	Long: `
Команда выводит подробную информацию о филиале по его идентификатору.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Проверка наличия аргумента (ID филиала)
		if len(args) < 1 {
			return errors.New("необходимо указать ID филиала")
		}

		// Выполняем получение филиала по переданному ID
		branch, err := data.GetBranch(args[0]) // Аргумент передается первым элементом массива args
		if err != nil {
			return fmt.Errorf("ошибка при получении филиала: %w", err)
		}

		// Формирование вывода информации о филиале
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

		return nil
	},
}

func init() {
	RootCmd.AddCommand(listBranchCmd)
}

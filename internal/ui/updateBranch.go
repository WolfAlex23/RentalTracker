package ui

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/wolfalex23/rental-tracker/internal/data"
	"github.com/wolfalex23/rental-tracker/internal/model"
)

var updateCmd = &cobra.Command{
	Use:   "update",                     // Имя команды
	Short: "Обновляет данные о филиале", // Краткое описание
	Long: `
Эта команда позволяет обновить данные о филиале по указанным параметрам.
Используйте флаги для передачи новых значений полей.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		// Получаем значения аргументов
		idStr, _ := cmd.Flags().GetString("id")
		department, _ := cmd.Flags().GetString("department")
		address, _ := cmd.Flags().GetString("address")
		contract, _ := cmd.Flags().GetString("contract")
		ariaStr, _ := cmd.Flags().GetFloat64("aria")
		meterInYearStr, _ := cmd.Flags().GetFloat64("meter-in-year")
		totalInYearStr, _ := cmd.Flags().GetFloat64("total-in-year")

		// Преобразуем строку с ID в число
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return fmt.Errorf("невозможно преобразовать ID (%s)", idStr)
		}

		// Формируем объект Branch
		updatedBranch := &model.Branch{
			ID:          id,
			Department:  department,
			Address:     address,
			Contract:    contract,
			Aria:        ariaStr,
			MeterInYear: meterInYearStr,
			TotalInYear: totalInYearStr,
		}

		// Обновляем филиал
		err = data.UpdateBranch(updatedBranch)
		if err != nil {
			return err
		}

		fmt.Println("Филиал успешно обновлён.")
		return nil
	},
}

// Регистрация новой команды
func init() {
	RootCmd.AddCommand(updateCmd)

	// Устанавливаем обязательные и необязательные флаги
	updateCmd.Flags().StringP("id", "i", "", "Уникальный идентификатор филиала (обязательно)")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().StringP("department", "d", "", "Название филиала")
	updateCmd.Flags().StringP("address", "a", "", "Адрес филиала")
	updateCmd.Flags().StringP("contract", "c", "", "Контракт филиала")
	updateCmd.Flags().Float64P("aria", "A", 0, "Занимаемая площадь (опционально)")
	updateCmd.Flags().Float64P("meter-in-year", "m", 0, "Стоимость метра в год (опционально)")
	updateCmd.Flags().Float64P("total-in-year", "T", 0, "Общая стоимость в год (опционально)")
}

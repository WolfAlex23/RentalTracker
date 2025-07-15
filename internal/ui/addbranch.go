package ui

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wolfalex23/rental-tracker/internal/data"
	"github.com/wolfalex23/rental-tracker/internal/model"
)

var AddCmd = &cobra.Command{
	Use:   "add",                              // Имя команды
	Short: "Добовляет новые данные о филиале", // Краткое описание
	Long: `
Эта команда позволяет добавить данные о новом филиале.
Используйте флаги для передачи новых значений полей.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		// Получение значений аргументов

		department, _ := cmd.Flags().GetString("department")
		address, _ := cmd.Flags().GetString("address")
		contract, _ := cmd.Flags().GetString("contract")
		ariaStr, _ := cmd.Flags().GetFloat64("aria")
		meterInYearStr, _ := cmd.Flags().GetFloat64("meter-in-year")
		totalInYearStr, _ := cmd.Flags().GetFloat64("total-in-year")

		// Проверка обязательных аргументов
		if department == "" || address == "" || contract == "" ||
			ariaStr <= 0 || meterInYearStr <= 0 || totalInYearStr <= 0 {
			return errors.New("все обязательные поля должны быть заполнены положительными числами")
		}

		// Формирование объекта Branch
		newBranch := &model.Branch{
			Department:  department,
			Address:     address,
			Contract:    contract,
			Aria:        ariaStr,
			MeterInYear: meterInYearStr,
			TotalInYear: totalInYearStr,
		}

		// Добавление филиала
		err := data.AddBranch(newBranch)
		if err != nil {
			return fmt.Errorf("Ошибка при добавлении филиала: %w", err)
		}

		fmt.Println("Новый филиал успешно добавлен.")
		return nil
	},
}

// Регистрация новой команды
func init() {
	RootCmd.AddCommand(AddCmd)

	// Устанавка обязательных и необязательных флагов
	AddCmd.Flags().StringP("department", "d", "", "Наименование подразделения (обязательно)")
	AddCmd.MarkFlagRequired("department")
	AddCmd.Flags().StringP("address", "a", "", "Адрес филиала (обязательно)")
	AddCmd.MarkFlagRequired("address")
	AddCmd.Flags().StringP("contract", "c", "", "Контракт (обязательно)")
	AddCmd.MarkFlagRequired("contract")
	AddCmd.Flags().Float64P("aria", "A", 0.0, "Занимаемая площадь (обязательно)")
	AddCmd.MarkFlagRequired("aria")
	AddCmd.Flags().Float64P("meterInYear", "m", 0.0, "Стоимость метра в год (обязательно)")
	AddCmd.MarkFlagRequired("meterInYear")
	AddCmd.Flags().Float64P("totalInYear", "T", 0.0, "Общая сумма в год (обязательно)")
	AddCmd.MarkFlagRequired("totalInYear")
}

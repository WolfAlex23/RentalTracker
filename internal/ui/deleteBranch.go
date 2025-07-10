package ui

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wolfalex23/rental-tracker/internal/data"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Удаляет данные о филиале",
	Long:  "Эта команда полностью удаляет данные филиала по указанному ID.",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := data.DeleteBranch(args[0])
		if len(args) < 1 {
			return errors.New("необходимо указать ID филиала")
		}

		if err != nil {
			return err
		}

		fmt.Println("Данные удалены")

		return nil
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

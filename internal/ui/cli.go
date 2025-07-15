package ui

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "rental-tracker",
	Short: "Приложение для управления информацией об аренде филиалов",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Добро пожаловать в Rental Tracker!")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

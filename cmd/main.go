package main

import (
	"log"

	"github.com/wolfalex23/rental-tracker/internal/data"
	"github.com/wolfalex23/rental-tracker/internal/menu"
)

func main() {
	dbPath := "branches.db"

	err := data.Init(dbPath)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)

	}

	defer data.Close()
	menu.MenuLoop()
}

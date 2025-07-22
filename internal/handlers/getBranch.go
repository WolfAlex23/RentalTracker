package handlers

import (
	"github.com/wolfalex23/rental-tracker/internal/data"
)

func ListOneHandler() {

	id := readPositiveInt("Номер филиала: ")

	data.GetBranch(id)

}

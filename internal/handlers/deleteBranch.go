package handlers

import (
	"github.com/wolfalex23/rental-tracker/internal/data"
)

func DeleteHandler() {

	id := readPositiveInt("Номер филиала: ")

	data.DeleteBranch(id)

}

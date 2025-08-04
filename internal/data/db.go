package data

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/wolfalex23/rental-tracker/internal/model"
	_ "modernc.org/sqlite"
)

var db *sql.DB

const schema = `
CREATE TABLE IF NOT EXISTS branches (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    department TEXT NOT NULL DEFAULT '',
    address TEXT NOT NULL DEFAULT '',
    contract TEXT NOT NULL DEFAULT '',
    aria DECIMAL(10, 2) NOT NULL DEFAULT '0.00',
    meterInYear DECIMAL(10, 2) NOT NULL DEFAULT '0.00',
    totalInYear DECIMAL(10, 2) NOT NULL DEFAULT '0.00',
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX branches_department ON branches(department);

CREATE TRIGGER trigger_update_branch_updatedAt
   AFTER UPDATE ON branches
   FOR EACH ROW BEGIN
       UPDATE branches SET updatedAt=date('now') WHERE id=NEW.id;
   END;
`

func Init(dbFile string) error {

	_, err := os.Stat(dbFile)

	var install bool
	if os.IsNotExist(err) {
		install = true
	} else if err != nil {
		// Другая ошибка, возможно проблема с правами доступа

		return fmt.Errorf("DB status check failed: %v", err)
	}

	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("failed to open DB: %v", err)
	}

	if install {
		_, err = db.Exec(schema)
		if err != nil {
			return fmt.Errorf("failed to execute schema: %v", err)
		}
		fmt.Println("new DB creation success")
	}

	return nil
}

func Close() error {
	return db.Close()
}

func SetLastUpdated(branch *model.Branch) {
	branch.UpdatedAt = time.Now().UTC()
}

func AddBranch(branch *model.Branch) error {

	query := `INSERT INTO branches (department, address, contract, aria, meterInYear, totalInYear) VALUES (:department, :address, :contract, :aria, :meterInYear, :totalInYear)`
	_, err := db.Exec(query,
		sql.Named("department", branch.Department),
		sql.Named("address", branch.Address),
		sql.Named("contract", branch.Contract),
		sql.Named("aria", branch.Aria),
		sql.Named("meterInYear", branch.MeterInYear),
		sql.Named("totalInYear", branch.TotalInYear))

	return err

}

func GetBranches() ([]*model.Branch, error) {

	var query string

	query = "SELECT id, department, address, contract, aria, meterInYear, totalInYear, updatedAt FROM branches ORDER BY updatedAt"

	branches := make([]*model.Branch, 0)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed SELECT-request: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		branch := model.Branch{}

		err := rows.Scan(&branch.ID, &branch.Department, &branch.Address, &branch.Contract, &branch.Aria, &branch.MeterInYear, &branch.TotalInYear, &branch.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %v", err)
		}

		branches = append(branches, &branch)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return branches, nil
}

func GetBranch(id int) (*model.Branch, error) {

	branch := &model.Branch{}

	err := db.QueryRow("SELECT id, department, address, contract, aria, meterInYear, totalInYear, updatedAt FROM branches WHERE id = :id",
		sql.Named("id", id)).Scan(&branch.ID, &branch.Department, &branch.Address, &branch.Contract, &branch.Aria, &branch.MeterInYear, &branch.TotalInYear, &branch.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("branch not found")
		}
		return nil, fmt.Errorf("row scan failed: %v", err)
	}
	return branch, nil
}

func UpdateBranch(branch *model.Branch) error {
	query := `UPDATE branches SET department = :department, address = :address, contract = :contract, aria = :aria, meterInYear = :meterInYear, totalInYear = :totalInYear WHERE id = :id`
	res, err := db.Exec(query,
		sql.Named("department", branch.Department),
		sql.Named("address", branch.Address),
		sql.Named("contract", branch.Contract),
		sql.Named("aria", branch.Aria),
		sql.Named("meterInYear", branch.MeterInYear),
		sql.Named("totalInYear", branch.TotalInYear),
		sql.Named("id", branch.ID))
	if err != nil {
		return fmt.Errorf("branch update failed: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating branch`)
	}
	return nil
}

func DeleteBranch(id int) error {

	res, err := db.Exec("DELETE FROM branches WHERE id = :id", sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("branch delete failed: %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`no such id branch to delete`)
	}
	return nil
}

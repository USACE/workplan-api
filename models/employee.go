package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID                     uuid.UUID `json:"id"`
	Name                   string    `json:"name"`
	Rate                   float32   `json:"rate"`
	AvailabilityMultiplier float32   `json:"availability_multiplier" db:"availability_multiplier"`
}

func ListEmployees(db *sqlx.DB) ([]Employee, error) {
	// Fetch employee commitments summaries by month
	ee := make([]Employee, 0)
	if err := db.Select(
		&ee,
		"SELECT id, name, rate, availability_multiplier FROM employee",
	); err != nil {
		return make([]Employee, 0), err
	}
	return ee, nil
}

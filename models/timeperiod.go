package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Timeperiod struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	End      time.Time `json:"timeperiod_end" db:"timeperiod_end"`
	Workdays int       `json:"workdays"`
}

func ListTimeperiods(db *sqlx.DB) ([]Timeperiod, error) {
	tt := make([]Timeperiod, 0)
	if err := db.Select(&tt, `SELECT * FROM timeperiod`); err != nil {
		return make([]Timeperiod, 0), err
	}
	return tt, nil
}

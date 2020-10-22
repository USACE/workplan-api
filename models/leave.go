package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Leave is a number of days for a single employee, timeperiod, and project
type Leave struct {
	ID             uuid.UUID `json:"id"`
	EmployeeID     uuid.UUID `json:"employee_id" db:"employee_id"`
	EmployeeName   string    `json:"employee_name" db:"employee_name"`
	TimeperiodID   uuid.UUID `json:"timeperiod_id" db:"timeperiod_id"`
	TimeperiodName string    `json:"timeperiod_name" db:"timeperiod_name"`
	Days           int       `json:"days"`
}

func GetLeave(db *sqlx.DB, ID *uuid.UUID) (*Leave, error) {
	var e Leave
	if err := db.Get(&e, `SELECT * FROM v_leave WHERE id=$1`, ID); err != nil {
		return nil, err
	}
	return &e, nil
}

func ListLeave(db *sqlx.DB) ([]Leave, error) {

	ee := make([]Leave, 0)
	if err := db.Select(
		&ee,
		`SELECT id,
	            employee_id,
	            employee_name,
	            timeperiod_id,
	            timeperiod_name,
				days
         FROM v_leave
         ORDER BY employee_id, timeperiod_id
    `); err != nil {
		return make([]Leave, 0), err
	}
	return ee, nil
}

func CreateLeave(db *sqlx.DB, leave *Leave) (*Leave, error) {
	var id uuid.UUID
	if err := db.Get(
		&id,
		`INSERT INTO leave (employee_id, timeperiod_id, days) VALUES
			($1,$2,$3)
		 ON CONFLICT ON CONSTRAINT employee_timeperiod_unique DO UPDATE SET days = EXCLUDED.days
		 RETURNING ID`,
		leave.EmployeeID, leave.TimeperiodID, leave.Days,
	); err != nil {
		return nil, err
	}
	return GetLeave(db, &id)
}

// DeleteLeave deletes a single leave using the leave ID
func DeleteLeave(db *sqlx.DB, ID *uuid.UUID) error {
	if _, err := db.Exec("DELETE FROM leave WHERE id = $1", ID); err != nil {
		return err
	}
	return nil
}

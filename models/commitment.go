package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Commitment is a number of days for a single employee, timeperiod, and project
type Commitment struct {
	ID             uuid.UUID `json:"id"`
	EmployeeID     uuid.UUID `json:"employee_id" db:"employee_id"`
	EmployeeName   string    `json:"employee_name" db:"employee_name"`
	TimeperiodID   uuid.UUID `json:"timeperiod_id" db:"timeperiod_id"`
	TimeperiodName string    `json:"timeperiod_name" db:"timeperiod_name"`
	ProjectID      uuid.UUID `json:"project_id" db:"project_id"`
	ProjectName    string    `json:"project_name" db:"project_name"`
	Days           int       `json:"days"`
	Cost           float32   `json:"cost"`
}

func GetCommitment(db *sqlx.DB, ID *uuid.UUID) (*Commitment, error) {
	var c Commitment
	if err := db.Get(&c, `SELECT * FROM v_commitment WHERE id=$1`, ID); err != nil {
		return nil, err
	}
	return &c, nil
}

func ListCommitments(db *sqlx.DB) ([]Commitment, error) {

	cc := make([]Commitment, 0)
	if err := db.Select(
		&cc,
		`SELECT id,
	            employee_id,
	            employee_name,
	            timeperiod_id,
	            timeperiod_name,
	            project_id,
	            project_name,
				days,
				cost
         FROM v_commitment
         ORDER BY employee_id, timeperiod_id, project_id
    `); err != nil {
		return make([]Commitment, 0), err
	}
	return cc, nil
}

func CreateCommitment(db *sqlx.DB, commitment *Commitment) (*Commitment, error) {
	var id uuid.UUID
	if err := db.Get(
		&id,
		`INSERT INTO commitment (employee_id, timeperiod_id, project_id, days) VALUES
			($1,$2,$3,$4)
		 ON CONFLICT ON CONSTRAINT employee_timeperiod_project_unique DO UPDATE SET days = EXCLUDED.days
		 RETURNING ID`,
		commitment.EmployeeID, commitment.TimeperiodID, commitment.ProjectID, commitment.Days,
	); err != nil {
		return nil, err
	}
	return GetCommitment(db, &id)
}

// DeleteCommitment deletes a single commitment using the commitment ID
func DeleteCommitment(db *sqlx.DB, ID *uuid.UUID) error {
	if _, err := db.Exec("DELETE FROM commitment WHERE id = $1", ID); err != nil {
		return err
	}
	return nil
}

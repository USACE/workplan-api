package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Project struct {
	ID                 uuid.UUID  `json:"id"`
	Name               string     `json:"name"`
	Funding            *float32   `json:"funding,omitempty"`
	FundsRemaining     *float32   `json:"funds_remaining" db:"funds_remaining"`
	LatestRealityCheck *time.Time `json:"latest_reality_check" db:"latest_reality_check"`
	FeedbackEnabled    *bool      `json:"feedback_enabled,omitempty" db:"feedback_enabled"`
}

func ListProjects(db *sqlx.DB) ([]Project, error) {
	pp := make([]Project, 0)
	if err := db.Select(&pp, `SELECT id, name, funding, funds_remaining, latest_reality_check, feedback_enabled FROM project`); err != nil {
		return make([]Project, 0), err
	}
	return pp, nil
}

func ListFeedbackProjects(db *sqlx.DB) ([]Project, error) {
	pp := make([]Project, 0)
	if err := db.Select(&pp, `SELECT id, name FROM project WHERE feedback_enabled`); err != nil {
		return make([]Project, 0), err
	}
	return pp, nil
}

func CreateProject(db *sqlx.DB, p *Project) (*Project, error) {
	var pNew Project
	sql := "INSERT INTO project (name, funding) VALUES ($1, $2) RETURNING id, name, funding"
	if err := db.Get(&pNew, sql, p.Name, p.Funding); err != nil {
		return nil, err
	}
	return &pNew, nil
}

func UpdateProject(db *sqlx.DB, p *Project) (*Project, error) {
	var pUpdated Project
	sql := "UPDATE project SET name=$2, funding=$3, feedback_enabled=$4 WHERE id=$1 RETURNING id, name, funding, feedback_enabled"
	if err := db.Get(&pUpdated, sql, p.ID, p.Name, p.Funding, p.FeedbackEnabled); err != nil {
		return nil, err
	}
	return &pUpdated, nil
}

func DeleteProject(db *sqlx.DB, projectID *uuid.UUID) error {

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM commitment WHERE project_id=$1", projectID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	_, err = tx.Exec("DELETE FROM project WHERE id=$1", projectID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

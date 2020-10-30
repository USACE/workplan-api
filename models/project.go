package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Project struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Funding float32   `json:"funding"`
}

func ListProjects(db *sqlx.DB) ([]Project, error) {
	pp := make([]Project, 0)
	if err := db.Select(&pp, `SELECT id, name, funding FROM project`); err != nil {
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
	sql := "UPDATE project SET name=$2, funding=$3 WHERE id=$1"
	if err := db.Get(&pUpdated, sql, p.ID, p.Name, p.Funding); err != nil {
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

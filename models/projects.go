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

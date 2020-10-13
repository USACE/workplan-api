package handlers

import (
	"net/http"

	"github.com/USACE/workplan-api/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func ListCommitments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc, err := models.ListCommitments(db)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, &cc)
	}
}

func CreateCommitment(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var commitment models.Commitment
		c.Bind(&commitment)
		cNew, err := models.CreateCommitment(db, &commitment)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusCreated, &cNew)

	}
}

// DeleteCommitment calls models.DeleteCommitment using the ID from the request
func DeleteCommitment(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		err = models.DeleteCommitment(db, &id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

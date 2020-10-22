package handlers

import (
	"net/http"

	"github.com/USACE/workplan-api/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func ListLeave(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ee, err := models.ListLeave(db)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, &ee)
	}
}

func CreateLeave(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var leave models.Leave
		c.Bind(&leave)
		eNew, err := models.CreateLeave(db, &leave)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusCreated, &eNew)

	}
}

// DeleteLeave deletes leave using the ID from the request
func DeleteLeave(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		err = models.DeleteLeave(db, &id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

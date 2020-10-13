package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/USACE/workplan-api/models"

	"github.com/labstack/echo/v4"
)

func ListEmployees(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ee, err := models.ListEmployees(db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, ee)
	}
}

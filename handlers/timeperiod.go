package handlers

import (
	"net/http"

	"workplan-api/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func ListTimeperiods(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		tt, err := models.ListTimeperiods(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, tt)
	}
}

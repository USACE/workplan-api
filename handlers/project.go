package handlers

import (
	"fmt"
	"net/http"

	"github.com/USACE/workplan-api/models"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func ListFeedbackProjects(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pp, err := models.ListFeedbackProjects(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, pp)
	}
}

func ListProjects(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		pp, err := models.ListProjects(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, pp)
	}
}

func CreateProject(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var p models.Project
		if err := c.Bind(&p); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		pNew, err := models.CreateProject(db, &p)
		if err != nil {
			fmt.Println(err.Error())
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, &pNew)
	}
}

func UpdateProject(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var p models.Project
		if err := c.Bind(&p); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		pUpdated, err := models.UpdateProject(db, &p)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, &pUpdated)
	}
}

func DeleteProject(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		err = models.DeleteProject(db, &id)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

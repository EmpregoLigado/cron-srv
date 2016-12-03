package handlers

import (
	"github.com/labstack/echo"
	"github.com/rafaeljesus/cron-srv/models"
	"net/http"
	"strconv"
)

func (env *Env) CronIndex(c echo.Context) error {
	status := c.QueryParam("status")
	expression := c.QueryParam("expression")
	query := models.Query{status, expression}

	crons := []models.Cron{}
	if err := env.Repo.Search(&query, &crons).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, crons)
}

func (env *Env) CronCreate(c echo.Context) error {
	cron := models.Cron{}
	if err := c.Bind(&cron); err != nil {
		return err
	}

	if err := env.Repo.CreateCron(&cron).Error; err != nil {
		return err
	}

	if err := env.Sched.Create(&cron); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, cron)
}

func (env *Env) CronShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cron := models.Cron{}
	if err := env.Repo.FindCronById(&cron, id).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cron)
}

func (env *Env) CronUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cron := models.Cron{}
	if err := env.Repo.FindCronById(&cron, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	cr := models.Cron{}
	if err := c.Bind(&cr); err != nil {
		return err
	}

	cron.Status = cr.Status
	cron.Expression = cr.Expression
	cron.Url = cr.Url

	if err := env.Repo.UpdateCron(&cron).Error; err != nil {
		return err
	}

	if err := env.Sched.Update(&cron); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cron)
}

func (env *Env) CronDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cron := models.Cron{}
	if err := env.Repo.FindCronById(&cron, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	if err := env.Repo.DeleteCron(&cron).Error; err != nil {
		return err
	}

	if err := env.Sched.Delete(cron.Id); err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, nil)
}

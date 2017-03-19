package api

import (
	"encoding/json"
	"github.com/EmpregoLigado/cron-srv/models"
	"github.com/nbari/violetear"
	"log"
	"net/http"
	"strconv"
)

func (h *APIHandler) EventsIndex(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var status, expression string

	if len(q["status"]) != 0 {
		status = q["status"][0]
	}

	if len(q["expression"]) != 0 {
		expression = q["expression"][0]
	}

	query := models.Query{status, expression}
	crons := []models.Cron{}
	if err := h.Repo.Search(&query, &crons); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&crons)
}

func (h *APIHandler) EventsCreate(w http.ResponseWriter, r *http.Request) {
	cron := new(models.Cron)
	if err := json.NewDecoder(r.Body).Decode(cron); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateCron(cron); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := h.Scheduler.Create(cron); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cron)
}

func (h *APIHandler) EventsShow(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cron := new(models.Cron)
	if err := h.Repo.FindCronById(cron, id); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(cron)
}

func (h *APIHandler) EventsUpdate(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cron := new(models.Cron)
	if err := h.Repo.FindCronById(cron, id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cr := new(models.Cron)
	if err := json.NewDecoder(r.Body).Decode(cr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cron.Status = cr.Status
	cron.Expression = cr.Expression
	cron.Url = cr.Url
	cron.MaxRetries = cr.MaxRetries
	cron.RetryTimeout = cr.RetryTimeout

	if err := h.Repo.UpdateCron(cron); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := h.Scheduler.Update(cron); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cron)
}

func (h *APIHandler) EventsDelete(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].([]string)[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cron := new(models.Cron)
	if err := h.Repo.FindCronById(cron, id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := h.Repo.DeleteCron(cron); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := h.Scheduler.Delete(cron.Id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

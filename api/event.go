package api

import (
	"encoding/json"
	"github.com/EmpregoLigado/cron-srv/models"
	"github.com/nbari/violetear"
	"net/http"
	"strconv"
)

func (h *APIHandler) EventsIndex(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	expression := r.URL.Query().Get("expression")
	query := models.NewQuery(status, expression)
	events := []models.Event{}

	if err := h.Repo.FindEvents(&events, query); err != nil {
		JSON(w, http.StatusPreconditionFailed, err)
		return
	}

	JSON(w, http.StatusOK, &events)
}

func (h *APIHandler) EventsCreate(w http.ResponseWriter, r *http.Request) {
	event := new(models.Event)
	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		JSON(w, http.StatusBadRequest, err)
		return
	}

	if err := h.Repo.CreateEvent(event); err != nil {
		JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := h.Scheduler.Create(event); err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusCreated, event)
}

func (h *APIHandler) EventsShow(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].(string))
	if err != nil {
		JSON(w, http.StatusBadRequest, err)
		return
	}

	event := new(models.Event)
	if err := h.Repo.FindEventById(event, id); err != nil {
		JSON(w, http.StatusNotFound, err)
		return
	}

	JSON(w, http.StatusOK, event)
}

func (h *APIHandler) EventsUpdate(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].(string))
	if err != nil {
		JSON(w, http.StatusBadRequest, err)
		return
	}

	event := new(models.Event)
	if err := h.Repo.FindEventById(event, id); err != nil {
		JSON(w, http.StatusNotFound, err)
		return
	}

	e := new(models.Event)
	if err := json.NewDecoder(r.Body).Decode(e); err != nil {
		JSON(w, http.StatusBadRequest, err)
		return
	}

	event.Status = e.Status
	event.Expression = e.Expression
	event.Url = e.Url
	event.Retries = e.Retries
	event.Timeout = e.Timeout

	if err := h.Repo.UpdateEvent(event); err != nil {
		JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := h.Scheduler.Update(event); err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, event)
}

func (h *APIHandler) EventsDelete(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].([]string)[0])
	if err != nil {
		JSON(w, http.StatusBadRequest, err)
		return
	}

	event := new(models.Event)
	if err := h.Repo.FindEventById(event, id); err != nil {
		JSON(w, http.StatusNotFound, err)
		return
	}

	if err := h.Repo.DeleteEvent(event); err != nil {
		JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := h.Scheduler.Delete(event.Id); err != nil {
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusNoContent, nil)
}

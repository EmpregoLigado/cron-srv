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

	query := models.NewQuery(status, expression)
	events := []models.Event{}
	if err := h.Repo.FindEvents(&events, query); err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&events)
}

func (h *APIHandler) EventsCreate(w http.ResponseWriter, r *http.Request) {
	event := new(models.Event)
	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateEvent(event); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := h.Scheduler.Create(event); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func (h *APIHandler) EventsShow(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event := new(models.Event)
	if err := h.Repo.FindEventById(event, id); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(event)
}

func (h *APIHandler) EventsUpdate(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event := new(models.Event)
	if err := h.Repo.FindEventById(event, id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	e := new(models.Event)
	if err := json.NewDecoder(r.Body).Decode(e); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event.Status = e.Status
	event.Expression = e.Expression
	event.Url = e.Url
	event.Retries = e.Retries
	event.Timeout = e.Timeout

	if err := h.Repo.UpdateEvent(event); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := h.Scheduler.Update(event); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

func (h *APIHandler) EventsDelete(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].([]string)[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event := new(models.Event)
	if err := h.Repo.FindEventById(event, id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := h.Repo.DeleteEvent(event); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := h.Scheduler.Delete(event.Id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

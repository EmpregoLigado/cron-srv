package mock

import "github.com/EmpregoLigado/cron-srv/models"

type RepoMock struct {
	Created      bool
	Updated      bool
	Deleted      bool
	Found        bool
	Searched     bool
	ByStatus     bool
	ByExpression bool
}

func NewRepo() *RepoMock {
	return &RepoMock{
		Created:      false,
		Updated:      false,
		Deleted:      false,
		Found:        false,
		Searched:     false,
		ByStatus:     false,
		ByExpression: false,
	}
}

func (repo *RepoMock) FindEvents(events *[]models.Event, sc *models.Query) (err error) {
	*events = append(*events, models.Event{Expression: "* * * * * *"})
	switch true {
	case sc.Status != "":
		repo.ByStatus = true
	case sc.Expression != "":
		repo.ByExpression = true
	default:
		repo.Searched = true
	}
	return
}

func (repo *RepoMock) FindEventById(event *models.Event, id int) (err error) {
	repo.Found = true
	return
}

func (repo *RepoMock) CreateEvent(event *models.Event) (err error) {
	repo.Created = true
	return
}

func (repo *RepoMock) UpdateEvent(event *models.Event) (err error) {
	repo.Updated = true
	return
}

func (repo *RepoMock) DeleteEvent(event *models.Event) (err error) {
	repo.Deleted = true
	return
}

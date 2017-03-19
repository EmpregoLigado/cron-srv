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

func (repo *RepoMock) Search(sc *models.Query, crons *[]models.Cron) (err error) {
	*crons = append(*crons, models.Cron{Expression: "* * * * * *"})
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

func (repo *RepoMock) FindCronById(cron *models.Cron, id int) (err error) {
	repo.Found = true
	return
}

func (repo *RepoMock) CreateCron(cron *models.Cron) (err error) {
	repo.Created = true
	return
}

func (repo *RepoMock) UpdateCron(cron *models.Cron) (err error) {
	repo.Updated = true
	return
}

func (repo *RepoMock) DeleteCron(cron *models.Cron) (err error) {
	repo.Deleted = true
	return
}

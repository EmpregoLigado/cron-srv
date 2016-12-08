package models

type RepoMock struct{}

func (repo *RepoMock) Search(*Query, *[]Cron) error {
	crons := make([]*Cron, 0)
	crons = append(crons, &Cron{Expression: "* * * * * *"})
	return nil
}

func (repo *RepoMock) FindCronById(cron *Cron, id int) error {
	return nil
}

func (repo *RepoMock) CreateCron(cron *Cron) error {
	return nil
}

func (repo *RepoMock) UpdateCron(cron *Cron) error {
	return nil
}

func (repo *RepoMock) DeleteCron(cron *Cron) error {
	return nil
}

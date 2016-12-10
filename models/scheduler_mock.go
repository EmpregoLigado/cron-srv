package models

type SchedMock struct{}

func (sm *SchedMock) Create(cron *Cron) error {
	return nil
}

func (sm *SchedMock) Update(cron *Cron) error {
	return nil
}

func (sm *SchedMock) Delete(id uint) error {
	return nil
}

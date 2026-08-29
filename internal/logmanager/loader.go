package logmanager

import (
	"github.com/juiposa/Automitic/internal/model"
)

func (l *logManager) LoadLogs(reportId string, fightNumber int) (model.Fight, error) {
	err := l.cache.InitCache(reportId, fightNumber)
	if err != nil {
		return model.Fight{}, err
	}
	report, err := l.LoadReport(reportId, fightNumber)
	if err != nil {
		return model.Fight{}, err
	}
	fight, err := l.LoadFight(reportId, fightNumber, report)
	if err != nil {
		return model.Fight{}, err
	}
	return fight, nil
}

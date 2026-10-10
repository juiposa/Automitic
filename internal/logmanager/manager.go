package logmanager

import (
	"github.com/juiposa/Automitic/internal/config"
	"github.com/juiposa/Automitic/internal/data"
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/model"
)

const LogUrlTemplate = "https://www.fflogs.com/reports/%s"
const FightUrlTemplate = LogUrlTemplate + "?fight=%d"

type LogManager interface {
	LoadLogs(reportId string, fightNumber int) (model.Fight, error)
	LoadReport(reportId string, fightNumber int) (fflogs.Report, error)
	LoadFight(reportId string, fightNumber int, report fflogs.Report) (model.Fight, error)
}

type logManager struct {
	gql  fflogs.GQL
	data data.DataManager
}

func NewLogLoader(cfg *config.Config, gql fflogs.GQL, dataManager data.DataManager) LogManager {
	return &logManager{
		gql:  gql,
		data: dataManager,
	}
}

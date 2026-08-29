package logmanager

import (
	"github.com/juiposa/Automitic/internal/config"
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/model"
)

const LogUrlTemplate = "https://www.fflogs.com/reports/%s"
const FightUrlTemplate = LogUrlTemplate + "?fight=%d"

type LogManager interface {
	LoadLogs(reportId string, fightNumber int) (fflogs.Report, error)
	LoadReport(reportId string, fightNumber int) (fflogs.Report, error)
	LoadFight(reportId string, fightNumber int, report fflogs.Report) (model.Fight, error)

	SaveReport(reportId string, report fflogs.Report) error
	SaveFightEvents(reportId string, fightNumber int, fight model.Fight) error
}

type logManager struct {
	gql           fflogs.FFLogsGQL
	cacheLocation string
}

func NewLogLoader(cfg *config.Config, gql fflogs.FFLogsGQL) LogManager {
	return &logManager{
		gql:           gql,
		cacheLocation: cfg.FightCacheLocation,
	}
}

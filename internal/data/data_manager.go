package data

import (
	"os"

	"github.com/juiposa/Automitic/internal/config"
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/model"
	"github.com/juiposa/Automitic/internal/model/game"
)

type DataManager interface {

	// Report and Timeline
	InitCache(reportId string, fightNumber int) error
	SaveFightEvents(reportId string, fightNumber int, fight model.Fight) error
	SaveReport(reportId string, report fflogs.Report) error
	LoadFightEvents(reportId string, fightNumber int) (model.Fight, error)
	LoadReport(reportId string) (fflogs.Report, error)
	SaveResult(reportId string, fightNumber int, timeline model.Timeline) error

	// Game Data
	GetActions() (map[int64]game.Action, error)
}

type dataManager struct {
	cacheLocation string
	dataLocation  string
	actions       map[int64]game.Action
}

func InitDataManager(cfg *config.Config) DataManager {
	return &dataManager{
		cacheLocation: cfg.FightCacheLocation,
		dataLocation:  cfg.DataLocation,
	}
}

func (d *dataManager) saveToFile(filename string, fileContents []byte) error {
	return os.WriteFile(filename, fileContents, 0644)
}

func (d *dataManager) loadFromFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (d *dataManager) openFile(filename string) (*os.File, error) {
	return os.Open(filename)
}

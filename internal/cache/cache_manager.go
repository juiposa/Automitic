package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/juiposa/Automitic/internal/config"
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/model"
)

type CacheManager interface {
	InitCache(reportId string, fightNumber int) error
	SaveFightEvents(reportId string, fightNumber int, fight model.Fight) error
	SaveReport(reportId string, report fflogs.Report) error

	LoadFightEvents(reportId string, fightNumber int) (model.Fight, error)
	LoadReport(reportId string) (fflogs.Report, error)

	SaveResult(reportId string, fightNumber int, timeline model.Timeline) error
}

type cacheManager struct {
	cacheLocation string
}

const FightEventsFilename = "fight_events.json"
const ReportFilename = "report.json"

const ResultFileName = "result-%s-%s.json"

const DateFormat = "2006-01-02T15-04-05"

func NewCacheManager(cfg *config.Config) CacheManager {
	return &cacheManager{
		cacheLocation: cfg.FightCacheLocation,
	}
}

func (c *cacheManager) InitCache(reportId string, fightNumber int) error {
	cacheDir := c.getFightDir(reportId, fightNumber)
	err := os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (c *cacheManager) SaveReport(reportId string, report fflogs.Report) error {
	filename := c.getLogDir(reportId) + "/" + ReportFilename
	contents, err := json.Marshal(report)
	if err != nil {
		return err
	}
	return c.saveToFile(filename, contents)
}

func (c *cacheManager) SaveFightEvents(reportId string, fightNumber int, fight model.Fight) error {
	filename := c.getFightDir(reportId, fightNumber) + "/" + FightEventsFilename
	contents, err := json.MarshalIndent(fight, "", "  ")
	if err != nil {
		return err
	}
	return c.saveToFile(filename, contents)
}
func (c *cacheManager) LoadFightEvents(reportId string, fightNumber int) (model.Fight, error) {
	file, err := os.ReadFile(c.getFightDir(reportId, fightNumber) + "/" + FightEventsFilename)
	if err != nil {
		return model.Fight{}, err
	}
	fight := model.Fight{}
	err = json.Unmarshal(file, &fight)
	if err != nil {
		return model.Fight{}, err
	}
	return fight, nil
}

func (c *cacheManager) LoadReport(reportId string) (fflogs.Report, error) {
	file, err := os.ReadFile(c.getLogDir(reportId) + "/" + ReportFilename)
	if err != nil {
		return fflogs.Report{}, err
	}
	report := fflogs.Report{}
	err = json.Unmarshal(file, &report)
	if err != nil {
		return fflogs.Report{}, err
	}
	return report, nil
}

func (c *cacheManager) SaveResult(reportId string, fightNumber int, timeline model.Timeline) error {
	filename := c.getFightDir(reportId, fightNumber) + "/" + fmt.Sprintf(ResultFileName, reportId, time.Now().Format(DateFormat))
	contents, err := json.MarshalIndent(timeline, "", "  ")
	if err != nil {
		return err
	}
	return c.saveToFile(filename, contents)
}

func (c *cacheManager) getLogDir(reportId string) string {
	return fmt.Sprintf("%s/%s", c.cacheLocation, reportId)
}

func (c *cacheManager) getFightDir(reportId string, fightNumber int) string {
	return fmt.Sprintf("%s/fight%d", c.getLogDir(reportId), fightNumber)
}

func (c *cacheManager) saveToFile(filename string, fileContents []byte) error {
	return os.WriteFile(filename, fileContents, 0644)
}

func (c *cacheManager) loadFromFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

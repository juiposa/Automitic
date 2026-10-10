package data

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/model"
)

const FightEventsFilename = "fight_events.json"
const ReportFilename = "report.json"

const ResultFileName = "result-%s-%s.json"

const DateFormat = "2006-01-02T15-04-05"

func (d *dataManager) InitCache(reportId string, fightNumber int) error {
	cacheDir := d.getFightDir(reportId, fightNumber)
	err := os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (d *dataManager) SaveReport(reportId string, report fflogs.Report) error {
	filename := d.getLogDir(reportId) + "/" + ReportFilename
	contents, err := json.Marshal(report)
	if err != nil {
		return err
	}
	return d.saveToFile(filename, contents)
}

func (d *dataManager) SaveFightEvents(reportId string, fightNumber int, fight model.Fight) error {
	filename := d.getFightDir(reportId, fightNumber) + "/" + FightEventsFilename
	contents, err := json.MarshalIndent(fight, "", "  ")
	if err != nil {
		return err
	}
	return d.saveToFile(filename, contents)
}
func (d *dataManager) LoadFightEvents(reportId string, fightNumber int) (model.Fight, error) {
	file, err := os.ReadFile(d.getFightDir(reportId, fightNumber) + "/" + FightEventsFilename)
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

func (d *dataManager) LoadReport(reportId string) (fflogs.Report, error) {
	file, err := os.ReadFile(d.getLogDir(reportId) + "/" + ReportFilename)
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

func (d *dataManager) SaveResult(reportId string, fightNumber int, timeline model.Timeline) error {
	filename := d.getFightDir(reportId, fightNumber) + "/" + fmt.Sprintf(ResultFileName, reportId, time.Now().Format(DateFormat))
	contents, err := json.MarshalIndent(timeline, "", "  ")
	if err != nil {
		return err
	}
	return d.saveToFile(filename, contents)
}

func (d *dataManager) getLogDir(reportId string) string {
	return fmt.Sprintf("%s/%s", d.cacheLocation, reportId)
}

func (d *dataManager) getFightDir(reportId string, fightNumber int) string {
	return fmt.Sprintf("%s/fight%d", d.getLogDir(reportId), fightNumber)
}

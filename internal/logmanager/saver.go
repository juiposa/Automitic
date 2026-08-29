package logmanager

import (
	"encoding/json"
	"os"

	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/model"
)

func (l *logManager) SaveReport(reportId string, report fflogs.Report) error {
	filename := l.getLogDir(reportId) + "/" + ReportFilename
	contents, err := json.Marshal(report)
	if err != nil {
		return err
	}
	return l.saveToFile(filename, contents)
}

func (l *logManager) SaveFightEvents(reportId string, fightNumber int, fight model.Fight) error {
	filename := l.getFightDir(reportId, fightNumber) + "/" + FightEventsFilename
	contents, err := json.Marshal(fight)
	if err != nil {
		return err
	}
	return l.saveToFile(filename, contents)
}

func (l *logManager) saveToFile(filename string, fileContents []byte) error {
	return os.WriteFile(filename, fileContents, 0644)
}

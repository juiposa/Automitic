package logmanager

import (
	"encoding/json"
	"os"

	"github.com/juiposa/Automitic/internal/fflogs"
)

const ReportFilename = "report.json"

func (l *logManager) LoadReport(reportId string, fightNumber int) (fflogs.Report, error) {
	err := l.initCache(reportId, fightNumber)
	if err != nil {
		return fflogs.Report{}, err
	}
	report, err := l.tryLoadReportFromCache(reportId)
	if err == nil { //was found in files
		// TOOD break refresh into its own method
		if len(report.Fights) < fightNumber { //missing the fight, fetch a fresh list
			fights, err := l.gql.GetReportFights(reportId)
			if err != nil {
				return report, err
			}
			report.Fights = fights
		}
	} else {
		//can't load from file, fetch from FFLogs
		report, err = l.fetchFreshReport(reportId)
		if err != nil {
			return report, err
		}
	}
	err = l.SaveReport(reportId, report)
	if err != nil {
		return report, err
	}
	return report, err
}

func (l *logManager) tryLoadReportFromCache(reportId string) (fflogs.Report, error) {
	file, err := os.ReadFile(l.getLogDir(reportId) + "/" + ReportFilename)
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

func (l *logManager) fetchFreshReport(reportId string) (fflogs.Report, error) {
	fullReport := fflogs.Report{}

	// get fights list
	fights, err := l.gql.GetReportFights(reportId)
	if err != nil {
		return fflogs.Report{}, err
	}
	fullReport.Fights = fights
	return fullReport, nil
}

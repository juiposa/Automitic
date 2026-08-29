package logmanager

import (
	"github.com/juiposa/Automitic/internal/fflogs"
)

func (l *logManager) LoadReport(reportId string, fightNumber int) (fflogs.Report, error) {
	report, err := l.cache.LoadReport(reportId)
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
	err = l.cache.SaveReport(reportId, report)
	if err != nil {
		return report, err
	}
	return report, err
}

func (l *logManager) fetchFreshReport(reportId string) (fflogs.Report, error) {
	fullReport := fflogs.Report{}

	// get metadata
	metadata, err := l.gql.GetReportMetadata(reportId)
	if err != nil {
		return fflogs.Report{}, err
	}
	fullReport.MasterData = metadata

	// get fights list
	fights, err := l.gql.GetReportFights(reportId)
	if err != nil {
		return fflogs.Report{}, err
	}
	fullReport.Fights = fights
	return fullReport, nil
}

package logmanager

import (
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/model"
)

func (l *logManager) LoadFight(reportId string, fightNumber int, report fflogs.Report) (model.Fight, error) {
	fight, err := l.cache.LoadFightEvents(reportId, fightNumber)
	if err == nil { //was found in files
		return fight, nil
	} else {
		fightMetadata := report.Fights[fightNumber-1]
		//can't load from file, fetch from FFLogs
		fightEvents, err := l.gql.GetFightEvents(reportId, fightNumber, fightMetadata.StartTime, fightMetadata.EndTime)
		if err != nil {
			return model.Fight{}, err
		}
		returnFight := model.Fight{
			ReportFight: fightMetadata,
			Events:      fightEvents,
			Actors:      GetFightActors(report.MasterData.Actors, fightMetadata),
		}
		err = l.cache.SaveFightEvents(reportId, fightNumber, returnFight)
		if err != nil {
			return model.Fight{}, err
		}
		return returnFight, nil
	}
}

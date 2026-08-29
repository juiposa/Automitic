package logmanager

import (
	"encoding/json"
	"os"

	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/model"
)

const FightEventsFilename = "fight_events.json"

func (l *logManager) LoadFight(reportId string, fightNumber int, report fflogs.Report) (model.Fight, error) {
	fight, err := l.tryLoadFightFromCache(reportId, fightNumber)
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
		}
		err = l.SaveFightEvents(reportId, fightNumber, returnFight)
		if err != nil {
			return model.Fight{}, err
		}
		return returnFight, nil
	}
}

func (l *logManager) tryLoadFightFromCache(reportId string, fightNumber int) (model.Fight, error) {
	file, err := os.ReadFile(l.getFightDir(reportId, fightNumber) + "/" + FightEventsFilename)
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

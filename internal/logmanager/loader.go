package logmanager

import (
	"fmt"
	"os"

	"github.com/juiposa/Automitic/internal/model"
)

func (l *logManager) LoadLogs(reportId string, fightNumber int) (model.Fight, error) {
	report, err := l.LoadReport(reportId, fightNumber)
	if err != nil {
		return model.Fight{}, err
	}
	fight, err := l.LoadFight(reportId, fightNumber, report)
	if err != nil {
		return model.Fight{}, err
	}
	//typeList := map[string]bool{} //TODO
	//for _, event := range fight.Events {
	//	typeList[event.Type] = true
	//}
	//for key, _ := range typeList {
	//	println(key)
	//}
	return fight, nil
}

func (l *logManager) initCache(reportId string, fightNumber int) error {
	cacheDir := l.getFightDir(reportId, fightNumber)
	err := os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (l *logManager) getLogDir(reportId string) string {
	return fmt.Sprintf("%s/%s", l.cacheLocation, reportId)
}

func (l *logManager) getFightDir(reportId string, fightNumber int) string {
	return fmt.Sprintf("%s/fight%d", l.getLogDir(reportId), fightNumber)
}

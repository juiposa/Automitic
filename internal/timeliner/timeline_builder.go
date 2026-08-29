package timeliner

import (
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/logmanager"
	"github.com/juiposa/Automitic/internal/model"
)

/*
headmarker
removedebuff
targetabilityupdate
removedebuffstack
death
dispel
encounterend
removebuffstack
combatantinfo
calculateddamage
calculatedheal
limitbreakupdate
heal
applydebuff
tether
cast
damage
applybuffstack
refreshbuff
gaugeupdate
begincast
refreshdebuff
applydebuffstack
applybuff
removebuff
absorbed
*/

type Timeliner interface {
	BuildTimeline(reportId string, fightNumber int) (model.Timeline, error)
}

type timeliner struct {
	logManager logmanager.LogManager
}

func NewTimeliner(logManager logmanager.LogManager) Timeliner {
	return &timeliner{
		logManager: logManager,
	}
}

func (t *timeliner) BuildTimeline(reportId string, fightNumber int) (model.Timeline, error) {
	_, err := t.logManager.LoadLogs(reportId, fightNumber)
	if err != nil {
		return model.Timeline{}, err
	}
	return model.Timeline{}, nil
}

func CompactTimeline(events []fflogs.Event) map[int64][]fflogs.Event {
	timeline := make(map[int64][]fflogs.Event)
	for _, event := range events {
		timeline[event.Timestamp] = append(timeline[event.Timestamp], event)
	}
	return timeline
}

func MakeTimelineEvent(event fflogs.Event) model.TimelineEvent {
	//tl := model.TimelineEvent{
	//
	//}

	return model.TimelineEvent{}
}

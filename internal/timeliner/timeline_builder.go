package timeliner

import (
	"slices"

	"github.com/juiposa/Automitic/internal/config"
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
	logManager         logmanager.LogManager
	acceptedEventTypes []RawEventType
}

func NewTimeliner(cfg *config.Config, logManager logmanager.LogManager) Timeliner {
	return &timeliner{
		logManager: logManager,
		acceptedEventTypes: []RawEventType{
			BeginCast,
			Cast,
			CalculatedDamage,
			Damage,
		},
	}
}

func (t *timeliner) BuildTimeline(reportId string, fightNumber int) (model.Timeline, error) {
	fight, err := t.logManager.LoadLogs(reportId, fightNumber)
	if err != nil {
		return model.Timeline{}, err
	}
	events := t.FilterEvents(fight)
	rawTimeline := RawTimeline{
		StartTime: fight.StartTime,
		EndTime:   fight.EndTime,
		RawEvents: events,
	}
	rawTimeline.CompactTimeline()
	return model.Timeline{}, nil
}

func (t *timeliner) FilterEvents(fight model.Fight) []fflogs.Event {
	var returnList []fflogs.Event
	bossIDs := t.GetBossIDs(fight)
	for _, event := range fight.Events {
		if slices.Contains(t.acceptedEventTypes, RawEventType(event.Type)) &&
			slices.Contains(bossIDs, event.SourceID) {
			returnList = append(returnList, event)
		}
	}
	return returnList
}

func (t *timeliner) GetBossIDs(fight model.Fight) []int64 {
	var returnList []int64
	for _, bossActor := range fight.Actors.Bosses {
		returnList = append(returnList, bossActor.ID)
	}
	return returnList
}

package timeliner

import (
	"slices"

	"github.com/juiposa/Automitic/internal/config"
	"github.com/juiposa/Automitic/internal/data"
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/log"
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
	logManager           logmanager.LogManager
	dataManager          data.DataManager
	acceptedBossEvents   []fflogs.RawEventType
	acceptedPlayerEvents []fflogs.RawEventType
}

func NewTimeliner(cfg *config.Config, logManager logmanager.LogManager, dataManager data.DataManager) Timeliner {
	return &timeliner{
		logManager:  logManager,
		dataManager: dataManager,
		acceptedBossEvents: []fflogs.RawEventType{
			fflogs.BeginCast,
			fflogs.Cast,
			fflogs.CalculatedDamage,
			fflogs.Damage,
		},
		acceptedPlayerEvents: []fflogs.RawEventType{},
	}
}

func (t *timeliner) BuildTimeline(reportId string, fightNumber int) (model.Timeline, error) {
	fight, err := t.logManager.LoadLogs(reportId, fightNumber)
	if err != nil {
		return model.Timeline{}, err
	}
	actions, err := t.dataManager.GetActions()
	if err != nil {
		return model.Timeline{}, err
	}
	events := t.FilterEvents(fight)
	rawTimeline := RawTimeline{
		StartTime: fight.StartTime,
		EndTime:   fight.EndTime,
		RawEvents: events,
		Actors:    fight.Actors,
		Actions:   actions,
	}
	rawTimeline.CompactTimeline()
	timeline := model.Timeline{
		PullTime:    rawTimeline.GetTimestamp(fight.EndTime),
		Actors:      fight.Actors,
		FightNumber: fight.ID,
		ReportID:    reportId,
	}
	for len(rawTimeline.Sequence) > 0 {
		event := rawTimeline.PopTimelineEvent()
		if event != nil {
			timeline.Events = append(timeline.Events, event)
		}
	}
	return timeline, nil
}

func (t *timeliner) FilterEvents(fight model.Fight) []fflogs.Event {
	var returnList []fflogs.Event
	bossIDs, playerIDs := t.GetEntityIDs(fight)
	for _, event := range fight.Events {
		if t.IsAcceptedEvent(event, bossIDs, playerIDs) {
			returnList = append(returnList, event)
		}
	}
	return returnList
}

func (t *timeliner) IsAcceptedEvent(event fflogs.Event, bossIDs, playerIDs []int64) bool {
	result := (slices.Contains(t.acceptedBossEvents, event.Type) &&
		slices.Contains(bossIDs, event.SourceID)) ||
		(slices.Contains(t.acceptedPlayerEvents, event.Type) &&
			slices.Contains(playerIDs, event.SourceID))
	log.Debug("ISACCEPTED", "result", result, "source", event.SourceID, "type", event.Type, "boss", bossIDs, "player", playerIDs)
	return result
}

func (t *timeliner) GetEntityIDs(fight model.Fight) ([]int64, []int64) {
	var bossList []int64
	for _, bossActor := range fight.Actors.Bosses {
		bossList = append(bossList, bossActor.ID)
	}
	var playerList []int64
	for _, playerActor := range fight.Actors.Players {
		playerList = append(playerList, playerActor.ID)
	}
	return bossList, playerList
}

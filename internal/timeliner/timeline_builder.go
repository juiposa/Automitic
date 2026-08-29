package timeliner

import (
	"github.com/juiposa/Automitic/internal/fflogs"
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

func BuildTimeline() model.Timeline {
	//report := logloader.LoadLogs()
	//events := report.Events.Data
	//compactedTimeline := CompactTimeline(events)
	//println(compactedTimeline)
	return model.Timeline{}
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

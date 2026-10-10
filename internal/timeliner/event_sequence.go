package timeliner

import (
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/model"
)

// GetNextEventInSequence - gets the subsequent events for a cast, presumes the cast has already been popped off the top.
func (t *RawTimeline) GetNextEventInSequence(parentEvent fflogs.Event, sequenceSlice []float64, timelineEvent *model.TimelineEvent) []fflogs.Event {
	for index, timestamp := range sequenceSlice { // search ahead in sequence
		events := t.TimeseriesEvents[timestamp]
		if t.IsSameCast(parentEvent, events[0]) && t.IsSameTarget(parentEvent, events[0]) && fflogs.IsFollowingEventType(parentEvent.Type, events[0].Type) { // this is a timestamp that follows our cast
			for _, event := range events {
				t.FillInEvent(event, timelineEvent)
				t.GetNextEventInSequence(event, sequenceSlice[index:], timelineEvent)
			}
			break
		}
	}
	return nil
}

func (t *RawTimeline) IsSameTarget(parentEvent fflogs.Event, event fflogs.Event) bool {
	return parentEvent.TargetID == event.TargetID
}

func (t *RawTimeline) IsSameCast(parentEvent fflogs.Event, event fflogs.Event) bool {
	return parentEvent.AbilityGameID == event.AbilityGameID && parentEvent.SourceID == event.SourceID
}

func (t *RawTimeline) TimeInSeconds(time float64) float64 {
	return time / 1000
}

package timeliner

import (
	"fmt"
	"slices"
	"time"

	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/model"
	"github.com/juiposa/Automitic/internal/model/actors"
	"github.com/juiposa/Automitic/internal/model/game"
)

type RawTimeline struct {
	StartTime        float64
	EndTime          float64
	Sequence         []float64
	RawEvents        []fflogs.Event
	TimeseriesEvents map[float64][]fflogs.Event
	Actions          map[int64]game.Action
	Actors           actors.FightActors
}

/* Heuristics
A) Different RawEventTypes will never happen on the same exact timestamp, the game server processes each event sequentially.
B) Thusly we only need to check the first element of a timestamp's events to determine type.
C) Every `damage` or `heal` event for a given cast will always happen on the exact same timestamp.
D) Once a root event type (Begin, Cast, etc.) finds it's first related event, it should know how to determine it's ended.
   - If `damage`, once it finds a `calculateddamage` for each or a death, that completes the event.
   - If `heal`, once it finds a `calculatedheal` and/or `absorbed` for each or a death, that completes the event.
*/

// CompactTimeline - map events by timestamp, record the sequence of timestamps
func (t *RawTimeline) CompactTimeline() {
	compactedTimeline := make(map[float64][]fflogs.Event)
	currentTimestamp := float64(0)
	for _, event := range t.RawEvents {
		if event.Timestamp > currentTimestamp {
			currentTimestamp = event.Timestamp
			t.Sequence = append(t.Sequence, currentTimestamp)
		}
		compactedTimeline[event.Timestamp] = append(compactedTimeline[event.Timestamp], event)
	}
	t.TimeseriesEvents = compactedTimeline
}

// PopTimelineEvent - scan the timeseries and create a TimelineEvent, advance sequence
func (t *RawTimeline) PopTimelineEvent() *model.TimelineEvent {
	var currTimestamp float64
	currTimestamp, t.Sequence = t.Sequence[0], t.Sequence[1:]
	tEvents := t.TimeseriesEvents[currTimestamp]
	if len(tEvents) == 0 || tEvents[0].Processed || (!slices.Contains(fflogs.RootEventTypes, tEvents[0].Type)) { //check if time timestamp is a list of root events and hasn't already been processed
		return nil
	}
	timelineEvent := t.InitEvent(tEvents[0]) //init the event from the first element
	for _, event := range tEvents {
		if !event.Processed { //has been processed by a different sequence
			t.FillInEvent(event, &timelineEvent)
			t.GetNextEventInSequence(event, t.Sequence, &timelineEvent) // traverse each root event and scan for it's downstreams
		}
	}
	return &timelineEvent
}

func (t *RawTimeline) GetTimestamp(timestamp float64) string {
	timeDiff := timestamp - t.StartTime
	duration, _ := time.ParseDuration(fmt.Sprintf("%.0fms", timeDiff))
	minutes := duration / time.Minute
	duration -= minutes * time.Minute
	seconds := duration / time.Second
	duration -= seconds * time.Second
	milliseconds := duration / time.Millisecond
	return fmt.Sprintf("%02d:%02d.%03d", minutes, seconds, milliseconds)
}

package timeliner

import (
	"fmt"
	"time"

	"github.com/juiposa/Automitic/internal/fflogs"
)

type RawTimeline struct {
	StartTime             float64
	EndTime               float64
	Sequence              []float64
	CurrentTimestampIndex int
	RawEvents             []fflogs.Event
	TimeseriesEvents      map[float64][]fflogs.Event
}

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
	t.CurrentTimestampIndex = 0
	t.TimeseriesEvents = compactedTimeline
}

// PopTimelineEvent - scan the timeseries and create a TimelineEvent, removing the raw events processed from the timeseries
func (t *RawTimeline) PopTimelineEvent() []fflogs.Event {
	//tIndex := t.CurrentTimestampIndex
	return []fflogs.Event{}
}

func (t *RawTimeline) GetTimestamp(timestamp float64) string {
	timeDiff := timestamp - t.StartTime
	duration, _ := time.ParseDuration(fmt.Sprintf("%.0fms", timeDiff))
	minutes := duration / time.Minute
	duration -= minutes * time.Minute
	seconds := duration / time.Second
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

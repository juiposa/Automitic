package timeliner

import (
	"fmt"

	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/log"
	"github.com/juiposa/Automitic/internal/model"
)

func (t *RawTimeline) InitEvent(event fflogs.Event) model.TimelineEvent {
	return model.TimelineEvent{
		Time:            int64(event.Timestamp),
		Timestamp:       t.GetTimestamp(event.Timestamp),
		SourceID:        event.SourceID,
		ActionId:        event.AbilityGameID,
		ActionName:      t.GetActionName(event.AbilityGameID),
		DamageToTargets: make(map[int64]model.DamageToTarget),
	}
}

func (t *RawTimeline) FillInEvent(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	switch event.Type {
	case fflogs.BeginCast:
		t.FillInBeginCast(event, timelineEvent)
	case fflogs.Cast:
		t.FillInCast(event, timelineEvent)
	case fflogs.Damage:
		t.FillInDamage(event, timelineEvent)
	case fflogs.CalculatedDamage:
		t.FillInCalculatedDamage(event, timelineEvent)
	default:
		log.Debug("Skipping event: ", event.Type)
	}
	//event.Processed = true
}

func (t *RawTimeline) FillInBeginCast(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	timelineEvent.CastDuration = t.TimeInSeconds(float64(event.Duration))
}

func (t *RawTimeline) FillInCast(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	//NOOP
}

func (t *RawTimeline) FillInDamage(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	tl := timelineEvent.DamageToTargets[event.TargetID]
	tl.ApplicationDelta = t.TimeInSeconds(event.Timestamp - float64(timelineEvent.Time))
	tl.ApplicationTime = t.GetTimestamp(event.Timestamp)
	tl.FinalAmount = event.Amount
	tl.TotalAmountMitigated = event.Mitigated
	tl.TotalMultiplier = event.Multiplier
	timelineEvent.DamageToTargets[event.TargetID] = tl
}

func (t *RawTimeline) FillInCalculatedDamage(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	timelineEvent.SnapshotDelta = t.TimeInSeconds(event.Timestamp - float64(timelineEvent.Time))
	timelineEvent.SnapshotTime = t.GetTimestamp(event.Timestamp)

	target := t.Actors.Players[event.TargetID]
	dmg := model.DamageToTarget{
		TargetID:          target.ID,
		TargetName:        fmt.Sprintf("%s - %s", target.Job, target.Name),
		UnmitigatedAmount: event.UnmitigatedAmount,
	}
	timelineEvent.DamageToTargets[event.TargetID] = dmg
}

func (t *RawTimeline) FillInCalculatedHeal(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	//TODO
}

func (t *RawTimeline) FillInBuffApplication(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	//TODO
}

func (t *RawTimeline) GetActionName(id int64) string {
	name := t.Actions[id].Name
	if name == "" {
		name = "Auto Attack"
	}
	return name
}

// TODO so on so forth

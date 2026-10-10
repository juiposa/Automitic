package timeliner

import (
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/log"
	"github.com/juiposa/Automitic/internal/model"
)

func InitEvent(event fflogs.Event) model.TimelineEvent {
	return model.TimelineEvent{
		Time:            int64(event.Timestamp),
		SourceID:        event.SourceID,
		AbilityId:       event.AbilityGameID,
		DamageToTargets: make(map[int64]model.DamageToTarget),
	}
}

func FillInEvent(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	switch event.Type {
	case fflogs.Cast:
		FillInCast(event, timelineEvent)
	case fflogs.Damage:
		FillInDamage(event, timelineEvent)
	case fflogs.CalculatedDamage:
		FillInCalculatedDamage(event, timelineEvent)
	default:
		log.Debug("Skipping event: ", event.Type)
	}
}

func FillInCast(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	//NOOP, might not be needed
}

func FillInDamage(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	tl := timelineEvent.DamageToTargets[event.TargetID]
	tl.FinalAmount = event.Amount
	tl.TotalAmountMitigated = event.Mitigated
	tl.TotalMultiplier = event.Multiplier
	timelineEvent.DamageToTargets[event.TargetID] = tl
}

func FillInCalculatedDamage(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	dmg := model.DamageToTarget{
		TargetID:          event.TargetID,
		UnmitigatedAmount: event.UnmitigatedAmount,
	}
	timelineEvent.DamageToTargets[event.TargetID] = dmg
}

func FillInCalculatedHeal(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	//TODO
}

func FillInBuffApplication(event fflogs.Event, timelineEvent *model.TimelineEvent) {
	//TODO
}

// TODO so on so forth

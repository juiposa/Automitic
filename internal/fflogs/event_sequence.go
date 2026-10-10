package fflogs

import "slices"

var (
	RootEventTypes = []RawEventType{
		Cast,
	}
	EventSequence = map[RawEventType][]RawEventType{
		BeginCast:        {Cast},
		Cast:             {CalculatedDamage, CalculatedHeal, ApplyBuff, ApplyBuff},
		CalculatedDamage: {Damage},
		CalculatedHeal:   {Heal},
		Damage:           {},
		Heal:             {},
	}
)

func IsFollowingEventType(eventType RawEventType, nextType RawEventType) bool {
	return slices.Contains(EventSequence[eventType], nextType)
}

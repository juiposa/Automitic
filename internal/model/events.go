package model

import "github.com/juiposa/Automitic/internal/model/actors"

type EventType string

const (
	DAMAGE     EventType = "DAMAGE"
	MITIGATION EventType = "MITIGATION"
)

type Event struct {
	PrePlayerState []actors.PlayerState
}

type DamageEvent struct {
	AbilityId         int
	UnmitigatedAmount int
	Crit              bool // hitType == 2

	MitPresent           []Mitigation
	TotalMultiplier      float64
	TotalAmountMitigated int

	FinalAmount int
}

type MitigationType string

const (
	HEALING MitigationType = "HEALING"
	SHIELD  MitigationType = "SHIELD"
	MIT     MitigationType = "MIT"
)

type Mitigation struct {
	//abilityGameID is the buff
	//extraAbilityGameID is the action that applies it
	AbiltiyId int
	Type      MitigationType
	Amount    float64
}

package model

import (
	"time"

	"github.com/juiposa/Automitic/internal/model/actors"
)

type Timeline struct {
	Timestamp        time.Time                `json:"timestamp"`
	PullTime         int                      `json:"pull_time"`
	Enemies          []actors.Enemy           `json:"enemies"`
	PlayerCharacters []actors.PlayerCharacter `json:"player_characters"`
	FightNumber      int                      `json:"fight_number"`
	ReportID         string                   `json:"report_id"`
	Events           []TimelineEvent          `json:"events"`
}

type TimelineEvent struct {
	Time int64 `json:"time"`
	Type EventType
	DamageEvent
}

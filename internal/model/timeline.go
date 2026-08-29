package model

import (
	"time"
)

type Timeline struct {
	Timestamp        time.Time         `json:"timestamp"`
	PullTime         int               `json:"pull_time"`
	Enemies          []Enemy           `json:"enemies"`
	PlayerCharacters []PlayerCharacter `json:"player_characters"`
	FightNumber      int               `json:"fight_number"`
	ReportID         string            `json:"report_id"`
	Events           []TimelineEvent   `json:"events"`
}

type TimelineEvent struct {
	Time int64 `json:"time"`
	Type EventType
	DamageEvent

}

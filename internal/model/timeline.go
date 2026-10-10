package model

import (
	"time"

	"github.com/juiposa/Automitic/internal/model/actors"
)

type Timeline struct {
	Timestamp   time.Time          `json:"timestamp"`
	PullTime    time.Duration      `json:"pull_time"`
	Actors      actors.FightActors `json:"fight_actors"`
	FightNumber int64              `json:"fight_number"`
	ReportID    string             `json:"report_id"`
	Events      []*TimelineEvent   `json:"events"`
}

type TimelineEvent struct {
	Time            int64                    `json:"time"`
	Timestamp       string                   `json:"timestamp"`
	SourceID        int64                    `json:"source_id"`
	AbilityId       int64                    `json:"ability_id"`
	DamageToTargets map[int64]DamageToTarget `json:"damage_to_targets"`
}

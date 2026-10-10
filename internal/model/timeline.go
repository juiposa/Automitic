package model

import (
	"time"

	"github.com/juiposa/Automitic/internal/model/actors"
)

type Timeline struct {
	Timestamp   time.Time          `json:"timestamp"`
	PullTime    string             `json:"pull_time"`
	Actors      actors.FightActors `json:"fight_actors"`
	FightNumber int64              `json:"fight_number"`
	ReportID    string             `json:"report_id"`
	Events      []*TimelineEvent   `json:"events"`
}

type TimelineEvent struct {
	Time       int64  `json:"time"`
	Timestamp  string `json:"timestamp"`
	SourceID   int64  `json:"source_id"`
	SourceName string `json:"source_name"`

	ActionId   int64  `json:"action_id"`
	ActionName string `json:"action_name"`

	CastDuration  float64 `json:"cast_duration"`
	SnapshotDelta float64 `json:"snapshot_delta"`
	SnapshotTime  string  `json:"snapshot_time"`

	DamageToTargets map[int64]DamageToTarget `json:"damage_to_targets"`
}

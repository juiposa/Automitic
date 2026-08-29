package model

import (
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/model/actors"
)

type Fight struct {
	*fflogs.ReportFight
	Events []fflogs.Event     `json:"events"`
	Actors actors.FightActors `json:"actors"`
}

package model

import "github.com/juiposa/Automitic/internal/fflogs"

type Fight struct {
	*fflogs.ReportFight
	Events []fflogs.Event `json:"events"`
}

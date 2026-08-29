package jobs

import "github.com/juiposa/Automitic/internal/model/actors"

func GetJobType(fflogsActorSubType string) actors.Job {
	switch fflogsActorSubType {
	case "Warrior":
		return actors.WAR
	case "Paladin":
		return actors.PLD
	case "DarkKnight":
		return actors.DRK
	case "Gunbreaker":
		return actors.GNB
	case "WhiteMage":
		return actors.WHM
	case "Scholar":
		return actors.SCH
	case "Astrologian":
		return actors.AST
	case "Sage":
		return actors.SGE
	case "Monk":
		return actors.MNK
	case "Dragoon":
		return actors.DRG
	case "Ninja":
		return actors.NIN
	case "Samurai":
		return actors.SAM
	case "Reaper":
		return actors.RPR
	case "Viper":
		return actors.VPR
	case "BlackMage":
		return actors.BLM
	case "Summoner":
		return actors.SMN
	case "RedMage":
		return actors.RDM
	case "Pictomancer":
		return actors.PCT
	case "Bard":
		return actors.BRD
	case "Machinist":
		return actors.MCH
	case "Dancer":
		return actors.DNC
	default:
		return "UNKNOWN"
	}
}

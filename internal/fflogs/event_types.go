package fflogs

type RawEventType string

const (

	// Casts
	BeginCast = RawEventType("begincast")
	Cast      = RawEventType("cast")

	// Damage
	CalculatedDamage = RawEventType("calculateddamage") //snapshot
	Damage           = RawEventType("damage")           //application

	// Heals
	CalculatedHeal = RawEventType("calculatedheal") //snapshot
	Heal           = RawEventType("heal")           //application
	Absorbed       = RawEventType("absorbed")       //shield absorption

	// Buffs
	ApplyBuff       = RawEventType("applybuff")
	RemoveBuff      = RawEventType("removebuff")
	RefreshBuff     = RawEventType("refreshbuff")
	ApplyBuffStack  = RawEventType("applybuffstack")
	RemoveBuffStack = RawEventType("removebuffstack")

	// Debuffs
	ApplyDebuff       = RawEventType("applydebuff")
	RemoveDebuff      = RawEventType("removedebuff")
	RefreshDebuff     = RawEventType("refreshdebuff")
	ApplyDebuffStack  = RawEventType("applydebuffstack")
	RemoveDebuffStack = RawEventType("removedebuffstack")
	HeadMarker        = RawEventType("headmarker")

	// Others
	TargetabilityUpdate = RawEventType("targetabilityupdate")
	Death               = RawEventType("death")
	Dispel              = RawEventType("dispel")
	EncounterEnd        = RawEventType("encounterend")
	CombatantInfo       = RawEventType("combatantinfo")
	LimitBreakUpdate    = RawEventType("limitbreakupdate")
	Tether              = RawEventType("tether")
	GaugeUpdate         = RawEventType("gaugeupdate")
	MapChange           = RawEventType("mapchange")
)

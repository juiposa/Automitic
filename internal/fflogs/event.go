package fflogs

type Event struct {
	//Meta
	PacketID  int   `json:"packetID,omitempty"`
	Fight     int   `json:"fight,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`

	//Actors
	SourceID       int `json:"sourceID,omitempty"`
	TargetID       int `json:"targetID,omitempty"`
	SourceInstance int `json:"sourceInstance,omitempty"`
	AbilityGameID  int `json:"abilityGameID,omitempty"`

	//Damage
	Type    string `json:"type,omitempty"`
	HitType int    `json:"hitType,omitempty"`

	UnmitigatedAmount int `json:"unmitigatedAmount,omitempty"`

	Mitigated  int     `json:"mitigated,omitempty"`
	Blocked    int     `json:"blocked,omitempty"`
	Absorbed   int     `json:"absorbed,omitempty"`
	Multiplier float64 `json:"multiplier,omitempty"`

	Amount int `json:"amount,omitempty"`

	//Death Events
	Overkill int `json:"overkill,omitempty"`

	//Stupid shit
	Gear                 []any   `json:"gear,omitempty"`
	Auras                []Auras `json:"auras,omitempty"`
	Level                any     `json:"level,omitempty"`
	SimulatedCrit        any     `json:"simulatedCrit,omitempty"`
	SimulatedDirectHit   any     `json:"simulatedDirectHit,omitempty"`
	Strength             any     `json:"strength,omitempty"`
	Dexterity            any     `json:"dexterity,omitempty"`
	Vitality             any     `json:"vitality,omitempty"`
	Intelligence         any     `json:"intelligence,omitempty"`
	Mind                 any     `json:"mind,omitempty"`
	Piety                any     `json:"piety,omitempty"`
	Attack               any     `json:"attack,omitempty"`
	DirectHit            any     `json:"directHit,omitempty"`
	CriticalHit          any     `json:"criticalHit,omitempty"`
	AttackMagicPotency   any     `json:"attackMagicPotency,omitempty"`
	HealMagicPotency     any     `json:"healMagicPotency,omitempty"`
	Determination        any     `json:"determination,omitempty"`
	SkillSpeed           any     `json:"skillSpeed,omitempty"`
	SpellSpeed           any     `json:"spellSpeed,omitempty"`
	Tenacity             any     `json:"tenacity,omitempty"`
	Bars                 any     `json:"bars,omitempty"`
	GuessAmount          any     `json:"guessAmount,omitempty"`
	ActorPotencyRatio    any     `json:"actorPotencyRatio,omitempty"`
	Tick                 any     `json:"tick,omitempty"`
	Simulated            any     `json:"simulated,omitempty"`
	ExpectedCritRate     any     `json:"expectedCritRate,omitempty"`
	TargetInstance       any     `json:"targetInstance,omitempty"`
	Unpaired             any     `json:"unpaired,omitempty"`
	ExtraAbilityGameID   any     `json:"extraAbilityGameID,omitempty"`
	Duration             any     `json:"duration,omitempty"`
	Overheal             any     `json:"overheal,omitempty"`
	Absorb               any     `json:"absorb,omitempty"`
	Value                any     `json:"value,omitempty"`
	Melee                any     `json:"melee,omitempty"`
	Stack                any     `json:"stack,omitempty"`
	GaugeID              any     `json:"gaugeID,omitempty"`
	Data1                any     `json:"data1,omitempty"`
	Data2                any     `json:"data2,omitempty"`
	Data3                any     `json:"data3,omitempty"`
	Data4                any     `json:"data4,omitempty"`
	FinalizedAmount      any     `json:"finalizedAmount,omitempty"`
	ExpectedAmount       any     `json:"expectedAmount,omitempty"`
	DirectHitPercentage  any     `json:"directHitPercentage,omitempty"`
	ExtraInfo            any     `json:"extraInfo,omitempty"`
	BonusPercent         any     `json:"bonusPercent,omitempty"`
	AttackerID           any     `json:"attackerID,omitempty"`
	AttackerInstance     any     `json:"attackerInstance,omitempty"`
	MarkerID             any     `json:"markerID,omitempty"`
	TetherID             any     `json:"tetherID,omitempty"`
	Fake                 any     `json:"fake,omitempty"`
	Targetable           any     `json:"targetable,omitempty"`
	Immune               any     `json:"immune,omitempty"`
	SourceMarker         any     `json:"sourceMarker,omitempty"`
	TargetMarker         any     `json:"targetMarker,omitempty"`
	KillerID             any     `json:"killerID,omitempty"`
	KillerInstance       any     `json:"killerInstance,omitempty"`
	KillingAbilityGameID any     `json:"killingAbilityGameID,omitempty"`
	IsBuff               any     `json:"isBuff,omitempty"`
	EncounterID          any     `json:"encounterID,omitempty"`
	Name                 any     `json:"name,omitempty"`
	Difficulty           any     `json:"difficulty,omitempty"`
	Size                 any     `json:"size,omitempty"`
	Kill                 any     `json:"kill,omitempty"`
}

type Auras struct {
	Source  any `json:"source,omitempty"`
	Ability any `json:"ability,omitempty"`
	Stacks  any `json:"stacks,omitempty"`
	Icon    any `json:"icon,omitempty"`
	Name    any `json:"name,omitempty"`
}

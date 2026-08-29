package actors

// Actor Types
const (
	ActorTypePlayer = "Player"
	ActorTypeNPC    = "NPC"
)

// Actor Subtypes
const (
	ActorSubtypeNPC        = "NPC"
	ActorSubTypeBoss       = "Boss"
	ActorSubTypeLimitBreak = "LimitBreak"
)

type FightActors struct {
	Players []PlayerCharacter `json:"players"`
	Bosses  []Enemy           `json:"bosses"`
	Adds    []Enemy           `json:"adds"`
}

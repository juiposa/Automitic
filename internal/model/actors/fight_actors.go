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
	Players map[int64]PlayerCharacter `json:"players"`
	Bosses  map[int64]Enemy           `json:"bosses"`
	Adds    map[int64]Enemy           `json:"adds"`
}

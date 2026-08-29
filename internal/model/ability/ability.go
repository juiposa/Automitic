package ability

type AbilitySource string

const (
	PLAYER AbilitySource = "PLAYER"
	ENEMY  AbilitySource = "ENEMY"
)

type AbilityType string

const (
	DAMAGE     AbilityType = "DAMAGE"
	MITIGATION AbilityType = "MITIGATION"
	SHIELD     AbilityType = "SHIELD"
	HEAL       AbilityType = "HEAL"
)

type Ability struct {
	Id            int `json:"id"`
	Name          string
	AbilitySource `json:"ability_source"`
	AbilityType   `json:"ability_type"`
}

func GetAbility(id int) Ability {
	return Ability{
		Id:            id,
		Name:          "Placeholder",
		AbilitySource: PLAYER,
		AbilityType:   MITIGATION,
	}
}

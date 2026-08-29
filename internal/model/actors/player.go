package actors

type PlayerCharacter struct {
	ID     int64 `json:"id"`
	GameID float64
	Name   string
	Job    Job
}

type PlayerState struct {
	MaxHealth     int
	CurrentHealth int
	CurrentMana   int
	Resources     map[string]int
}

type Pet struct {
}

package model

type PlayerCharacter struct {
	Name string
	Job  Job
}

type PlayerState struct {
	MaxHealth     int
	CurrentHealth int
	CurrentMana   int
	Resources     map[string]int
}

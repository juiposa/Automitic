package logmanager

import (
	"slices"

	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/model/actors"
	"github.com/juiposa/Automitic/internal/model/game"
)

func GetFightActors(actorsList []*fflogs.ReportActor, fight *fflogs.ReportFight) actors.FightActors {
	fightActors := actors.FightActors{
		Players: make(map[int64]actors.PlayerCharacter),
		Bosses:  make(map[int64]actors.Enemy),
		Adds:    make(map[int64]actors.Enemy),
	}
	actorIDs := GetActorIDs(fight)

	for _, actor := range actorsList {
		if !slices.Contains(actorIDs, *actor.ID) { //ignore all report actors not present in this specific pull
			continue
		}
		switch *actor.Type {
		case actors.ActorTypePlayer:
			fightActors.Players[*actor.ID] = MakePlayerActor(actor)
		case actors.ActorTypeNPC:
			if *actor.SubType == actors.ActorSubTypeBoss {
				fightActors.Bosses[*actor.ID] = MakeEnemyActor(actor)
			} else if *actor.SubType == actors.ActorSubtypeNPC { // Adds might not always be NPC/NPC, idk
				fightActors.Adds[*actor.ID] = MakeEnemyActor(actor)
			}
		}
	}
	return fightActors
}

func GetActorIDs(fight *fflogs.ReportFight) []int64 {
	var ids []int64
	for _, actor := range fight.EnemyNPCs {
		ids = append(ids, *actor.ID)
	}
	for _, actorID := range fight.FriendlyPlayers {
		ids = append(ids, *actorID)
	}
	for _, actor := range fight.FriendlyPets {
		ids = append(ids, *actor.GameID)
	}
	return ids
}

func MakePlayerActor(actor *fflogs.ReportActor) actors.PlayerCharacter {
	return actors.PlayerCharacter{
		ID:     *actor.ID,
		GameID: *actor.GameID,
		Name:   *actor.Name,
		Job:    game.GetJobType(*actor.SubType),
	}
}

func MakeEnemyActor(actor *fflogs.ReportActor) actors.Enemy {
	return actors.Enemy{
		ID:     *actor.ID,
		GameID: *actor.GameID,
		Name:   *actor.Name,
	}
}

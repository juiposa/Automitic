package data

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/juiposa/Automitic/internal/model/game"
)

const ActionsCSV = "Action.csv"
const ActionsDescriptionCSV = "ActionTransient.csv"

func (d *dataManager) GetActions() (map[int64]game.Action, error) {
	if d.actions != nil {
		return d.actions, nil
	}
	actions := make(map[int64]game.Action)
	actionfile, err := d.openFile(d.getDataFile(ActionsCSV))
	if err != nil {
		return nil, err
	}
	defer actionfile.Close()
	descriptionfile, err := d.openFile(d.getDataFile(ActionsDescriptionCSV))
	if err != nil {
		return nil, err
	}
	defer descriptionfile.Close()
	acCSV := csv.NewReader(actionfile)
	descCSV := csv.NewReader(descriptionfile)

	_, _ = acCSV.Read()
	_, _ = descCSV.Read()

	for {
		action, err := acCSV.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		id, _ := strconv.Atoi(action[0])
		description, _ := descCSV.Read()
		actions[int64(id)] = game.Action{
			Id:          int64(id),
			Name:        action[1],
			Description: description[1],
		}
	}
	d.actions = actions
	return actions, nil
}

func (d *dataManager) getDataFile(filename string) string {
	return fmt.Sprintf("%s/%s", d.dataLocation, filename)
}

package commands

import (
	"fmt"
	"parking_lot/internal/model"
	"parking_lot/internal/pkg/log"
	"strings"
)

type ParkCmd struct {
	rawArguments string

	registrationNum string
	color           string
}

func NewParkCmd(args string) Commands {
	return &ParkCmd{
		rawArguments: args,
	}

}

func (c *ParkCmd) Parse() error {
	lines := strings.Split(c.rawArguments, " ")
	if len(lines) != 3 {
		return InvalidArgs
	}

	c.registrationNum = lines[1]
	c.color = lines[2]
	return nil
}

func (c ParkCmd) Execute(parkingLot model.ParkingLot) (model.ParkingLot, error) {

	slot, err := parkingLot.FindAvailableSlot()
	if err != nil {
		return parkingLot, err
	}

	err = parkingLot.Park(slot.ID, model.NewCar(c.registrationNum, c.color))
	if err != nil {
		return parkingLot, err
	}

	log.Write(fmt.Sprintf("Allocated slot number: %d\n", slot.ID))
	return parkingLot, nil

}

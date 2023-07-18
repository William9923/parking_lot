package commands

import (
	"fmt"
	"parking_lot/internal/model"
	"parking_lot/internal/pkg/log"
	"strconv"
	"strings"
)

type UnparkCmd struct {
	rawArguments string

	slotID int
}

func NewUnparkCmd(args string) Commands {
	return &UnparkCmd{
		rawArguments: args,
	}

}

func (c *UnparkCmd) Parse() error {
	lines := strings.Split(c.rawArguments, " ")
	if len(lines) != 2 {
		return InvalidArgs
	}

	slotID, err := strconv.Atoi(lines[1])
	if err != nil {
		return err
	}

	c.slotID = slotID
	return nil
}

func (c UnparkCmd) Execute(parkingLot model.ParkingLot) (model.ParkingLot, error) {
	err := parkingLot.UnPark(c.slotID)
	log.Write(fmt.Sprintf("Slot number %d is free\n", c.slotID))
	return parkingLot, err
}

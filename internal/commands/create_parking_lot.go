package commands

import (
	"fmt"
	"parking_lot/internal/model"
	"parking_lot/internal/pkg/log"
	"strconv"
	"strings"
)

type CreateParkingLotCmd struct {
	rawArguments string

	size int
}

func NewCreateParkingLotCmd(args string) Commands {
	return &CreateParkingLotCmd{
		rawArguments: args,
	}

}

func (c *CreateParkingLotCmd) Parse() error {
	lines := strings.Split(c.rawArguments, " ")
	if len(lines) != 2 {
		return InvalidArgs
	}

	size, err := strconv.Atoi(lines[1])
	if err != nil {
		return err
	}

	c.size = size
	return nil
}

func (c CreateParkingLotCmd) Execute(parkingLot model.ParkingLot) (model.ParkingLot, error) {
	log.Write(fmt.Sprintf("Created a parking lot with %d slots\n", c.size))
	return model.NewEmptyParkingLot(c.size), nil
}

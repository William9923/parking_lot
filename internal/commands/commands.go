package commands

import (
	"fmt"
	"parking_lot/internal/model"
	"strings"
)

var InvalidArgs = fmt.Errorf("invalid args")
var UnsupportedCmd = fmt.Errorf("unsupported command")

type Commands interface {
	Parse() error
	Execute(parkingLot model.ParkingLot) (model.ParkingLot, error)
}

func NewCommandFactory(rawData string) (Commands, error) {
	if strings.HasPrefix(rawData, "create_parking_lot") {
		return NewCreateParkingLotCmd(rawData), nil
	}

	if strings.HasPrefix(rawData, "park") {
		return NewParkCmd(rawData), nil
	}

	if strings.HasPrefix(rawData, "leave") {
		return NewUnparkCmd(rawData), nil
	}

	if strings.HasPrefix(rawData, "status") {
		return NewStatusCmd(), nil
	}

	if strings.HasPrefix(rawData, "registration_numbers_for_cars_with_colour") {
		return NewQueryByColorCmd(rawData), nil
	}

	return nil, UnsupportedCmd

}

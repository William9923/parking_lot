package commands

import (
	"fmt"
	"parking_lot/internal/model"
	"parking_lot/internal/pkg/log"
	"strings"
)

type QueryByColorCmd struct {
	rawArguments string

	color string
}

func NewQueryByColorCmd(args string) Commands {
	return &QueryByColorCmd{
		rawArguments: args,
	}

}

func (c *QueryByColorCmd) Parse() error {
	lines := strings.Split(c.rawArguments, " ")
	if len(lines) != 2 {
		return InvalidArgs
	}

	c.color = lines[1]
	return nil
}

func (c QueryByColorCmd) Execute(parkingLot model.ParkingLot) (model.ParkingLot, error) {
	slots := parkingLot.GetSlots()

	var cars []model.Car
	// O(n)
	for _, slot := range slots {
		if slot.Car.Color == c.color {
			cars = append(cars, slot.Car)
		}
	}

	for i, car := range cars {
		log.Write(fmt.Sprintf("%s", car.RegistrationNumber))
		if i < len(cars)-1 {
			log.Write(fmt.Sprint(", "))
		}
	}
	log.Write("\n")

	return parkingLot, nil
}

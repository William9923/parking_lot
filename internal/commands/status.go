package commands

import (
	"fmt"
	"parking_lot/internal/model"
	"parking_lot/internal/pkg/log"
)

type StatusCmd struct {
}

func NewStatusCmd() Commands {
	return &StatusCmd{}

}

func (c *StatusCmd) Parse() error {
	return nil
}

func (c StatusCmd) Execute(parkingLot model.ParkingLot) (model.ParkingLot, error) {
	log.Write(fmt.Sprint("Slot No. Registration No Colour\n"))
	slots := parkingLot.GetSlots()

	for _, slot := range slots {
		if slot.IsAllocated {
			log.Write(fmt.Sprintf("%d %s %s\n", slot.ID, slot.Car.RegistrationNumber, slot.Car.Color))
		}
	}

	return parkingLot, nil
}

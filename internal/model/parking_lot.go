package model

import "fmt"

var FullParkingLotErr = fmt.Errorf("Sorry, parking lot is full")
var OutOfBoundSlot = fmt.Errorf("Sorry, no slot is found")
var SlotAlreadyAllocated = fmt.Errorf("Sorry, the slot is already allocated")

type ParkingLot struct {
	Size  int
	Slots []Slot
}

func NewEmptyParkingLot(size int) ParkingLot {

	var slots = make([]Slot, size)

	for i := 0; i < size; i++ {
		slots[i] = NewEmptySlot(i + 1)
	}

	return ParkingLot{
		Size:  size,
		Slots: slots,
	}
}

func (p *ParkingLot) GetSlots() []Slot {
	return p.Slots
}

func (p *ParkingLot) FindAvailableSlot() (Slot, error) {
	for i := 0; i < p.Size; i++ {
		if !p.Slots[i].IsAllocated {
			return p.Slots[i], nil
		}
	}

	return Slot{}, FullParkingLotErr
}

func (p *ParkingLot) GetSlotByID(slotID int) (Slot, error) {
	id := slotID - 1
	if (id) >= p.Size {
		return Slot{}, OutOfBoundSlot
	}

	return p.Slots[id], nil
}

func (p *ParkingLot) Park(slotID int, car Car) error {
	// NOTE: because we use 1-based index
	id := slotID - 1
	if (id) >= p.Size {
		return OutOfBoundSlot
	}

	if p.Slots[id].IsAllocated {
		return SlotAlreadyAllocated
	}

	p.Slots[id].IsAllocated = true
	p.Slots[id].Car = car
	return nil
}

func (p *ParkingLot) UnPark(slotID int) error {
	// NOTE: because we use 1-based index
	id := slotID - 1
	p.Slots[id].IsAllocated = false
	p.Slots[id].Car = NewEmptyCar()
	return nil
}

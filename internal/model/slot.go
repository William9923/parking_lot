package model

type Slot struct {
	ID          int
	IsAllocated bool
	Car         Car
}

func NewEmptySlot(id int) Slot {
	return Slot{
		ID:          id,
		IsAllocated: false,
		Car:         NewEmptyCar(),
	}
}

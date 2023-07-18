package model

type Car struct {
	RegistrationNumber string
	Color              string
}

func NewCar(regisNum string, color string) Car {
	return Car{
		RegistrationNumber: regisNum,
		Color:              color,
	}
}

func NewEmptyCar() Car {
	return Car{
		RegistrationNumber: "-",
		Color:              "-",
	}
}

func IsEmptyCar(car Car) bool {
	return car.Color == "-" && car.RegistrationNumber == "-"
}

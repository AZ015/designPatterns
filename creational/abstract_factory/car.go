package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	LuxuryCarType = iota + 1
	FamilyCarType
)

type Car interface {
	NumDoors() int
	NumWheels() int
	NumSeats() int
}

type CarFactory struct {
}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}

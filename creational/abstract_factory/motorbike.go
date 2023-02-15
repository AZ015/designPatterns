package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	SportMotorbikeType = iota + 1
	CruiseMotorbikeType
)

type Motorbike interface {
	NumWheels() int
	NumSeats() int
	GetMotorbikeType() int
}

type MotorbikeFactory struct {
}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized", v))
	}
}

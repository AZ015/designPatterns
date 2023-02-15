package abstract_factory

type SportMotorbike struct {
}

func (sb *SportMotorbike) NumWheels() int {
	return 2
}

func (sb *SportMotorbike) NumSeats() int {
	return 1
}

func (sb *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}

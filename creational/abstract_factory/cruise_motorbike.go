package abstract_factory

type CruiseMotorbike struct {
}

func (cb *CruiseMotorbike) NumWheels() int {
	return 2
}

func (cb *CruiseMotorbike) NumSeats() int {
	return 2
}

func (cb *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}

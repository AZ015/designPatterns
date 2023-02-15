package abstract_factory

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := BuildFactory(MotorbikeFactoryType)
	require.NoError(t, err)

	motorbikeVehicle, err := motorbikeF.NewVehicle(SportMotorbikeType)
	require.NoError(t, err)

	sportBike, ok := motorbikeVehicle.(Motorbike)
	require.True(t, ok)
	t.Logf("Sport bike has type %v\n", sportBike.GetMotorbikeType())
}

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	require.NoError(t, err)

	carVehicle, err := carF.NewVehicle(LuxuryCarType)
	require.NoError(t, err)

	luxuryCar, ok := carVehicle.(Car)
	require.True(t, ok)
	t.Logf("Luxury car has doors %v\n", luxuryCar.NumDoors())
}

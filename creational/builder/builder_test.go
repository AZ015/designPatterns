package builder

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var manufacturingComplex = ManufacturingDirector{}

func TestBuilderPattern_CarBuilder(t *testing.T) {
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	require.Equal(t, car.Wheels, 4)
	require.Equal(t, car.Structure, "Car")
	require.Equal(t, car.Seats, 5)
}

func TestBuilderPattern_BikeBuilder(t *testing.T) {
	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1

	require.Equal(t, motorbike.Wheels, 2)
	require.Equal(t, motorbike.Structure, "Motorbike")
}

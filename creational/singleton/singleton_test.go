package singleton

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	require.NotNil(t, counter1)

	currentCount := counter1.AddOne()
	require.Equal(t, currentCount, 1)

	counter2 := GetInstance()
	require.NotNil(t, counter1)

	currentCount = counter2.AddOne()
	require.Equal(t, currentCount, 2)
}

package prototype

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	require.NotNil(t, shirtCache)

	item1, err := shirtCache.GetClone(White)
	require.NoError(t, err)
	require.Equal(t, item1, whitePrototype)

	shirt1, ok := item1.(*Shirt)
	require.True(t, ok)
	shirt1.SKU = "abbcc"

	item2, err := shirtCache.GetClone(White)
	require.NoError(t, err)

	shirt2, ok := item2.(*Shirt)
	require.True(t, ok)

	require.NotEqual(t, shirt1.SKU, shirt2.SKU)
	require.NotEqual(t, shirt1, shirt2)
}

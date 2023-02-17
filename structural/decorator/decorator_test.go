package decorator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"
	require.Contains(t, pizzaResult, expectedText)
}

func TestOnion_AddIngredient(t *testing.T) {
	onion := &Onion{}
	onionResult, err := onion.AddIngredient()
	require.Error(t, err)

	onion = &Onion{&PizzaDecorator{}}
	onionResult, err = onion.AddIngredient()
	require.NoError(t, err)

	require.Contains(t, onionResult, "onion")
}

func TestMeat_AddIngredient(t *testing.T) {
	meat := &Meat{}
	meatResult, err := meat.AddIngredient()
	require.Error(t, err)

	meat = &Meat{&PizzaDecorator{}}
	meatResult, err = meat.AddIngredient()
	require.NoError(t, err)

	require.Contains(t, meatResult, "meat")
}

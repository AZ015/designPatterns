package factory

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	require.NoError(t, err)

	msg := payment.Pay(10.30)
	require.Contains(t, msg, "paid using cash")
}

func TestCreatePaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	require.NoError(t, err)

	msg := payment.Pay(10.30)
	require.Contains(t, msg, "paid using debit card")
}

func TestCreatePaymentMethodNotExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)
	require.Error(t, err, "A payment method with ID 20 must return an error")
}

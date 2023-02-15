package factory

import (
	"errors"
	"fmt"
)

const (
	Cash = iota + 1
	DebitCard
	CreditCard
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case CreditCard:
		return new(CreditCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized", m))
	}
}

type CashPM struct {
}

func (c CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%.2f paid using cash\n", amount)
}

type DebitCardPM struct {
}

func (dc DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%.2f paid using debit card\n", amount)
}

type CreditCardPM struct {
}

func (cc CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%.2f paid using credit card\n", amount)
}

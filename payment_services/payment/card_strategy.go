package payment

import (
	"errors"
)

type CardPayment struct {
}

func NewCardPayment() *CardPayment {
	return &CardPayment{}
}

func (c *CardPayment) Pay(amount float64) error {
	return errors.New("el pago con tarjeta no está implementado en esta versión de ejemplo")
}

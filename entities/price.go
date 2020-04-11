package entities

import "errors"

type Price struct {
	initialized bool
	currency    *string
	amount      *uint64
}

func (p Price) Initialized() bool {
	return p.initialized
}

func (p Price) Currency() string {
	return *p.currency
}

func (p Price) Amount() uint64 {
	return *p.amount
}

func NewPrice(currency string, amount uint64) (Price, error) {

	if currency != "USD" {
		return Price{}, errors.New("currency is not USD")
	}

	if amount == 0 {
		return Price{}, errors.New("price amount can not be 0")
	}

	return Price{
		currency:    &currency,
		amount:      &amount,
		initialized: true,
	}, nil

}

// for a API server, the idiomatic approach to design packages is to package by domain entities
package options

import (
	"strings"
	"time"

	appErrors "github.com/aries-financial-inc/options-service/errors"
)

// string comparision is inefficient. okay for the use case.
// TODO: use int type. with custom codec
type optionsType string

const (
	CALL optionsType = "call"
	PUT  optionsType = "put"
)

// values are case insensitive
func (o optionsType) IsValid() error {
	switch (optionsType)(strings.ToLower((string)(o))) {
	case CALL, PUT:
	default:
		return appErrors.ErrInvalidOptionsType
	}
	return nil
}

// TODO: not a readable name. being consistent with the contract.
// maybe name it position along with contract change
// do not introduce new terminology
type longShort string

const (
	LONG  longShort = "long"
	SHORT longShort = "short"
)

func (l longShort) IsValid() error {
	switch (longShort)(strings.ToLower((string)(l))) {
	case LONG, SHORT:
	default:
		return appErrors.ErrInvalidLongShort
	}
	return nil
}

type OptionsContract struct {
	// cannot name the variable "type"
	OptionsType    optionsType `json:"type"`
	StrikePrice    float64     `json:"strike_price"`
	Bid            float64     `json:"bid"`
	Ask            float64     `json:"ask"`
	ExpirationDate time.Time   `json:"expiration_date"`
	LongShort      longShort   `json:"long_short"`
}

func (o OptionsContract) IsValid() error {
	if err := o.OptionsType.IsValid(); err != nil {
		return err
	}

	if o.StrikePrice <= 0 {
		return appErrors.ErrInvalidStrikePrice
	}

	if o.Bid <= 0 {
		return appErrors.ErrInvalidBidPrice
	}

	if o.Ask <= 0 {
		return appErrors.ErrInvalidAskPrice
	}

	if o.Ask < o.Bid {
		return appErrors.ErrAskBidMismatch
	}

	if err := o.LongShort.IsValid(); err != nil {
		return err
	}

	if o.ExpirationDate.IsZero() || o.ExpirationDate.Before(time.Now()) {
		return appErrors.ErrInvalidExpirationDate
	}

	return nil
}

func (o OptionsContract) CalculateBreakEvenPoint() float64 {
	// long call
	if o.LongShort == LONG && o.OptionsType == CALL {
		return precisionTotwodecimalPlaces(o.StrikePrice + o.Ask)
	}
	// short call
	if o.LongShort == SHORT && o.OptionsType == CALL {
		return precisionTotwodecimalPlaces(o.StrikePrice + o.Bid)
	}
	// long put
	if o.LongShort == LONG && o.OptionsType == PUT {
		return precisionTotwodecimalPlaces(o.StrikePrice - o.Ask)
	}
	// short put
	if o.LongShort == SHORT && o.OptionsType == PUT {
		return precisionTotwodecimalPlaces(o.StrikePrice - o.Bid)
	}
	return 0.0
}

func (o OptionsContract) CalculateProfitOrLoss(price float64) float64 {
	// long call
	if o.LongShort == LONG && o.OptionsType == CALL {
		return precisionTotwodecimalPlaces(maxFloat64(0, price-o.StrikePrice) - o.Ask)
	}
	// short call
	if o.LongShort == SHORT && o.OptionsType == CALL {
		return precisionTotwodecimalPlaces(o.Bid - maxFloat64(0, price-o.StrikePrice))
	}
	// long put
	if o.LongShort == LONG && o.OptionsType == PUT {
		return precisionTotwodecimalPlaces(maxFloat64(o.StrikePrice-price, 0) - o.Ask)
	}
	// short put
	if o.LongShort == SHORT && o.OptionsType == PUT {
		return precisionTotwodecimalPlaces(o.Bid - max(0, o.StrikePrice-price))
	}
	return 0.0
}

// TODO: move to separate utils
func maxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}

// a simple and naive solution, assuming prices are not too large
func precisionTotwodecimalPlaces(f float64) float64 {
	return float64(int(f * 100)) / 100
}
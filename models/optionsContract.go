package models

import (
	"time"

	appErrors "github.com/aries-financial-inc/options-service/errors"
)

type optionsType int

const (
	CALL optionsType = iota + 1
	PUT
)

type position int

const (
	LONG position = iota + 1
	SHORT
)

type OptionsContract struct {
	OptionsType    optionsType
	StrikePrice    float64
	BidPrice       float64
	AskPrice       float64
	ExpirationDate time.Time
	Position       position
}

func (o OptionsContract) IsValid() error {
	if o.OptionsType == 0 {
		return appErrors.ErrInvalidOptionsType
	}

	if o.StrikePrice <= 0 {
		return appErrors.ErrInvalidStrikePrice
	}

	if o.BidPrice <= 0 {
		return appErrors.ErrInvalidBidPrice
	}

	if o.AskPrice <= 0 {
		return appErrors.ErrInvalidAskPrice
	}

	if o.AskPrice < o.BidPrice {
		return appErrors.ErrAskBidMismatch
	}

	if o.Position == 0 {
		return appErrors.ErrInvalidPosition
	}

	if o.ExpirationDate.IsZero() || o.ExpirationDate.Before(time.Now()) {
		return appErrors.ErrInvalidExpirationDate
	}

	return nil
}

package models_test

import (
	"testing"

	"time"

	"github.com/aries-financial-inc/options-service/errors"
	"github.com/aries-financial-inc/options-service/models"
	"github.com/stretchr/testify/assert"
)

func TestOptionsContractInValidOptionType(t *testing.T) {
	assert.ErrorIs(t, models.OptionsContract{
		OptionsType: 0,
	}.IsValid(), errors.ErrInvalidOptionsType)
}

func TestOptionsContractInValidStrikePrice(t *testing.T) {
	assert.ErrorIs(t, models.OptionsContract{
		OptionsType: models.CALL,
	}.IsValid(), errors.ErrInvalidStrikePrice)
}

func TestOptionsContractInValidBidPrice(t *testing.T) {
	assert.ErrorIs(t, models.OptionsContract{
		OptionsType: models.CALL,
		StrikePrice: 100.0,
	}.IsValid(), errors.ErrInvalidBidPrice)
}

func TestOptionsContractInValidAskPrice(t *testing.T) {
	assert.ErrorIs(t, models.OptionsContract{
		OptionsType: models.CALL,
		StrikePrice: 100.0,
		BidPrice: 10.05,
	}.IsValid(), errors.ErrInvalidAskPrice)

	assert.ErrorIs(t, models.OptionsContract{
		OptionsType: models.CALL,
		StrikePrice: 100.0,
		AskPrice: 10.0,
		BidPrice: 12.4,
	}.IsValid(), errors.ErrAskBidMismatch)
}


func TestOptionsContractInValidPosition(t *testing.T) {
	assert.ErrorIs(t, models.OptionsContract{
		OptionsType: models.CALL,
		StrikePrice: 100.0,
		BidPrice: 10.05,
		AskPrice: 12.04,
	}.IsValid(), errors.ErrInvalidPosition)
}

func TestOptionsContractInExpirationDate(t *testing.T) {
	assert.ErrorIs(t, models.OptionsContract{
		OptionsType: models.CALL,
		StrikePrice: 100.0,
		BidPrice: 10.05,
		AskPrice: 12.04,
		Position: models.LONG,
	}.IsValid(), errors.ErrInvalidExpirationDate)

	assert.ErrorIs(t, models.OptionsContract{
		OptionsType: models.CALL,
		StrikePrice: 100.0,
		BidPrice: 10.05,
		AskPrice: 12.04,
		Position: models.LONG,
		ExpirationDate: time.Now().Add(-24 * time.Hour),
	}.IsValid(), errors.ErrInvalidExpirationDate)
}

func TestOptionsContractIsValid(t *testing.T) {
	expirationDate, err := time.Parse(time.RFC3339, "2025-12-17T00:00:00Z")
	assert.NoError(t, err)
	assert.NoError(t, models.OptionsContract{
		OptionsType: models.CALL,
		StrikePrice: 100.0,
		BidPrice: 10.05,
		AskPrice: 12.04,
		Position: models.LONG,
		ExpirationDate: expirationDate,
	}.IsValid())
}
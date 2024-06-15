package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/aries-financial-inc/options-service/controllers"
	"github.com/aries-financial-inc/options-service/models"
	"github.com/stretchr/testify/assert"

	"github.com/aries-financial-inc/options-service/errors"
)

// TODO: idiomatic way of writing tests in golang is to keep tests and code together. fix the folder structure
func TestOptionsContractModelValidation(t *testing.T) {
	t.Run("invalid options type", func(t *testing.T) {
		assert.ErrorIs(t, models.OptionsContract{
			OptionsType: 0,
		}.IsValid(), errors.ErrInvalidOptionsType)
	})

	t.Run("invalid strike price", func(t *testing.T) {
		assert.ErrorIs(t, models.OptionsContract{
			OptionsType: models.CALL,
		}.IsValid(), errors.ErrInvalidStrikePrice)
	})

	t.Run("invalid ask price", func(t *testing.T) {
		assert.ErrorIs(t, models.OptionsContract{
			OptionsType: models.CALL,
			StrikePrice: 100.0,
			BidPrice:    10.05,
		}.IsValid(), errors.ErrInvalidAskPrice)

		// ask price is smaller than bid price
		assert.ErrorIs(t, models.OptionsContract{
			OptionsType: models.CALL,
			StrikePrice: 100.0,
			AskPrice:    10.0,
			BidPrice:    12.4,
		}.IsValid(), errors.ErrAskBidMismatch)
	})

	t.Run("invalid expiration date", func(t *testing.T) {
		assert.ErrorIs(t, models.OptionsContract{
			OptionsType: models.CALL,
			StrikePrice: 100.0,
			BidPrice:    10.05,
			AskPrice:    12.04,
			Position:    models.LONG,
		}.IsValid(), errors.ErrInvalidExpirationDate)

		// expiration date in in the past
		assert.ErrorIs(t, models.OptionsContract{
			OptionsType:    models.CALL,
			StrikePrice:    100.0,
			BidPrice:       10.05,
			AskPrice:       12.04,
			Position:       models.LONG,
			ExpirationDate: time.Now().Add(-24 * time.Hour),
		}.IsValid(), errors.ErrInvalidExpirationDate)
	})

	t.Run("invalid options type", func(t *testing.T) {
		expirationDate, err := time.Parse(time.RFC3339, "2025-12-17T00:00:00Z")
		assert.NoError(t, err)
		assert.NoError(t, models.OptionsContract{
			OptionsType:    models.CALL,
			StrikePrice:    100.0,
			BidPrice:       10.05,
			AskPrice:       12.04,
			Position:       models.LONG,
			ExpirationDate: expirationDate,
		}.IsValid())
	})

}

func TestAnalysisEndpoint(t *testing.T) {
	// Your code here
}

func TestIntegration(t *testing.T) {
	// Your code here
}

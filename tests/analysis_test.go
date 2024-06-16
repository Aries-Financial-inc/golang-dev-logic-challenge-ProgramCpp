package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/aries-financial-inc/options-service/controllers"
	"github.com/aries-financial-inc/options-service/options"
	"github.com/stretchr/testify/assert"

	"github.com/aries-financial-inc/options-service/errors"
)

// TODO: idiomatic way of writing tests in golang is to keep tests and code together. fix the folder structure
func TestOptionsContractModelValidation(t *testing.T) {
	t.Run("invalid options type", func(t *testing.T) {
		assert.ErrorIs(t, options.OptionsContract{}.IsValid(), errors.ErrInvalidOptionsType)

		assert.ErrorIs(t, options.OptionsContract{
			OptionsType: "xxx",
		}.IsValid(), errors.ErrInvalidOptionsType)
	})

	t.Run("invalid strike price", func(t *testing.T) {
		assert.ErrorIs(t, options.OptionsContract{
			OptionsType: options.CALL,
		}.IsValid(), errors.ErrInvalidStrikePrice)
	})

	t.Run("invalid ask price", func(t *testing.T) {
		assert.ErrorIs(t, options.OptionsContract{
			OptionsType: options.CALL,
			StrikePrice: 100.0,
			Bid:         10.05,
		}.IsValid(), errors.ErrInvalidAskPrice)

		// ask price is smaller than bid price
		assert.ErrorIs(t, options.OptionsContract{
			OptionsType: options.CALL,
			StrikePrice: 100.0,
			Ask:         10.0,
			Bid:         12.4,
		}.IsValid(), errors.ErrAskBidMismatch)
	})

	t.Run("invalid position", func(t *testing.T) {
		assert.ErrorIs(t, options.OptionsContract{
			OptionsType: options.CALL,
			StrikePrice: 100.0,
			Bid:         10.05,
			Ask:         12.04,
		}.IsValid(), errors.ErrInvalidLongShort)

		assert.ErrorIs(t, options.OptionsContract{
			OptionsType: options.CALL,
			StrikePrice: 100.0,
			Bid:         10.05,
			Ask:         12.04,
			LongShort:   "xxx",
		}.IsValid(), errors.ErrInvalidLongShort)
	})

	t.Run("invalid expiration date", func(t *testing.T) {
		assert.ErrorIs(t, options.OptionsContract{
			OptionsType: options.CALL,
			StrikePrice: 100.0,
			Bid:         10.05,
			Ask:         12.04,
			LongShort:   options.LONG,
		}.IsValid(), errors.ErrInvalidExpirationDate)

		// expiration date in in the past
		assert.ErrorIs(t, options.OptionsContract{
			OptionsType:    options.CALL,
			StrikePrice:    100.0,
			Bid:            10.05,
			Ask:            12.04,
			LongShort:      options.LONG,
			ExpirationDate: time.Now().Add(-24 * time.Hour),
		}.IsValid(), errors.ErrInvalidExpirationDate)
	})

	t.Run("valid options contract", func(t *testing.T) {
		expirationDate, err := time.Parse(time.RFC3339, "2025-12-17T00:00:00Z")
		assert.NoError(t, err)
		assert.NoError(t, options.OptionsContract{
			OptionsType:    options.CALL,
			StrikePrice:    100.0,
			Bid:            10.05,
			Ask:            12.04,
			LongShort:      options.LONG,
			ExpirationDate: expirationDate,
		}.IsValid())
	})

}

// TODO: idiomatic way of writing tests in golang is to keep tests and code together. fix the folder structure.
// move this to controllers package
func TestAnalysisEndpoint(t *testing.T) {
	t.Run("error on no options", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/analyze", strings.NewReader(`{}`))
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		controllers.AnalysisHandler(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("error on less than 4 options", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/analyze", strings.NewReader(`[
  		{
   			"strike_price": 100, 
    		"type": "Call", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		}]`))
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		controllers.AnalysisHandler(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("error on more than 4 options", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/analyze", strings.NewReader(`[
  		{
   			"strike_price": 100, 
    		"type": "Call", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		},
		{
   			"strike_price": 100, 
    		"type": "Call", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		},
		{
   			"strike_price": 100, 
    		"type": "Call", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		},
		{
   			"strike_price": 100, 
    		"type": "Call", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		},
		{
   			"strike_price": 100, 
    		"type": "Call", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		}]`))
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		controllers.AnalysisHandler(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)

	})

	t.Run("error on invalid options", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/analyze", strings.NewReader(`[
  		{
   			"strike_price": 100, 
    		"type": "xxx", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		},
		{
   			"strike_price": 100, 
    		"type": "xxx", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		},
		{
   			"strike_price": 100, 
    		"type": "Call", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		},
		{
   			"strike_price": 100, 
    		"type": "Call", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		}]`))
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		controllers.AnalysisHandler(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)

	})

	t.Run("successful analysis", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/analyze", strings.NewReader(`[
		{
   			"strike_price": 100, 
    		"type": "Call", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		},
		{
   			"strike_price": 100, 
    		"type": "Call", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		},
		{
   			"strike_price": 100, 
    		"type": "Call", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		},
		{
   			"strike_price": 100, 
    		"type": "Call", 
    		"bid": 10.05, 
    		"ask": 12.04, 
    		"long_short": "long", 
    		"expiration_date": "2025-12-17T00:00:00Z"
  		}]`))
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		controllers.AnalysisHandler(res, req)
		assert.Equal(t, http.StatusOK, res.Code)

		// TODO: validate response body
	})
}

func TestIntegration(t *testing.T) {
	// Your code here
}

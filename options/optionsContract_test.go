package options_test

import (
	"testing"

	"github.com/aries-financial-inc/options-service/options"
	"github.com/stretchr/testify/assert"
)

func TestCalculateBreakEvenPoint(t *testing.T) {
	// strike + ask
	assert.Equal(t, 116.5, options.OptionsContract{
		LongShort:   options.LONG,
		OptionsType: options.CALL,
		StrikePrice: 102.5,
		Ask:         14.00,
		Bid:         12.10,
	}.CalculateBreakEvenPoint())

	// strike + bid
	assert.Equal(t, 114.6, options.OptionsContract{
		LongShort:   options.SHORT,
		OptionsType: options.CALL,
		StrikePrice: 102.5,
		Ask:         14.00,
		Bid:         12.10,
	}.CalculateBreakEvenPoint())

	// strike - ask
	assert.Equal(t, 88.5, options.OptionsContract{
		LongShort:   options.LONG,
		OptionsType: options.PUT,
		StrikePrice: 102.5,
		Ask:         14.00,
		Bid:         12.10,
	}.CalculateBreakEvenPoint())

	// strike - bid
	assert.Equal(t, 90.4, options.OptionsContract{
		LongShort:   options.SHORT,
		OptionsType: options.PUT,
		StrikePrice: 102.5,
		Ask:         14.00,
		Bid:         12.10,
	}.CalculateBreakEvenPoint())
}

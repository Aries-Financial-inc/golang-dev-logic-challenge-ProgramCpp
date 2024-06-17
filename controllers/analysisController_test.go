package controllers_test

import (
	"testing"

	"github.com/aries-financial-inc/options-service/controllers"
	"github.com/aries-financial-inc/options-service/options"
	"github.com/stretchr/testify/assert"
)

// TODO: mock options contract to fine tune unit tests
func TestCalculateXY(t *testing.T) {
	assert.ElementsMatch(t, []controllers.XYValue{
		{
			X: 100,
			Y: -12.04,
		},
		{
			X: 112.04,
			Y: 0,
		},
		{
			X: 102.50,
			Y: -14,
		},
		{
			X: 116.5,
			Y: 0,
		},
		{
			X: 103,
			Y: 14,
		},
		{
			X: 89,
			Y: 0,
		},
		{
			X: 105,
			Y: -18,
		},
		{
			X: 87,
			Y: 0,
		},
		{
			X: 0,
			Y: -12.04,
		},
		{
			X: 210,
			Y: 97.96,
		},
		{
			X: 0,
			Y: -14,
		},
		{
			X: 210,
			Y: 93.50,
		},
		{
			X: 0,
			Y: -89,
		},
		{
			X: 210,
			Y: 14,
		},
		{
			X: 0,
			Y: 87,
		},
		{
			X: 210,
			Y: -18,
		},
	}, controllers.CalculateXYValues([]options.OptionsContract{
		{
			StrikePrice: 100,
			OptionsType: "Call",
			Bid:         10.05,
			Ask:         12.04,
			LongShort:   "long",
		},
		{
			StrikePrice: 102.50,
			OptionsType: "Call",
			Bid:         12.10,
			Ask:         14,
			LongShort:   "long",
		},
		{
			StrikePrice: 103,
			OptionsType: "Put",
			Bid:         14,
			Ask:         15.50,
			LongShort:   "short",
		},
		{
			StrikePrice: 105,
			OptionsType: "Put",
			Bid:         16,
			Ask:         18,
			LongShort:   "long",
		},
	},
	))
}


func TestCalculateMaxProfit(t *testing.T) {
	assert.Equal(t, 97.96, controllers.CalculateMaxProfit([]options.OptionsContract{
		{
			StrikePrice: 100,
			OptionsType: "Call",
			Bid:         10.05,
			Ask:         12.04,
			LongShort:   "long",
		},
		{
			StrikePrice: 102.50,
			OptionsType: "Call",
			Bid:         12.10,
			Ask:         14,
			LongShort:   "long",
		},
		{
			StrikePrice: 103,
			OptionsType: "Put",
			Bid:         14,
			Ask:         15.50,
			LongShort:   "short",
		},
		{
			StrikePrice: 105,
			OptionsType: "Put",
			Bid:         16,
			Ask:         18,
			LongShort:   "long",
		},
	},
	))
}

func TestCalculateMaxLoss(t *testing.T) {
	assert.Equal(t, -89.0, controllers.CalculateMaxLoss([]options.OptionsContract{
		{
			StrikePrice: 100,
			OptionsType: "Call",
			Bid:         10.05,
			Ask:         12.04,
			LongShort:   "long",
		},
		{
			StrikePrice: 102.50,
			OptionsType: "Call",
			Bid:         12.10,
			Ask:         14,
			LongShort:   "long",
		},
		{
			StrikePrice: 103,
			OptionsType: "Put",
			Bid:         14,
			Ask:         15.50,
			LongShort:   "short",
		},
		{
			StrikePrice: 105,
			OptionsType: "Put",
			Bid:         16,
			Ask:         18,
			LongShort:   "long",
		},
	},
	))
}

func TestCalculateBreakEvenPoints(t * testing.T){
	assert.ElementsMatch(t, []float64{112.04, 116.5, 89.0, 87.0}, controllers.CalculateBreakEvenPoints([]options.OptionsContract{
		{
			StrikePrice: 100,
			OptionsType: "Call",
			Bid:         10.05,
			Ask:         12.04,
			LongShort:   "long",
		},
		{
			StrikePrice: 102.50,
			OptionsType: "Call",
			Bid:         12.10,
			Ask:         14,
			LongShort:   "long",
		},
		{
			StrikePrice: 103,
			OptionsType: "Put",
			Bid:         14,
			Ask:         15.50,
			LongShort:   "short",
		},
		{
			StrikePrice: 105,
			OptionsType: "Put",
			Bid:         16,
			Ask:         18,
			LongShort:   "long",
		},
	},
	))
}

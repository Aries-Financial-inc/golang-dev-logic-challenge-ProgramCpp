package controllers

import (
	"encoding/json"
	"io"
	"math"
	"net/http"

	"github.com/aries-financial-inc/options-service/options"
)

// AnalysisResponse represents the data structure of the analysis result
// TODO: group XY values for each option
type AnalysisResponse struct {
	XYValues        []XYValue `json:"xy_values"`
	MaxProfit       float64   `json:"max_profit"`
	MaxLoss         float64   `json:"max_loss"`
	BreakEvenPoints []float64 `json:"break_even_points"`
}

// XYValue represents a pair of X and Y values
type XYValue struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// TODO: add logging for all failures
func AnalysisHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	options := []options.OptionsContract{}
	err = json.Unmarshal(body, &options)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(options) != 4 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, o := range options {
		if err := o.IsValid(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// TODO: fix repeated computations of X and Y values
	resp := AnalysisResponse{
		XYValues:        calculateXYValues(options),
		MaxProfit:       calculateMaxProfit(options),
		MaxLoss:         calculateMaxLoss(options),
		BreakEvenPoints: calculateBreakEvenPoints(options),
	}

	res, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

// for a option, the range of X is (0, 2 * strike price)
// the range of X for the graph is the (minimum X, maximum X), for all options
// for a given X, calculate profits or losses for all options.  this enables comparision of options' profits and losses
// the values of X are min X, max X, all strike prices and all break even points
func calculateXYValues(contracts []options.OptionsContract) []XYValue {
	xValues := []float64{0}
	xMax := 0.0
	// calculate range of X for each option
	for _, c := range contracts {
		xValues = append(xValues, c.StrikePrice)
		xValues = append(xValues, c.CalculateBreakEvenPoint())
		if xMax < 2 * c.StrikePrice {
			xMax = 2 * c.StrikePrice
		}
	}
	xValues = append(xValues, xMax)

	xyValues := []XYValue{}
	// calculate Y values for all X values
	for _, x := range xValues {
		for _, c := range contracts {
			xyValues = append(xyValues, XYValue{x, c.CalculateProfitOrLoss(x)})
		}
	}

	return xyValues
}

func calculateMaxProfit(contracts []options.OptionsContract) float64 {
	maxProfit := 0.0
	profitLosses := calculateXYValues(contracts)

	for _, pl := range profitLosses {
		if pl.Y > maxProfit {
			maxProfit = pl.Y
		}
	}
	return maxProfit
}

func calculateMaxLoss(contracts []options.OptionsContract) float64 {
	maxLoss := math.MaxFloat64
	profitLosses := calculateXYValues(contracts)

	for _, pl := range profitLosses {
		if pl.Y < maxLoss {
			maxLoss = pl.Y
		}
	}
	return maxLoss
}

func calculateBreakEvenPoints(contracts []options.OptionsContract) []float64 {
	breakEvens := []float64{}
	for _, c := range contracts {
		breakEvens = append(breakEvens, c.CalculateBreakEvenPoint())
	}
	return breakEvens
}

package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/aries-financial-inc/options-service/options"
)

// AnalysisResponse represents the data structure of the analysis result
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

// for visialization of profit and loss, the range of X, is +/- 25% of break-even price
func calculateXYValues(contracts []options.OptionsContract) []XYValue {
	xyValues := []XYValue{}
	return xyValues
}

func calculateMaxProfit(contracts []options.OptionsContract) float64 {
	// Your code here
	return 0
}

func calculateMaxLoss(contracts []options.OptionsContract) float64 {
	// Your code here
	return 0
}

func calculateBreakEvenPoints(contracts []options.OptionsContract) []float64 {
	// Your code here
	return nil
}

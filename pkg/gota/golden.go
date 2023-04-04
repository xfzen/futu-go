package gota

import (
	"fmt"
	"strconv"
)

func CalcGoldenRatio(price float64, count int) ([]float64, []float64) {
	ratios := make([]float64, count)
	ratio := price

	for i := 0; i < count; i++ {
		ratios[i] = ratio
		ratio = ratio * 0.618
	}

	for i := 1; i < count; i++ {
		ratios[i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", ratios[i]), 64)
	}

	averages := make([]float64, count-1)
	for i := 0; i < count-1; i++ {
		averages[i] = (ratios[i] + ratios[i+1]) / 2
		averages[i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", averages[i]), 64)
	}

	return ratios, averages
}

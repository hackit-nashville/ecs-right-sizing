package lib

import (
	"math"
)

func calculateReservation(results []*float64) int64 {
	var hours [24]float64
	var hour = 0

	resultsLength := len(results)
	if resultsLength < 24 {
		resultsLength = 24
	}
	if resultsLength%24 != 0 {
		panic("Invalid number of data points. Must be intervals of 24")
	}

	for _, result := range results {
		hours[hour] = hours[hour] + *result
		hour = hour + 1
		if hour == 24 {
			hour = 0
		}
	}

	for hour, value := range hours {
		hours[hour] = value / (float64(resultsLength) / 24)
	}

	max := Max(hours)

	return int64(math.Round(max / .8))
}

// Max returns themaximum value of a slice with 24 datapoints
func Max(array [24]float64) float64 {
	var max float64
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

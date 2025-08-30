package neighbours

import (
	"fmt"
	"math"
)

var (
	sliceLengthMismatch = fmt.Errorf("Slices need to be of equal length.")
)

type DistanceFunc interface {
	Distance(x, y []float64) (float64, error)
}

// Can only move in straight lines from point x to point y
//
//	------y
//	|
//	|
//	x
type Manhattan struct{}

func (m *Manhattan) Distance(x, y []float64) (float64, error) {

	if len(x) != len(y) {
		return 0, sliceLengthMismatch
	}

	var result float64
	for i, val := range x {
		result += math.Abs(val - y[i])
	}
	return result, nil
}

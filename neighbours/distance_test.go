package neighbours

import (
	"errors"
	"math"
	"testing"
)

func TestDistanceFunctions(t *testing.T) {
	euclidean := &Euclidean{}
	tests := []struct {
		distanceFunc     DistanceFunc
		name             string
		point1           []float64
		point2           []float64
		expectedDistance float64
		wantError        error
	}{
		{
			euclidean, "Euclidean: x, y plane distance calculation", []float64{1, 2}, []float64{2, 3}, math.Sqrt2, nil,
		},
		{
			euclidean, "Euclidean: x, y, z plane distance calculation", []float64{1, 2, 3}, []float64{2, 3, 4}, math.Sqrt(3), nil,
		},
		{
			euclidean, "Euclidean: non-equal feature counts should error", []float64{1, 2}, []float64{2, 3, 4}, 1.73205, SliceLengthMismatch,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			res, err := test.distanceFunc.distance(test.point1, test.point2)

			if test.wantError != nil {
				if !errors.Is(test.wantError, err) {
					t.Errorf("Expected error: %v, Actual: %v", test.wantError, err)
				}
				return
			}

			if test.expectedDistance != res {
				t.Errorf("Expected: %f, Actual: %f", test.expectedDistance, res)
			}
		})
	}
}

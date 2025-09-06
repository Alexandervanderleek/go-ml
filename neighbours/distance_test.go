package neighbours

import (
	"math"
	"testing"
)

func TestEuclideanDistance(t *testing.T) {
	tests := []struct {
		name             string
		point1           []float64
		point2           []float64
		expectedDistance float64
		wantError        error
	}{
		{
			"x, y plane distance calculation", []float64{1, 2}, []float64{2, 3}, 1.414213, nil,
		},
		{
			"x, y, z plane distance calculation", []float64{1, 2, 3}, []float64{2, 3, 4}, 1.414213, nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			euclidean := &Euclidean{}
			res, _ := euclidean.Distance(test.point1, test.point2)

			const epsilon = 1e-6
			if diff := math.Abs(res - test.expectedDistance); diff > epsilon {
				t.Errorf("Expected: %f, Actual: %f", test.expectedDistance, res)
			}

		})
	}

}

package neighbours

import (
	"testing"

	goml "github.com/Alexandervanderleek/go-ml"
)

func TestKNNMahattanDistance(t *testing.T) {

	dataSet := &goml.DataSet{
		Samples: []goml.Sample{
			{Features: []float64{1, 2}, Label: "A"},
			{Features: []float64{2, 3}, Label: "A"},
			{Features: []float64{3, 4}, Label: "A"},
			{Features: []float64{6, 7}, Label: "B"},
			{Features: []float64{7, 8}, Label: "B"},
		},
		FeatureNames: []string{"Age", "Height"},
	}

	knn := NewKnn(WithDistance(&Manhattan{}))

	closest, err := knn.CalculateNeighbours(1, dataSet, []float64{4, 5})

	if err != nil {
		t.Fatal(err)
	}

	if closest[0].point.Label != "A" {
		t.Errorf("Expected label A, got %s", closest[0].point.Label)
	}

}

package neighbours

import (
	"fmt"
	"sort"

	goml "github.com/Alexandervanderleek/go-ml"
)

type KNN struct {
	distanceFunc DistanceFunc
}

type PointWithDistance struct {
	point    goml.Sample
	distance float64
}

type Option func(*KNN)

func NewKnn(options ...Option) {

	knn := &KNN{}
	for _, option := range options {
		option(knn)
	}

}

func (knn *KNN) CalculateNeighbours(k uint, data goml.DataSet, point goml.Sample) ([]PointWithDistance, error) {
	if k == 1 {
		return nil, fmt.Errorf("Require neighbours of at least 1.")
	}

	resultDistances := make([]PointWithDistance, len(data.Samples))

	for i, val := range data.Samples {
		distance, err := knn.distanceFunc.Distance(val.Features, point.Features)

		resultDistances = append(resultDistances, PointWithDistance{
			val,
			distance,
		})
	}

	sort.Slice(resultDistances, func(i, j int) bool {
		return resultDistances[i].distance < resultDistances[j].distance
	})

	return resultDistances, nil
}

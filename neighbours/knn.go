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

func WithDistance(df DistanceFunc) Option {
	return func(knn *KNN) {
		knn.distanceFunc = df
	}
}

func NewKnn(options ...Option) *KNN {

	knn := &KNN{}
	for _, option := range options {
		option(knn)
	}

	return knn

}

func (knn *KNN) CalculateNeighbours(k uint, data *goml.DataSet, point []float64) ([]PointWithDistance, error) {
	if k < 1 {
		return nil, fmt.Errorf("Require neighbours of at least 1.")
	}

	resultDistances := make([]PointWithDistance, 0, len(data.Samples))

	for _, val := range data.Samples {
		distance, err := knn.distanceFunc.Distance(val.Features, point)

		if err != nil {
			return nil, err
		}

		resultDistances = append(resultDistances, PointWithDistance{
			val,
			distance,
		})
	}

	sort.Slice(resultDistances, func(i, j int) bool {
		return resultDistances[i].distance < resultDistances[j].distance
	})

	return resultDistances[:k], nil
}

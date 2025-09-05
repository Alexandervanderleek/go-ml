package neighbours

import (
	"container/heap"
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

func (knn *KNN) FindNeighbors(k uint, data *goml.DataSet, point []float64) ([]PointWithDistance, error) {

	if k < 1 {
		return nil, fmt.Errorf("Require K value of at least 1")
	}

	maxHeap := &neighboursHeap{}
	heap.Init(maxHeap)

	for _, val := range data.Samples {
		distance, err := knn.distanceFunc.Distance(val.Features, point)

		if err != nil {
			return nil, err
		}

		if maxHeap.Len() < int(k) {
			heap.Push(maxHeap, PointWithDistance{
				val,
				distance,
			})
		} else {
			if distance < (*maxHeap)[0].distance {
				heap.Pop(maxHeap)
				heap.Push(maxHeap, PointWithDistance{
					val,
					distance,
				})
			}
		}
	}

	sort.Slice(*maxHeap, func(i, j int) bool {
		return (*maxHeap)[i].distance < (*maxHeap)[j].distance
	})

	return *maxHeap, nil
}

package goml

type Sample struct {
	Features []float64 // {100, 20, 10}
	Label    any       // e.g Japan
}

type DataSet struct {
	Samples      []Sample
	FeatureNames []string // height, age, ect
}

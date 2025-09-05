package neighbours

type neighboursHeap []PointWithDistance

func (n neighboursHeap) Len() int           { return len(n) }
func (n neighboursHeap) Less(i, j int) bool { return n[i].distance > n[j].distance }
func (n neighboursHeap) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n *neighboursHeap) Pop() any {
	oldValue := (*n)[len(*n)-1]
	*n = (*n)[0 : len(*n)-1]
	return oldValue
}
func (n *neighboursHeap) Push(in any) { *n = append(*n, in.(PointWithDistance)) }

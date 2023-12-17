package main

type Point struct {
	r        int
	c        int
	heatloss int
	steps    int
	dir      string // 'N', 'W', 'S', 'E'
}

type PointHeap []Point

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i].heatloss < h[j].heatloss }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() any {
	// old := *pq
	// n := len(old)
	// item := old[n-1]
	// old[n-1] = nil  // avoid memory leak
	// item.index = -1 // for safety
	// *pq = old[0 : n-1]
	// return item

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

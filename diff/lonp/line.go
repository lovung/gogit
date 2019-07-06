package lonp

import (
	"sync"
)

// LineDiff struct
type LineDiff struct {
	A, B string
	m, n int
}

// NewDiff create the new Diff
func NewDiff(str1, str2 string) *LineDiff {
	if len(str1) > len(str2) {
		return &LineDiff{str2, str1, len(str2), len(str1)}
	}
	return &LineDiff{str1, str2, len(str1), len(str2)}
}

func (d *LineDiff) snake(k, p, pp int) int {
	var y int
	if p > pp {
		y = p
	} else {
		y = pp
	}
	x := y - k
	for x < d.m && y < d.n && d.A[x] == d.B[y] {
		x++
		y++
	}
	return y
}

// EditDistance return the edit distance
func (d *LineDiff) EditDistance() int {
	offset := d.m + 1
	delta := d.n - d.m
	size := d.m + d.n + 3
	fp := make([]int, size)

	for i := range fp {
		fp[i] = -1
	}

	var w sync.WaitGroup
	for p := 0; ; p++ {
		w.Add(2)
		go func() {
			for k := -p; k <= delta-1; k++ {
				fp[k+offset] = d.snake(k, fp[k-1+offset]+1, fp[k+1+offset])
			}
			w.Done()
		}()

		go func() {
			for k := delta + p; k >= delta+1; k-- {
				fp[k+offset] = d.snake(k, fp[k-1+offset]+1, fp[k+1+offset])
			}
			w.Done()
		}()
		w.Wait()
		fp[delta+offset] = d.snake(delta, fp[delta-1+offset]+1, fp[delta+1+offset])
		if fp[delta+offset] >= d.n {
			return delta + 2*p
		}
	}
}

/**
 * The algorithm implemented here is based on "An O(NP) Sequence Comparison Algorithm"
 * by described by Sun Wu, Udi Manber and Gene Myers
 */

package onp

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Diff struct
type Diff struct {
	str1, str2 string
	m, n       int
}

// NewDiff create the new diff
func NewDiff(str1, str2 string) *Diff {
	if len(str1) > len(str2) {
		return &Diff{str2, str1, len(str2), len(str1)}
	}
	return &Diff{str1, str2, len(str1), len(str2)}
}

func (diff *Diff) snake(k, p, pp int) int {

	y := max(p, pp)
	x := y - k

	for x < diff.m && y < diff.n && diff.str1[x] == diff.str2[y] {
		x++
		y++
	}

	return y
}

const notSolved = -1

// EditDistance show the edit distance
func (diff *Diff) EditDistance() int {

	offset := diff.m + 1
	delta := diff.n - diff.m
	size := diff.m + diff.n + 3
	fp := make([]int, size)

	for i := range fp {
		fp[i] = -1
	}

	for p := 0; ; p++ {

		done := make(chan bool)
		go func() {
			for k := -p; k <= delta-1; k++ {
				fp[k+offset] = diff.snake(k, fp[k-1+offset]+1, fp[k+1+offset])
			}

			done <- true
		}()

		go func() {
			for k := delta + p; k >= delta+1; k-- {
				fp[k+offset] = diff.snake(k, fp[k-1+offset]+1, fp[k+1+offset])
			}
			done <- true
		}()

		for i := 0; i < 2; i++ {
			<-done
		}
		fp[delta+offset] = diff.snake(delta, fp[delta-1+offset]+1, fp[delta+1+offset])

		if fp[delta+offset] >= diff.n {
			return delta + 2*p
		}
	}

	return notSolved
}

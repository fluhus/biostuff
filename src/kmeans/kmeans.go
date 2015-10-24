// Performs k-means analysis on data.
package kmeans

import (
	"fmt"
	"math/rand"
)

// Performs k-means clustering on the given data. Each vector is an element in
// the clustering. Returns the tags each element was given, and the average
// distance of elements from their assigned means.
func Kmeans(data [][]float64, k int) (tags []int, means [][]float64,
		distortion float64) {
	return (&matrix{data}).kmeans(k)
}

// A matrix type for convenience functions.
type matrix struct {
	data [][]float64
}

// Returns the number of rows in the matrix.
func (m *matrix) nrows() int {
	return len(m.data)
}

// Returns the number of columns in the matrix.
func (m *matrix) ncols() int {
	if len(m.data) == 0 {
		return 0
	} else {
		return len(m.data[0])
	}
}

// Performs k-means clustering on the rows of the matrix.
func (m *matrix) kmeans(k int) (tags []int, cents [][]float64, dist float64) {
	// Must be at least 1.
	if k < 1 {
		panic(fmt.Sprint("Bad k:", k))
	}

	// If k is too large - that's ok just reduce to avoid out-of-range.
	if k > m.nrows() {
		k = m.nrows()
	}

	// Create initial centroids.
	initCents := rand.Perm(m.nrows())[:k]
	cents = make([][]float64, k)
	for i := range cents {
		cents[i] = make([]float64, m.ncols())
		copy(cents[i], m.data[initCents[i]])
	}

	// First tagging.
	tags = m.tag(cents)
	dist = m.distortion(cents, tags)
	distOld := 2 * dist

	// Iterate until converged.
	for dist > distOld || dist / distOld < 0.999 {
		distOld = dist
		cents = m.cent(tags, k)
		tags = m.tag(cents)
		dist = m.distortion(cents, tags)
	}

	return
}

// Tags each row with the index of its nearest centroid.
func (m *matrix) tag(cents [][]float64) []int {
	if len(cents) == 0 {
		panic("Cannot tag on 0 centroids.")
	}
	
	tags := make([]int, m.nrows())

	for i := range tags {
		// Find nearest centroid.
		tags[i] = 0
		d := distance(cents[0], m.data[i])
		for j := 1; j < len(cents); j++ {
			dj := distance(cents[j], m.data[i])
			if dj < d {
				d = dj
				tags[i] = j
			}
		}
	}

	return tags
}

// Calculates the distortion of the given tagging and centroids.
func (m *matrix) distortion(cents [][]float64, tags []int) float64 {
	if len(tags) != m.nrows() {
		panic(fmt.Sprintf("Non-matching lengths of matrix and tags: %d, %d",
				m.nrows(), len(tags)))
	}
	if m.nrows() == 0 {
		return 0
	}

	d := 0.0
	for i := range tags {
		d += distance(cents[tags[i]], m.data[i])
	}

	return d / float64(m.nrows())
}

// Calculates the new centroids, according to average of tagged rows in each
// group.
func (m *matrix) cent(tags []int, k int) [][]float64 {
	cents := make([][]float64, k)
	for i := range cents {
		cents[i] = make([]float64, m.ncols())
	}
	counts := make([]int, k)

	for i := range m.data {
		counts[tags[i]]++
		for j := range m.data[i] {
			cents[tags[i]][j] += m.data[i][j]
		}
	}

	for i := range cents {
		for j := range cents[i] {
			if counts[i] != 0 {
				cents[i][j] /= float64(counts[i])
			} else {
				cents[i][j] = 0
			}
		}
	}

	return cents
}

// Returns the L1 (Manhattan) distance between 2 vectors.
func distance(a, b []float64) float64 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Non-matching lengths: %d, %d",
				len(a), len(b)))
	}
	
	d := 0.0
	for i := range a {
		d += abs(a[i] - b[i])
	}

	return d
}

// Returns the absolute value of a number.
func abs(a float64) float64 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}


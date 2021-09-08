// Package regions provides an index for searching on intervals.
package regions

import (
	"fmt"
	"sort"
)

// Index is a searchable collection of intervals.
type Index struct {
	idx []interval
}

// NewIndex returns an index on the given interval starts and ends. Starts and ends
// should be of the same length. End positions are exclusive, meaning that an end
// value of n implies that the interval's last position is n-1.
func NewIndex(starts, ends []int) *Index {
	if len(starts) != len(ends) {
		panic(fmt.Sprintf("lengths of starts and ends don't match: %v!=%v",
			len(starts), len(ends)))
	}
	events := make([]event, 0, len(starts)+len(ends))
	for i := range starts {
		// TODO(amit): Check that start<end
		events = append(events, event{i, starts[i], true})
		events = append(events, event{i, ends[i], false})
	}
	sort.Slice(events, func(i, j int) bool {
		return eventLess(events[i], events[j])
	})
	var intervals []interval
	idxs := map[int]struct{}{}
	var pos int
	for i, e := range events {
		if i == 0 {
			pos = e.pos
		}
		if e.pos != pos {
			intervals = append(intervals, interval{pos, keys(idxs)})
			pos = e.pos
		}
		if e.start {
			idxs[e.idx] = struct{}{}
		} else {
			delete(idxs, e.idx)
		}
	}
	intervals = append(intervals, interval{pos, keys(idxs)})
	return &Index{intervals}
}

// At returns the intervals at position i. Returned values are the serial numbers of
// the start-end pairs for which start[n] <= i < end[n].
func (idx *Index) At(i int) []int {
	at := sort.Search(len(idx.idx), func(j int) bool {
		return idx.idx[j].start > i
	})
	if at == 0 {
		return nil
	}
	return cp(idx.idx[at-1].idxs) // Return a copy to keep the index read-only.
}

// A start or an end of an interval.
type event struct {
	idx   int
	pos   int
	start bool
}

// Compares 2 events for sorting.
func eventLess(a, b event) bool {
	if a.pos != b.pos {
		return a.pos < b.pos
	}
	if a.start != b.start {
		return !a.start // End comes before start
	}
	return a.idx < b.idx
}

// The start of a piece with the intervals that intersect with it.
type interval struct {
	start int
	idxs  []int
}

// Returns the keys of a map, sorted.
func keys(m map[int]struct{}) []int {
	if len(m) == 0 {
		return nil
	}
	result := make([]int, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	sort.Ints(result)
	return result
}

// Copies an int slice.
func cp(a []int) []int {
	if a == nil {
		return nil
	}
	result := make([]int, len(a))
	copy(result, a)
	return result
}

package main

import sim "git.saintnet.tech/stryan/spacetea/simulator"

type pagelist []sim.JournalPage

// Len is the number of elements in the collection.
func (e pagelist) Len() int {
	return len(e)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (e pagelist) Less(i int, j int) bool {
	return e[i].ID() < e[j].ID()
}

// Swap swaps the elements with indexes i and j.
func (e pagelist) Swap(i int, j int) {
	tmp := e[i]
	e[i] = e[j]
	e[j] = tmp
}

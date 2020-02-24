// This program is written to find the longest continuous number in a matrix.
// e.g.
// [1, 2, 3, 4]
// [3, 3, 3, 3]
// [3, 2, 2, 2]
// [1, 3, 3, 4]
// In the matrix below, there are maxmum 6 '3' next to each other continuously
// maxmux 3 '2' next to each other continuously, we call the continuous number a 'path'

package main

import (
	"fmt"
	"math"
)

func main() {
	run()
}

func run() {
	var m = matrix([][]int{
		[]int{1, 2, 3, 5},
		[]int{3, 3, 3, 3},
		[]int{3, 2, 2, 2},
		[]int{1, 3, 3, 4},
	})

	uniqueValues := findUniqueValues(m)
	for v := range uniqueValues {
		idxes := m.findIdxes(v)
		paths := findAllPath(idxes)
		path := findLongestPath(paths)
		fmt.Printf("The longest path for value: %v in the matrix is: %v with %v numbers\n", v, path, len(path))
	}
}

type matrix [][]int

type point struct {
	x, y int
}

type x struct{}

// use map of empty struct as set
func findUniqueValues(m matrix) map[int]struct{} {
	uniques := make(map[int]struct{})
	for _, values := range m {
		for _, value := range values {
			uniques[value] = struct{}{}
		}
	}
	return uniques
}

func (m matrix) findIdxes(v int) []point {
	var idxes []point
	for outerIdx, values := range m {
		for innerIdx, value := range values {
			if value == v {
				idxes = append(idxes, point{outerIdx, innerIdx})
			}
		}
	}
	return idxes
}

func findAllPath(ps []point) [][]point {
	var paths [][]point
	var path []point

	for _, p1 := range ps {
		//check if the point has already been evaluated
		if p1.isInPaths(paths) {
			continue
		}
		path = []point{p1}
		path = findOnePath(p1, ps, &path)
		paths = append(paths, path)
	}
	return paths
}

// has to pass pointer of a slice if the function is going to update a slice's length/capcity by passing it to a function
// because the values of the slice are passed by reference (since we pass a copy of the pointer),
// but all the metadata describing the slice itself are just copies
func findOnePath(p1 point, ps []point, path *[]point) []point {
	for _, p2 := range ps {
		if p1.isNextTo(p2) && !p2.isIn(*path) {
			*path = append(*path, p2)
			findOnePath(p2, ps, path)
		}
	}
	return *path
}

func (p1 point) isIn(ps []point) bool {
	for _, p2 := range ps {
		if p2 == p1 {
			return true
		}
	}
	return false
}

func (p1 point) isInPaths(pss [][]point) bool {
	for _, ps := range pss {
		for _, p2 := range ps {
			if p2 == p1 {
				return true
			}
		}
	}
	return false
}

func (p1 point) isNextTo(p2 point) bool {
	if (p1.x == p2.x && int(math.Abs(float64(p1.y-p2.y))) == 1) || (p1.y == p2.y && int(math.Abs(float64(p1.x-p2.x))) == 1) {
		return true
	}
	return false
}

// if 2 path have the longest path at the same time, return the first one
// (from left to right, from up to bottom)
func findLongestPath(paths [][]point) []point {
	longestPath := []point{}
	for _, path := range paths {
		if len(path) > len(longestPath) {
			longestPath = path
		}
	}
	return longestPath
}

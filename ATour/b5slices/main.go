package main

import (
	utl "autilities"
	"slices"

	"github.com/fatih/color"
)

var header = utl.Header{}

func main() {

	header.DisplayHeader("Showing Slices")

	color.Cyan("Prints text in cyan.")

	var arr1 [5]int
	utl.PLine("Array ")
	utl.ShowTypeAndValue(arr1)

	var slcx []string
	utl.PLine("Slice ")
	utl.ShowTypeAndValue(slcx)
	utl.PLine("uninit:", slcx, slcx == nil, len(slcx) == 0)

	slcx = make([]string, 3)
	slcx[0] = "a"
	slcx[1] = "b"
	slcx[2] = "c"
	utl.PLine("set(slcx):", slcx)

	var slc1 = make([]string, 3)
	utl.ShowTypeAndValue(slc1)
	utl.PLine("emp:", slc1, "len:", len(slc1), "cap:", cap(slc1))

	slc1[0] = "a"
	slc1[1] = "b"
	slc1[2] = "c"
	utl.PLine("set:", slc1)
	utl.PLine("get:", slc1[2])

	utl.PLine("len:", len(slc1))

	slc1 = append(slc1, "d")
	slc1 = append(slc1, "e", "f")
	utl.PLine("Append:", slc1)

	c := make([]string, len(slc1))
	copy(c, slc1)
	utl.PLine("Copy:", c)

	l := slc1[2:5]
	utl.PLine("Slice 1:", l)

	l = slc1[:5]
	utl.PLine("Slice 2:", l)

	l = slc1[2:]
	utl.PLine("Slice 3:", l)

	t := []string{"g", "h", "i"}
	utl.ShowTypeAndValue(t)
	utl.PLine("Declare and Initialize:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		utl.PLine("t == t2")
	}

	tdSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		tdSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			tdSlice[i][j] = i + j
		}
	}
	utl.PLine("2d Slice: ", tdSlice)

	ts := make([][]int, 3, 5)
	utl.PLine("ts: ", ts)

	ss := make([]string, 0, 5)
	utl.PLine("SS:", ss, "len:", len(ss), "cap:", cap(ss))

	ss = append(ss, "A", "B")
	utl.PLine("SS:", ss, "len:", len(ss), "cap:", cap(ss))

	ss = append(ss, "C", "D", "E", "F", "G", "H", "I", "J")
	utl.PLine("SS:", ss, "len:", len(ss), "cap:", cap(ss))
}

/*
Notes:

1. Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
2. Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).
3. We can set and get just like with arrays.
4. len returns the length of the slice as expected.
5. In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.
6. Slices can also be copy’d. Here we create an empty slice c of the same length as slc1 and copy into c from slc1.
7. Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements slc1[2], slc1[3], and slc1[4].
8. This slices up to (but excluding) slc1[5].
9. And this slices up from (and including) slc1[2].
10. We can declare and initialize a variable for slice in a single line as well.
11. Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
12. Note that while slices are different types than arrays, they are rendered similarly by utl.PLine.
*/
// Copyright 2013 The Gonum Authors. All rights reserved.
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file

package floats

import (
	"fmt"
)

// Set of examples for all the functions

func ExampleAdd_simple() {
	// Adding three slices together. Note that
	// the result is stored in the first slice
	s1 := []float64{1, 2, 3, 4}
	s2 := []float64{5, 6, 7, 8}
	s3 := []float64{1, 1, 1, 1}

	Add(s1, s2, s3)

	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
	fmt.Println("s3 =", s3)
	// Output:
	// s1 = [7 9 11 13]
	// s2 = [5 6 7 8]
	// s3 = [1 1 1 1]
}

func ExampleAdd_newslice() {
	// If one wants to store the result in a
	// new container, just make a new slice
	s1 := []float64{1, 2, 3, 4}
	s2 := []float64{5, 6, 7, 8}
	s3 := []float64{1, 1, 1, 1}
	dst := make([]float64, len(s1))

	Add(dst, s1, s2, s3)

	fmt.Println("dst =", dst)
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
	fmt.Println("s3 =", s3)
	// Output:
	// dst = [7 9 11 13]
	// s1 = [1 2 3 4]
	// s2 = [5 6 7 8]
	// s3 = [1 1 1 1]
}

func ExampleAdd_unequallengths() {
	// If the lengths of the slices are unknown,
	// use Eqlen to check
	s1 := []float64{1, 2, 3}
	s2 := []float64{5, 6, 7, 8}

	eq := EqLen(s1, s2)
	if eq {
		Add(s1, s2)
	} else {
		fmt.Println("Unequal lengths")
	}
	// Output:
	// Unequal lengths
}

func ExampleAdd_sliceofslicesum() {
	// Add can also be used to take a columnwise sum
	// of a slice of slices

	// First, initialize a slice of slices
	s := make([][]float64, 5)
	for i := range s {
		s[i] = make([]float64, 3)
		for j := range s[i] {
			s[i][j] = float64(j)
		}
	}
	// Remember that the first entry is the destination
	result := make([]float64, len(s[0]))
	Add(result, s...)

	fmt.Println("result =", result)
	// Output:
	// result = [0 5 10]
}

func ExampleAddConst() {
	s := []float64{1, -2, 3, -4}
	c := 5.0

	AddConst(c, s)

	fmt.Println("s =", s)
	// Output:
	// s = [6 3 8 1]
}

func ExampleCumProd() {
	s := []float64{1, -2, 3, -4}
	dst := make([]float64, len(s))

	CumProd(dst, s)

	fmt.Println("dst =", dst)
	fmt.Println("s =", s)
	// Output:
	// dst = [1 -2 -6 24]
	// s = [1 -2 3 -4]
}

func ExampleCumSum() {
	s := []float64{1, -2, 3, -4}
	dst := make([]float64, len(s))

	CumSum(dst, s)

	fmt.Println("dst =", dst)
	fmt.Println("s =", s)
	// Output:
	// dst = [1 -1 2 -2]
	// s = [1 -2 3 -4]
}

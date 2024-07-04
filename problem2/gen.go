package main

import (
	"math/rand"
)

func genRandomSlice(n int) []int {
	res := make([]int, 0, n)
	for j := 0; j < n; j++ {
		res = append(res, rand.Int())
	}

	return res
}

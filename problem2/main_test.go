package main

import (
	"fmt"
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func calculateMaxDifferenceSlow(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	diff := 0
	for i := 1; i < len(arr); i++ {
		diff = max(diff, arr[i]-arr[i-1])
	}

	return diff
}

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 10000},
	{input: 100000},
}

func Test_calculateMaxDifference(t *testing.T) {
	for _, v := range table {
		arr := genRandomSlice(v.input)
		arr1 := slices.Clone(arr)
		arr2 := slices.Clone(arr)

		want := calculateMaxDifferenceSlow(arr1)
		result := calculateMaxDifference(arr2)

		assert.Equal(t, want, result)
	}
}

func Benchmark_calculateMaxDifference(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := genRandomSlice(v.input)
				_ = calculateMaxDifference(arr)
			}
		})
	}
}

func Benchmark_calculateMaxDifferenceSlow(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := genRandomSlice(v.input)
				_ = calculateMaxDifferenceSlow(arr)
			}
		})
	}
}

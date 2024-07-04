package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func calculateUniquePricesSlow(node *CategoryNode) map[int]struct{} {
	if node == nil {
		return nil
	}

	if len(node.Children) == 0 {
		uniquePrices := make(map[int]struct{}, len(node.Prices))
		for _, price := range node.Prices {
			uniquePrices[price] = struct{}{}
		}
		// fmt.Printf("Категория %s: %d уникальных цен\n", node.Name, len(uniquePrices))
		node.UniquePricesCnt = len(uniquePrices)
		return uniquePrices
	}

	uniquePrices := map[int]struct{}{}
	for _, child := range node.Children {
		childUniquePrices := calculateUniquePricesSlow(child)
		mergeSets(uniquePrices, childUniquePrices)
	}

	// fmt.Printf("Категория %s: %d уникальных цен\n", node.Name, len(bigChild))
	node.UniquePricesCnt = len(uniquePrices)
	return uniquePrices
}

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 10000},
}

func Test_calculateUniquePrices(t *testing.T) {
	for _, v := range table {
		root := genRandomCategoryTree(v.input)
		root1 := root.Clone()
		root2 := root.Clone()

		_ = calculateUniquePricesSlow(root1)
		_ = calculateUniquePrices(root2)

		assert.Equal(t, root1, root2)
	}
}

func Benchmark_calculateUniquePrices(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				root := genRandomCategoryTree(v.input)
				_ = calculateUniquePrices(root)
			}
		})
	}
}

func Benchmark_calculateUniquePricesSlow(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				root := genRandomCategoryTree(v.input)
				_ = calculateUniquePricesSlow(root)
			}
		})
	}
}

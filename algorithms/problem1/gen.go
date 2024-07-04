package main

import (
	"math/rand"
)

func genRandomCategoryTree(n int) *CategoryNode {
	//return genBambooRandomCategoryTree(n)
	return genTrueRandomCategoryTree(n)
}

func genBambooRandomCategoryTree(n int) *CategoryNode {
	leavesCnt := rand.Int()%(n-1) + 1
	leaves := make([]*CategoryNode, 0, leavesCnt)
	for i := 0; i < leavesCnt; i++ {
		k := rand.Int()%10 + 1
		prices := make([]int, 0, k)
		for j := 0; j < k; j++ {
			v := rand.Int()
			prices = append(prices, v)
		}
		leaves = append(leaves, &CategoryNode{Prices: prices})
	}

	parentsCnt := n - leavesCnt
	parents := make([]*CategoryNode, 0, parentsCnt)
	root := &CategoryNode{}
	parents = append(parents, root)
	for idx := 1; idx < parentsCnt; idx++ {
		// bamboo
		parentIdx := idx - 1
		parents = append(parents, &CategoryNode{})
		parents[parentIdx].Children = append(parents[parentIdx].Children, parents[idx])
	}

	for idx := 0; idx < leavesCnt; idx++ {
		parentIdx := rand.Int() % parentsCnt
		parents[parentIdx].Children = append(parents[parentIdx].Children, leaves[idx])
	}

	return parents[0]
}

func genTrueRandomCategoryTree(n int) *CategoryNode {
	leavesCnt := n / 2
	leaves := make([]*CategoryNode, 0, leavesCnt)
	for i := 0; i < leavesCnt; i++ {
		k := rand.Int()%10 + 1
		prices := make([]int, 0, k)
		for j := 0; j < k; j++ {
			v := rand.Int()
			prices = append(prices, v)
		}
		leaves = append(leaves, &CategoryNode{Prices: prices})
	}

	parentsCnt := n - leavesCnt
	parents := make([]*CategoryNode, 0, parentsCnt)
	root := &CategoryNode{}
	parents = append(parents, root)
	for idx := 1; idx < parentsCnt; idx++ {
		// true random
		parentIdx := rand.Int() % idx
		parents = append(parents, &CategoryNode{})
		parents[parentIdx].Children = append(parents[parentIdx].Children, parents[idx])
	}

	for idx := 0; idx < leavesCnt; idx++ {
		parentIdx := rand.Int() % parentsCnt
		parents[parentIdx].Children = append(parents[parentIdx].Children, leaves[idx])
	}

	return parents[0]
}

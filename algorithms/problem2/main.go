package main

import (
	"fmt"
	"math"
)

func calculateMaxDifference(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	minV, maxV := arr[0], arr[0]
	for _, v := range arr {
		minV = min(minV, v)
		maxV = max(maxV, v)
	}
	if n <= 2 || minV == maxV {
		return maxV - minV
	}

	bucketMin := make([]int, n-1)
	bucketMax := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		bucketMin[i] = math.MaxInt
		bucketMax[i] = math.MinInt
	}

	bucketSize := float64(maxV-minV) / float64(n-1)
	for _, v := range arr {
		bucketIdx := int(float64(v-minV) / bucketSize)
		if bucketIdx == n-1 {
			bucketIdx--
		}
		bucketMin[bucketIdx] = min(bucketMin[bucketIdx], v)
		bucketMax[bucketIdx] = max(bucketMax[bucketIdx], v)
	}

	prevMaxV, ans := 0, 0
	for i := 0; i < n-1; i++ {
		if bucketMin[i] == math.MaxInt && bucketMax[i] == math.MinInt {
			continue
		}
		if i > 0 {
			ans = max(ans, bucketMin[i]-prevMaxV)
		}
		prevMaxV = bucketMax[i]
	}

	return ans
}

func main() {
	arr := []int{1, 12, 5, 10, 6, 6, 11, 2}
	result := calculateMaxDifference(arr)
	fmt.Println(result)
}

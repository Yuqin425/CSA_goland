package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sort(n []int) {
	m := len(n)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if n[i] < n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	fmt.Print(n)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var nums = make([]int, 0)
	n := rand.Intn(1000)
	for i := 0; i < n; i++ {
		number := rand.Intn(10000)
		nums = append(nums, number)
	}
	sort(nums)
}

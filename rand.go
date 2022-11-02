package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var nums = make([]int, 0)
	for i := 0; i < 100; i++ {
		number := rand.Intn(10000)
		nums = append(nums, number)
	}
	sort.Ints(nums)
	fmt.Print(nums)
}

package main

import (
	"fmt"
	"sort"
)

type number []int

func (n number) Len() int {
	return len(n)
}

func (n number) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n number) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func main() {
	nums := number{5, 1, 1, 2, 0, 0}
	
	fmt.Println("nums=", nums)

	sort.Sort(nums)

	fmt.Println(nums)
}

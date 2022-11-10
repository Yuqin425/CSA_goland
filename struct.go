package main

import "fmt"

type array struct {
	size int
	n    int
	a    []int
}

func main() {
	a := array{}
	fmt.Scanf("%d%d", &a.n, &a.size)
	fmt.Scanln()
	a.a = make([]int, a.n)
	for i := 0; i < a.n; i++ {
		fmt.Scanf("%d", &a.a[i])
	}
	reverse(a)
	for i := 0; i < a.n; i++ {
		fmt.Printf("%d ", a.a[i])
	}
}

func reverse(a array) {
	x := 0
	for i := 0; i < a.size/2+1; i++ {
		for j := 0; j < a.size; j++ {
			if i+j == a.size-1 {
				x = a.a[i]
				a.a[i] = a.a[j]
				a.a[j] = x
			}

		}
	}
}

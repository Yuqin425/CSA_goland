package main

import "fmt"

func main() {
	var (
		a  int
		b  int
		ch string
	)
	fmt.Scanf("%d %s %d", &a, &ch, &b)
	switch ch {
	case "+":
		fmt.Println(a, ch, b, "=", a+b)
	case "-":
		fmt.Println(a, ch, b, "=", a-b)
	case "*":
		fmt.Println(a, ch, b, "=", a*b)
	case "/":
		fmt.Println(a, ch, b, "=", a/b)
	case "%":
		fmt.Println(a, ch, b, "=", a%b)
	default:
		fmt.Println("It is error!")
	}
}

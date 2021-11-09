package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println("N1 + N2 = ", sum(30, 15))
}

func sum(n1 int, n2 int) int {
	return n1 + n2
}

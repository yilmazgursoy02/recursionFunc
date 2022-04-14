package main

import "fmt"

func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func main() {

	fmt.Println(factorial(5))

	fmt.Println(fibonaci(5))
	for i := 0; i < 6; i++ {

		fmt.Printf("%d ", fibonaci(i))
	}

}

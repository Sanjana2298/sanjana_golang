// Write program to find fibonacci series upto specific number
package main

import "fmt"

func generateFibonacci(n int) []int {
	fibSeries := []int{0, 1}
	for i := 2; ; i++ {
		next := fibSeries[i-1] + fibSeries[i-2]
		if next > n {
			break
		}
		fibSeries = append(fibSeries, next)
	}
	return fibSeries
}

func main() {
	var n int
	fmt.Print("Enter a number to generate the Fibonacci series up to: ")
	fmt.Scan(&n)

	fibSeries := generateFibonacci(n)

	fmt.Printf("Fibonacci series up to %d: %v\n", n, fibSeries)
}

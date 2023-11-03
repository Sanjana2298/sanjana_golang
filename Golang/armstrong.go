// GoLang program to find the given number is Armstrong or not
// using "for" loop.

package main

import "fmt"

func main() {
	var num int = 0
	var rem int = 0
	var res int = 0
	var tmp int = 0

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)

	tmp = num
	for num > 0 {
		rem = num % 10
		res = res + (rem * rem * rem)
		num = num / 10

	}

	if tmp == res {
		fmt.Printf("Number is armstrong")
	} else {
		fmt.Printf("Number is not armstrong")
	}
}

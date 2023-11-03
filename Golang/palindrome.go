// Write a program to check number is palindrome or not
package main

import (
	"fmt"
)

func isPalindrome(number int) bool {
	originalNumber := number
	reversedNumber := 0

	for number > 0 {
		remainder := number % 10
		reversedNumber = reversedNumber*10 + remainder
		number /= 10
	}

	return originalNumber == reversedNumber
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if isPalindrome(number) {
		fmt.Println("It's a palindrome!")
	} else {
		fmt.Println("It's not a palindrome.")
	}
}

package main

import (
	"fmt"
)

func main() {
	var day int
	fmt.Print("Please enter a number between 1 and 7: ")
	fmt.Scan(&day)

	var dayOfWeek string

	switch day {
	case 1:
		dayOfWeek = "Sunday"
	case 2:
		dayOfWeek = "Monday"
	case 3:
		dayOfWeek = "Tuesday"
	case 4:
		dayOfWeek = "Wednesday"
	case 5:
		dayOfWeek = "Thursday"
	case 6:
		dayOfWeek = "Friday"
	case 7:
		dayOfWeek = "Saturday"
	default:
		dayOfWeek = "Invalid input. Please enter a number between 1 and 7."
	}

	fmt.Printf("The day of the week is: %s\n", dayOfWeek)
}

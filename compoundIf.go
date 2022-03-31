package main

import "fmt"

func main() {
	//Insert your code here
	var year int
	fmt.Println("Enter the year: ")
	fmt.Scanln(&year)

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("It's a leapyear!")
	} else {
		fmt.Println("It's not a leapyear.")
	}
}

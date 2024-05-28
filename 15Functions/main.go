package main

import "fmt"

func main() {
	fmt.Println("Functions in golang!")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result of calculation: ", result)

	proResult, myMessage := proAdder(2, 4, 5, 6, 7, 8)
	fmt.Println("Pro result is: ", proResult)
	fmt.Println("My message is: ", myMessage)
}

func adder(valOne int, valTwo int) /*what type of value to enter*/ int /*signature i.e what type of value to return*/ {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi, Pro result function"
}

func greeter() {
	fmt.Println("Namastey from golang")
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input."
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for our Pizza: ")

	// comma ok syntax / error ok syntax
	// input(string), err(when we have to catch a error) / _(when we dont have to catch any error)
	input, _ := reader.ReadString('\n') //this \n highlights the fact that it will read untill new line(\n) is seen
	fmt.Println("Thanks for your rating!", input)
	fmt.Printf("Type of the rating is %T", input)

}

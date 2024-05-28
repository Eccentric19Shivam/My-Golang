package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang!")

	/*
		var *name* [*index*]data_type
		  To add :
		  	name[index] = value
	*/

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Pineapple"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "mushroom", "beans"}
	fmt.Println("Vegetable list is: ", vegList)
	fmt.Println("The length of vegetable list is: ", len(vegList))
}

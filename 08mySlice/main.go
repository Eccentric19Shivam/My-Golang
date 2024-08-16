package main

import (
	"fmt"
	// "sort"
)

func main() {
	fmt.Println("Welcome to Slices!")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruit list %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3]) //last range is alway non-inclusive
	fmt.Println(fruitList)

	// highScore := make([]int, 4)

	// highScore[0] = 234
	// highScore[1] = 945
	// highScore[2] = 465
	// highScore[3] = 867 // base memory allocation
	// // highScore[4] = 777

	// highScore = append(highScore, 555, 888, 999) // allocate more memory

	// fmt.Println(highScore)

	// fmt.Println(sort.IntsAreSorted((highScore)))

	// sort.Ints(highScore)
	// fmt.Println(highScore)
	// fmt.Println(sort.IntsAreSorted((highScore)))

	// // How to remove a value from slices based on index

	// var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	// fmt.Println(courses)
	// var index int = 2
	// courses = append(courses[:index], courses[index+1:]...)
	// fmt.Println(courses)

}

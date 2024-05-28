package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in golang!")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough // niche jao
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough // niche jao
	case 5: // ab mt jao
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and you can roll again")
	default:
		fmt.Println("What was that?")
	}
}

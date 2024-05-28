package main

import "fmt"

func main() {
	fmt.Println("Loops in golang!")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }
	// for _, day := range days {
	// 	fmt.Printf("index is and value is %v\n", day)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco // when 2 is reached, the message in the label is declared
		}

		if rougeValue == 5 {
			rougeValue++
			continue // 5 ko skip krke aur phir aage continue
			//break 5 aur 5 ke baad nhi dikhega
		}

		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeOnline.in")
}

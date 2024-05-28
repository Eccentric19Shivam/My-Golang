package main

import "fmt"

const LoginToken string = "ihdfkgjsb" //it is public or global variable as the var name declared with capital letter in the front

// jwtToken := 30000.0 walrus operator not allowed outside if the method, means main fucntion is the main method to write anything, so we use var in outside the method

func main() {
	var username string = "shivam"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float64 = 256.455748957593856848
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//no var style

	numberofUser := 30000.0
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}

package main

import "fmt"

//outside method not allowed
//jwtToken:=300000

// Public as name starts with capital letter
const LoginToken string = "anything"

func main() {
	var username string = "hitesh"
	//fmt.Println("Variables")
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLogged bool = false
	fmt.Println(isLogged)
	fmt.Printf("Variable is of type: %T \n", isLogged)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.4556463563
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default Values
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit types
	var website = "learn code"
	fmt.Println(website)

	//no var style-->Allowed inside method
	numberOfUsers := 300000.0
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)

}

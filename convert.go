package main

import "fmt"

func main() {
	yesNo := "No"

	var launch bool

	switch yesNo {
	case "True", "Yes", "1":
		launch = true
	case "False", "No", "0":
		launch = false
	default:
		fmt.Println("I don't know")
	}
	fmt.Println(yesNo, " is ", launch)

}

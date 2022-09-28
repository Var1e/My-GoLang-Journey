package main

import "fmt"

func main() {
	//assigning administrator with string pointer
	//any string can point to administrator
	var administrator *string

	scolese := "Cristopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)

	bolden = "Charles F. Bolden Jr."
	fmt.Println(*administrator)

	//changes to bolden can be made in one place
	//coz administrator points to bolden instead of
	//making a copy
	*administrator = "Maj. Gen. Charles F. Bolden Jr."
	fmt.Println(bolden)

	//assigning major to administrator creates a new pointer
	// that points to the same address(bolden)
	major := administrator
	*major = "Major General Charles F. Bolden Jr."
	fmt.Println(bolden)

	//this prints true coz major and admin both points to the same address
	fmt.Println(major == administrator)

	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot
	fmt.Println(major == administrator)

	//assigning dereferenced value of major to another variable
	//creates a clone of major changes to major
	//doesn't affect bolden
	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles)
	fmt.Println(bolden)

	charles = "Charles Bolden"

	//this prints true coz charles and bolden variable
	//both contain the same strings
	fmt.Println(charles == bolden)

	//this prints falls coz charles and bolden is not
	//in the same address
	fmt.Println(&charles == &bolden)

}

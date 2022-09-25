package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
func celsiusToFahrenheit(c float64) float64 {
	c = (c * 9.0 / 5.0) + 32.0
	return c
}
func kelvinToFahrenheit(f float64) float64 {
	f -= 459.67
	return f
}

func main() {
	temp := 294.5

	celsius := kelvinToCelsius(temp)
	fmt.Println("273.5 Kelvin is ", celsius, " Celsius")

	tempe := 100.0
	fahrenheit := celsiusToFahrenheit(tempe)
	fmt.Println(tempe, " is", fahrenheit, " Fahrenheit")

	tempra := 0.0
	foreignheight := kelvinToFahrenheit(tempra)
	fmt.Println(tempra, " is", foreignheight, " Fahrenheit")

}

package main

import "fmt"

type Planets []string

func (planets Planets) terraform() {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func main() {
	planets := []string{"Mercury", "Venus", "Earth", "Jupiter", "Saturn", "Uranus", "Neptune"}
	Planets(planets[0:1]).terraform()
	Planets(planets[2:3]).terraform()
	fmt.Println(planets)

}

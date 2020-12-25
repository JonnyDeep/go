package main

import "fmt"

type location struct {
	lat  float64
	long float64
}

func distance(loc1, loc2 location) dis {
	return dis{0.0, 0.0}
}

type dis struct {
	lat  float64
	long float64
}

func main() {
	// var planets [8]string

	// planets[0] = "Mercury"
	// planets[1] = "Venus"
	// planets[2] = "Earth"

	// earth := planets[2]
	// fmt.Println(earth)

	// fmt.Println(len(planets))
	// fmt.Println(planets[3] == "")

	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.47263

	var earth location
	earth.lat = 12.999
	earth.long = -128.2323

	distance(spirit, earth)

	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth"}

	fmt.Println(planets)

	for i, planet := range planets {
		fmt.Println(i, planet)

	}

}

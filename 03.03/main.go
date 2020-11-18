package main

import "fmt"

func kelvinToCelSius(k float64) float64 {
	k -= 273.15
	return k
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelSius(kelvin)
	fmt.Print(kelvin, " k is ", celsius)
}

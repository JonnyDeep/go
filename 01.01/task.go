package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Spaceline\t\t\tDays\tTrip type\tPrice")
	const distance = 62100000
	var days = 0
	for i := 0; i < 10; i++ {
		space, trip, speed, price := rand.Intn(3)+1, rand.Intn(2)+1, rand.Intn(14)+16, rand.Intn(14)+36
		days = distance / (speed * 24 * 60 * 60)
		var trip_type = ""
		if trip == 1 {
			trip_type = "Round trip"
		} else {
			trip_type = "one way"
		}
		switch space {
		case 1:
			if trip_type == "one way" {
				fmt.Printf("%v\t\t%v\t%v\t\t%v\n", "Space Adventures", days, trip_type, price)
			} else {
				fmt.Printf("%v\t\t%v\t%v\t%v\n", "Space Adventures", days, trip_type, price)
			}
		case 2:
			if trip_type == "one way" {
				fmt.Printf("%v\t\t\t\t%v\t%v\t%v\n", "Spacex", days, trip_type, price)
			} else {
				fmt.Printf("%v\t\t\t\t%v\t%v\t%v\n", "Spacex", days, trip_type, price)
			}
		case 3:
			if trip_type == "one way" {
				fmt.Printf("%v\t\t\t%v\t%v\t\t%v\n", "Virgin Galactic", days, trip_type, price)
			} else {
				fmt.Printf("%v\t\t\t%v\t%v\t%v\n", "Virgin Galactic", days, trip_type, price)
			}
		}
	}
}

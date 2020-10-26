package main

import (
	"fmt"
	"math/big"
)

func main() {
	// const lightSpeed = 299792
	// const secondsPerDay = 86400

	// var distance int64 = 41.3e12
	// fmt.Println("Alpha Centauri is", distance, "km away")
	// days := distance / lightSpeed / secondsPerDay
	// fmt.Println("That is", days, "days of travel at light speed.")

	//超过10e18 的整数big.Int
	//big.Float
	//big.Rat

	lightSpeed := big.NewInt(299792)
	secondsPerday := big.NewInt(86400)

	distance := new(big.Int)
	distance.SetString("24000000000000000000000000", 10)

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerday)
	fmt.Println("That is", days, "days of travel at light speed.")

	fmt.Println("the distance is", 236000000000000000/299792/31536000)
}

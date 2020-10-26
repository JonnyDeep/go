package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var sum float32 = 0

	for sum < 20 {
		var i = rand.Intn(3) + 1
		if i == 1 {
			sum += 0.05
		} else if i == 2 {
			sum += 0.10
		} else {
			sum += 0.25
		}
		fmt.Printf("储钱罐余额：%5.2f\n", sum)
	}

}

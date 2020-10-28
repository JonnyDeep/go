package main

import (
	"fmt"
	"math"
)

func main() {
	v := 42
	if v >= 0 && v <= math.MaxUint8 {
		v8 := uint8(v)
		fmt.Println("converted:", v8)
	}

	
	//想把数值转化为string，他的值必须能转化为code point
	//strconv.Itoa
	//strconv.Atoi

	
}

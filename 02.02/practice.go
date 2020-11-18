package main

import (
	"fmt"
	"strconv"
)

func main() {
	var yesno bool

	var text1 string = "string"
	var text2 string = "true"

	yesno, err := strconv.ParseBool(text2)
	if err != nil {
		fmt.Println("converted:", err)
	} else {
		fmt.Println("converted:", yesno)
	}

	ans, err := strconv.ParseBool(text1)
	if err != nil {
		fmt.Println("converted:", err)
	} else {
		fmt.Println("converted:", ans)
	}

}

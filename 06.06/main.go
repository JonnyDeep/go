package main

import "fmt"

func main() {
	answer := 12
	fmt.Println(&answer)

	address := &answer

	fmt.Println(*address)
}

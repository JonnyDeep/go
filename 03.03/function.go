package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v K\n", k)
		time.Sleep(time.Second)
	}
}
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

//闭包就是由于匿名函数封闭并包围作用域钟的变量而得名的
func main() {
	sensor := fakeSensor
	fmt.Println(sensor())
	sensor = realSensor
	fmt.Println(sensor())
	measureTemperature(3, fakeSensor)

	f := func(message string) {
		fmt.Println(message)
	}
	f("hello")

	func() {
		fmt.Println("funct")
	}()
}

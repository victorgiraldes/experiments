package main

import (
	"fmt"
)

type Car struct {
	color string
	engine string
	brandy string
}

func myCar(color string, engine string) {
	fmt.Println("Cor: ", color)
	fmt.Println("Motor: ", engine)
}

func 

func main() {
	// var (
	// 	buf    bytes.Buffer
	// 	logger = log.New(&buf, "logger: ", log.Lshortfile)
	// )

	// logger.Print("Error")

	// fmt.Print(&buf)

	var car Car

	car.color = "black"
	car.engine = "1.6"

	myCar(car.color, car.engine)
}

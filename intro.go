package main

import (
	"fmt"
)

func main() {

	//Write a program to determine how fast a ship would need to travel (in km/h)
	//in order to reach Malacandra in 28 days. Assume a distance of 56,000,000 km.

	//Speed = Distance/Time (km/hr)
	// var t = 28 * 24   //converts days to hrs
	// var d = 56000000  //Km
	// var speed = d / t // Km/hr

	// fmt.Println(speed)

	t, d, speed := 28*24, 56000000, 0
	speed = d / t
	fmt.Println(speed)

}

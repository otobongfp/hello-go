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

	//or method 2

	// t, d, speed := 28*24, 56000000, 0
	// speed = d / t
	// fmt.Println(speed)

	// var command = "Get Free"
	// exit := strings.Contains(command, "Free")
	// fmt.Println("You are free now!!!", exit)
	//------------------------------------------------

	// COMPARISONS
	// var age = 41
	// var toDrink = age > 18
	// fmt.Printf("I am %v, am I eligible for a drink? %v \n", age, toDrink)
	//------------------------------------------------
	//CALCULATE THE LEAP YEARS
	// fmt.Println("Is 2028 a Leap Year?")
	// var leap = 2028/400 == 0 || 2028%4 == 0
	// if leap {
	// 	fmt.Println("Yes it is a leap year")
	// } else {
	// 	fmt.Println("No it is Not!!!")
	// }

	var haveTorch = true
	var litTorch = false
	if !haveTorch || !litTorch {
		fmt.Println("Nothing to see here.")
	}

}

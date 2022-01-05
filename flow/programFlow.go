package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var result string;

	if i := 38; i < 0 {
		result = "Less than 0"
	} else if i == 0 {
		result = "Equal than 0"
	} else {
		result = "Greater than 0"
	}

	fmt.Println("If condition result: ", result)

	rand.Seed(time.Now().Unix())
	day := rand.Intn(8) + 1;
	switch day{
		case 1:
			result = "Sunday"
		case 2:
			result = "Monday"
		case 8:
			fallthrough
		default:
			result = "Unknown day"
	}

	fmt.Printf("Switch case, Day %v is: %v\n", day, result)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println("Print with normal 'for' loop (3 arguments)")
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	fmt.Println("Print with range 'for' loop (2 arguments)")
	for i := range colors {
		fmt.Println(colors[i])
	}

	fmt.Println("Print with range  and foreach loop (2 arguments)")
	for _, colors := range colors {
		fmt.Println(colors)
	}

	val := 0
	for val < 10 {
		fmt.Println("Value: ", val)
		if val >= 1 {
			//break
			goto theBinary
		}
		val++
	}

theBinary:
	fmt.Println("The Binary digits")
}

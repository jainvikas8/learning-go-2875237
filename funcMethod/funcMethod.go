package main

import (
	"fmt"
)

// Dog is a struct
type Dog struct {
	Breed string
	Weight int
	Sound string
	bark string
}

func main() {

	callFunc1()

	fmt.Println("Sum of 1 and 2 is: ", addValues(1,2))

	fmt.Println("Sum of 1, 2 and 3 is: ", addArrayValues(1,2,3))

	total, count := addArrayValuesCount(1,2,3,4,5,6,7,8)
	fmt.Println("Sum of 1, 2, 3, 4, 5, 6, 7 and 8 is: ", total)
	fmt.Println("Count is: ", count)

	poodle := Dog {"Poodle", 10, "Woof!", ""}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
	poodle.Bark()
}

func callFunc1() {
	fmt.Println("In callFunc1()")
}

func addValues(val1 int, val2 int) int {
	return val1 + val2
}

func addArrayValues(values ...int) int {
	total := 0
	for v := 0; v < len(values); v++ {
		total += values[v]
	}
	return total
}

func addArrayValuesCount(values ...int) (int, int) {
	total := 0
	for v := 0; v < len(values); v++ {
		total += values[v]
	}
	return total, len(values)
}

// Speak is how dog speaks
func (d Dog) Speak() {
	fmt.Println("This dog speaks as:", d.Sound)
}

// Bark is how dog barks
func (d Dog) Bark() {
	d.bark = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println("This dog barks as:", d.bark)
}

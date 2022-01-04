package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 1, 2, 3
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 1.1, 2.1, 3.1
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)

	fmt.Println("Float sum using math.Round operation:", (math.Round(floatSum*100)/100))

	fmt.Printf("Float sum using formatting: %.2f\n", floatSum)

	addedSum := floatSum + float64(intSum)

	fmt.Printf("Added sum using formatting: %.2f\n", addedSum)

	fmt.Printf("Added sum * math.pi: %.2f\n", (addedSum * math.Pi))
}

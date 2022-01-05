package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
)

func main() {
	fmt.Println("---Simple Calculator---")
	fmt.Print("Please enter value 1: ")
	float1 := readInput()

	fmt.Print("Please enter value 2: ")
	float2 := readInput()

	fmt.Println("Please select an operation")
	fmt.Println("1. Add (+)")
	fmt.Println("2. Subtract (-)")
	fmt.Println("3. Divide (/)")
	fmt.Println("4. Multiplication (*)")
	fmt.Print("Input: ")

	switch readInput() {
		case 1:
			printVal(float1, float2, (float1 + float2), "sum")
		case 2:
			printVal(float1, float2, (float1 - float2), "subtraction")
		case 3:
			printVal(float1, float2, (float1 / float2), "division")
		case 4:
			printVal(float1, float2, (float1 * float2), "multiplication")
		default:
			fmt.Println("Unsupported option")
	}
}

func readInput() float64 {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	floatVal, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}

	return floatVal
}

func printVal(val1 float64, val2 float64, result float64, operation string) {
	fmt.Printf("The %v of %v and %v is ", operation, val1, val2)

	if result == float64(int(result)) {
		fmt.Printf("%v\n", result)
	} else {
		fmt.Printf("%.4f\n", result)
	}
}

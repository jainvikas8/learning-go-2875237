package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter number: ")
	number_input, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(number_input), 64)
	if err != nil {
		fmt.Println(err)
		panic("This should be a number")
	} else {
		fmt.Println("Value of number: ", aFloat)
	}
}

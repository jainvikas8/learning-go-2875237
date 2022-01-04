package main

import (
	"fmt"
)

const aConst float64 = 123.45

func main() {
	var aString string = "This is Gopher"
	fmt.Print(aString)
	fmt.Printf(", The variable type is %T\n", aString)

	var i int = 42
	fmt.Print(i)
	fmt.Printf(", The variable type is %T\n", i)

	var j uint
	fmt.Print(j)
	fmt.Printf(", The variable type is %T\n", j)

	var newString = "This is another Gopher"
	fmt.Print(newString)
	fmt.Printf(", The variable type is %T\n", newString)

	var newInt = 32
	fmt.Print(newInt)
	fmt.Printf(", The variable type is %T\n", newInt)

	anotherString := "This is another Gopher"
	fmt.Print(anotherString)
	fmt.Printf(", The variable type is %T\n", anotherString)

	fmt.Print(aConst)
	fmt.Printf(", The variable type is %T\n", aConst)
}


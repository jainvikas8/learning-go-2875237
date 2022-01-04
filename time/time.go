package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("The time is - ", n)

	t := time.Date(2022, time.January, 4, 21, 8, 23, 0, time.UTC)
	fmt.Println("I'm learning go language on - ", t)
	fmt.Println("ANSIC formated - ", n.Format(time.ANSIC))
}

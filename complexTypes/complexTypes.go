package main

import (
	"fmt"
	"sort"
)

// Pets is a struct
type Pets struct{
	whichType string
	Weight int
}

func main() {
	var ptr *int
	i := 38
	fmt.Println("Value of ptr (pointer) before assignment:", ptr)
	ptr = &i
	fmt.Println("Value of ptr (pointer) after assignment:", ptr)
	fmt.Println("Value at the address pointed by ptr (pointer) after assignment:", *ptr)

	*ptr *= 2
	fmt.Println("Value of ptr after multiplying by 2:", *ptr)
	fmt.Println("Value of i:", i)

	var colors [3] string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println("Array of colors:", colors)

	var numbers = [5]int{1, 2, 3, 5, 8}
	fmt.Println("Array of numbers:", numbers)

	fmt.Println("Array length of colors:", len(colors))
	fmt.Println("Array length of numbers:", len(numbers))

	var newColors = []string{"Yellow", "Red", "Green", "Blue"}
	fmt.Println("Slice of colors:", newColors)

	newColors = append(newColors, "Purple")
	fmt.Println("Slice of colors (appended):", newColors)
	fmt.Println("Slice length of colors:", len(newColors))

	newColors = append(newColors[1:len(newColors)])
	fmt.Println("Slice of colors (remove first item):", newColors)
	fmt.Println("Slice length of colors", len(newColors))

	newNumbers := make([]int, 2)
	fmt.Println("Slice of numbers using make():", newNumbers)

	newNumbers = append(newNumbers, 12)
	fmt.Println("Slice of numbers (appended):", newNumbers)
	fmt.Println("Slice length of numbers:", len(newNumbers))

	newNumbers[1] = 56
	fmt.Println("Slice of numbers (modified):", newNumbers)

	sort.Ints(newNumbers)
	fmt.Println("Slice of numbers (sorted):", newNumbers)

	stations := make(map[string]string)
	stations["RYS"] = "Royston"
	stations["CMB"] = "Cambridge"
	stations["LGC"] = "Letchworth Garden City"
	fmt.Println("Map of stations:", stations)

	royston := stations["RYS"]
	fmt.Println("RYS is ", royston)

	delete(stations, "LGC")
	fmt.Println("Map of stations (deleted LGC):", stations)
	fmt.Println("Map length:", len(stations))

	stations["ELY"] = "ELY"
	fmt.Println("Map of stations (updated):", stations)

	//Create a slice to copy the keys from the map
	keys := make([]string, len(stations))
	itr := 0
	for k := range stations {
		keys[itr] = k
		itr++
	}
	fmt.Println("Slice of keys from the Map of stations", keys)
	sort.Strings(keys)
	fmt.Println("Slice of keys from the Map of stations (sorted)", keys)

	for k := range keys {
		fmt.Printf("%v: %v\n", keys[k], stations[keys[k]])
	}

	dog := Pets{"dog", 10}
	fmt.Println("Struct of Pets", dog)
	fmt.Printf("Displaying fields of struct : %+v\n", dog)
	fmt.Printf("Weight of the dog: %v\n", dog.Weight)
}

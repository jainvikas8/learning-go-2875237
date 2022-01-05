package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello gopher!"

	file, err := os.Create("./newFile.txt")
	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close()
	defer readFile("./newFile.txt")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Printf("Reading file: %v\n%v", fileName, string(data))
}

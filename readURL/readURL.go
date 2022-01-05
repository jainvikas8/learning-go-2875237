package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const url = "http://google.co.uk"

func main() {
	response, err := http.Get(url)
	checkError(err)

	fmt.Printf("Response type: %T\n", response)
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	checkError(err)

	fmt.Printf("Reading file: %v", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

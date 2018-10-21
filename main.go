package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//making a byte slice
	//make is a built in function that takes a type of slice and the second argument is empty space that
	//you want the slice to be initialized with. if we ,make it 0, Read function will assume that the slice is already full. so 99999 is kind of a guess about the empty space that would be needed for the slice
	bs := make([]byte, 99999)

	//the response body takes the byte slice and passes it to the Read function
	resp.Body.Read(bs)

	fmt.Println(string(bs))

}

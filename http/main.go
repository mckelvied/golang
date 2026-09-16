package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error 1:", err)
		return
	}
	defer resp.Body.Close()

	// Process the response...
	bs := make([]byte, 99999)
	n, err := resp.Body.Read(bs)
	if err != nil {
		fmt.Println("Error 2:", err)
		return
	}
	fmt.Println("Read bytes:", n)
	fmt.Println(string(bs[:n]))

}

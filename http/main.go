package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	website := "https://google.fr/"
	resp, err := http.Get(website)
	if err != nil {
		fmt.Printf("Can't check website %s\n", website)
		os.Exit(1)
	}

	// Read body of resp
	// takes a type of a slice and number of empty space initialized
	// responseBody := make([]byte, 99999)
	// resp.Body.Read(responseBody)
	// fmt.Println(string(responseBody))

	// Go helper to do what we've done previously with the slice of byte 99999 ...
	// Because it's use the Stdout helper, this is going to print the result

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}

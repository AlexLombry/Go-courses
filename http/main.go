package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	website := "https://google.fr/"
	resp, err := http.Get(website)
	if err != nil {
		fmt.Printf("Can't check website %s\n", website)
		os.Exit(1)
	}
	fmt.Printf("%+v", resp)
}

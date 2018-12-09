package main

import (
	"fmt"
	"net/http"
	"time"
)

// In this topic, channels is use to communicate between go routine
// to know if they are completed their task
// channels is the only way to communicate with go routines
func main() {
	links := getLinks()

	// create a brand new channel
	c := make(chan string)

	for _, link := range links {
		// go routine to avoid blocking call from http.Get
		// go child routine
		go checkLink(link, c)
	}

	// Print from the channel
	//fmt.Println(<-c)

	// Infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }
	for l := range c {
		// function literal (anonymous function)
		// Be aware of L variable is outer scope of func
		// Point in the same memory address
		// We have to pass it
		// link and l is the same variable in this func literal
		go func(link string) {
			time.Sleep(time.Second * 2)
			checkLink(link, c)
		}(l)
	}
}

func getLinks() []string {
	return []string{
		"https://laracasts.com",
		"https://golang.org",
		"http://localhost:10000/a.php",
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down !")
		c <- link
		return
	}
	// print into the channel
	fmt.Println(link, "is up !")
	c <- link
}

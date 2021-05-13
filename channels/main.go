package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.yahoo.com",
		"http://www.amazon.com",
		"http://edjx.io",
		"https://timesofindia.indiatimes.com",
		"https://golang.org",
	}

	c := make(chan string) // make a channel to communicate for go routine

	// access each link via go routine
	for _, link := range links {
		go checklink(link, c) // go routine to call function
	}

	// run loop everytime channel emmits a value
	for l := range c {
		fmt.Print("inside channel -->", l)
		// it is a function lietral(anoynymous function) acting as go routine
		// you don't access main variables as in child routines directly
		// pass and access should be the approach
		go func(link string) {
			checklink(link, c)
			time.Sleep(5 * time.Second)
		}(l) //pass an argument in paranthesis so that func can recieve it
	}
	fmt.Println()
}

func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "--> might be down, error --> ", err)
		c <- link // send link back to channel, it act as a block call
		return
	}
	c <- link
	fmt.Println(link, "link is up")
}

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	links := []string{"http://google.com", "http://facebook.com", "http://golang.org"}

	// create channel which will wait all go routines to finish, and then finish executing main routine
	c := make(chan string)

	for _, website := range links {
		fmt.Println(website)
		go checkLink(website, c) // creating go routines, not using code syncronosly
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) // blocking line of code, waits for the message
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		os.Exit(1)
	}

	c <- "Received" // send message to channel
	fmt.Println(link, "is up!")

}

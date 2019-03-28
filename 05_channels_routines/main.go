package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	links := []string{"http://google.com", "http://facebook.com", "http://golang.org"}

	// create channel which will wait all go routines to finish, and then finish executing main routine
	c := make(chan string)

	for _, website := range links {
		fmt.Println(website)
		go checkLink(website, c) // creating go routines, not using code syncronosly
	}

	// for {
	// 	// fmt.Println(<-c) // blocking line of code, waits for the message
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		fmt.Println("-------------------")
		go func(l string) {
			time.Sleep(time.Second * 5)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		os.Exit(1)
	}

	fmt.Println(link, "is up!")
	c <- link // send message to channel

}

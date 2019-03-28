package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	links := []string{"http://google.com", "http://facebook.com", "http://golang.org"}

	for _, website := range links {
		fmt.Println(website)
		go checkLink(website)
	}
}

func checkLink(link string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		os.Exit(1)
	}

	fmt.Println(link, "is up!")

}

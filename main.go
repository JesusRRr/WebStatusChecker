package main

import "fmt"

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	for _, link := range links {
		fmt.Println(link)
	}
}

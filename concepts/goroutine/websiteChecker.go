package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	siteList := []string{
		"http://www.google.com",
		"http://facebook.com",
		"http://yahoo.com",
		"http://example.com",
	}

	c := make(chan string)

	for _, url := range siteList {
		go checkSite(url, c)
	}
	for l := range c {
		func(il string) {
			time.Sleep(5 * time.Second)
			go checkSite(il, c)
		}(l)
	}
}

func checkSite(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, " site is down")
		c <- url
	} else {
		fmt.Println(url, " site is up")
		c <- url
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("go get google")
	fmt.Print("hello world ")
	res, err := http.Get("http://www3.google.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
	if err != nil {
		log.Fatal(err)
	}
}

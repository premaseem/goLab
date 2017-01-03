package main

import  (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)


func main() {
	fmt.Println("go get google")
	res, err := http.Get("http://www3.google.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s",page)
	if err != nil {
		log.Fatal(err)
	}
}

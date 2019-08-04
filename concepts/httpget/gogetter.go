package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("go get google")
	res, err := http.Get("http://www.google.com")

	// 1. Read using utils
	//page, _ := ioutil.ReadAll(res.Body)
	//res.Body.Close()
	//fmt.Printf("%s", page)

	// 2. Read using Reader
	//data := make([]byte,99999)
	//res.Body.Read(data)
	//fmt.Println(string(data))

	// 3. write to standout
	//io.Copy(os.Stdout, res.Body)

	// 4. custom write implementation
	myWriter := custmoWriter{}
	io.Copy(myWriter, res.Body)

	if err != nil {
		log.Fatal(err)
	}
}

type custmoWriter struct{}

func (custmoWriter) Write(data []byte) (int, error) {
	fmt.Println(string(data))
	return len(data), nil
}

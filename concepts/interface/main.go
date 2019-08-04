package main

import "fmt"

type bot interface {
	getGreeting() string
}

type chat interface {
	chitCat() string
}

type chatbot interface {
	chat
	bot
}

type engBot struct{}
type spBot struct{}

func print(b chatbot) {
	fmt.Println(b.getGreeting())
}

func (engBot) getGreeting() string {
	return "hello"
}

func (spBot) getGreeting() string {
	return "hola"
}

func (engBot) chitCat() string {
	return "hello"
}

func (spBot) chitCat() string {
	return "hola"
}

func main() {

	sb := spBot{}
	eb := engBot{}

	print(sb)
	print(eb)

}

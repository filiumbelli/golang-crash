package main

type bot interface {
	getGreetings() string
}
type englishBot struct{}
type spanishBot struct{}

func (e englishBot) getGreetings() string {
	return "Hello"
}

func (s spanishBot) getGreetings() string {
	return "Ola"
}

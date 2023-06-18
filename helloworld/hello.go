package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello() string {
	return englishHelloPrefix + "world"
}

func HelloPerson(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello())
	fmt.Println(HelloPerson("Talha"))
}

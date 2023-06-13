package main

import "fmt"

func Hello() string {
	return "Hello, world"
}

func HelloPerson(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello())
	fmt.Println(HelloPerson("Talha"))
}

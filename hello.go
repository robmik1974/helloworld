package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const polish = "Polish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const polishHelloPrefix = "Witaj, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case polish:
		prefix = polishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}

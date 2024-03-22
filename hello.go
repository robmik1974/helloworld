package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	polish             = "Polish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	polishHelloPrefix  = "Witaj, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingsPrefix(language) + name
}

func greetingsPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case polish:
		prefix = polishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}

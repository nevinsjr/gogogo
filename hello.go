package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	worldDefault       = "World"
)

func selectGreetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// Hello TODO: overload this function
func Hello(name string, language string) string {
	validatedName := worldDefault
	if name != "" {
		validatedName = name
	}

	return selectGreetingPrefix(language) + validatedName + "!"
}

func main() {
	fmt.Println(Hello("John", ""))
}

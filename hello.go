package main

import "fmt"

const (
	EN_PREFIX = "Hello, "
	SP_PREFIX = "Hola, "
	FR_PREFIX = "Bonjour, "

	SPANISH = "Spanish"
	FRENCH  = "French"
)

func greetingPrefix(language string) (prefix string) {

	switch language {
	case FRENCH:
		prefix = FR_PREFIX
	case SPANISH:
		prefix = SP_PREFIX
	default:
		prefix = EN_PREFIX
	}

	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	return fmt.Sprintf("%s%s", prefix, name)
}

func main() {
	fmt.Println(Hello("Brain", ""))
}

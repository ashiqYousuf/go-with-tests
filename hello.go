package main

import "fmt"

const HELLO = "Hello, "

func Hello(name string) string {
	// if name == "" {
	// 	name = "World"
	// }
	return fmt.Sprintf("%s%s", HELLO, name)
}

func main() {
	fmt.Println(Hello("Brain"))
}

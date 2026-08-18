package main

import "fmt"

func createGreeter(language string) func(string) string {
	if language == "Spanish" {
		return func(name string) string {
			return "Hola, " + name + "!"
		}
	}

	return func(name string) string {
		return "Vastagana huyee, " + name + "!"
	}
}

func main() {
	spanishTool := createGreeter("Spanish")

	englishTool := createGreeter("English")

	fmt.Println(spanishTool("Alex"))
	fmt.Println(englishTool("Chacha"))
}

package main

import (
	"fmt"
	
	"greeting"
)

func main() {
	// Get a greeting message and print it.
	message := greeting.Hello("World")
	fmt.Println(message)
}
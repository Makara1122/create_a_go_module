package main

import (
	"fmt"
	"log"
	"greeting"
)

func main() {
	// Set properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message,err := greeting.Hello("")
	// If an error was returned, print it to the console and 
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console
	fmt.Println(message)
}
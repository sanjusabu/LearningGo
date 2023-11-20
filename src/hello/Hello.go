package main

import (
	"fmt"
	"log"

	// "../greetings"
	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("Sanju")
	names := []string{"Sanju", "Yash", "Madhur", "Pranjay", "Abhay"}
	maps, errors := greetings.MultiHello(names)
	if err != nil || errors != nil {
		log.Fatal(err, errors)
	}
	fmt.Println(message)
	fmt.Println(maps)

}

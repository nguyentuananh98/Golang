package main

import (
	"awesomeProject/com/example/greeting"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Tuan Anh ", "Tuan Em "}

	messages, err := greeting.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}

package gosample

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	ID      int
	Name    string
	Email   string
	Age     int
	Address string
	memo    string
}

func Write() {
	person := &Person{
		ID:      1,
		Name:    "Writer",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
	}

	file, err := os.Create("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}
}

func Read() {
	file, err := os.Open("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var person Person
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(person)
	fmt.Print(string("\n"))
}

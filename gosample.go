package gosample

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	ID      int
	Name    string
	Email   string
	Age     int
	Address string
	memo    string
}

func ToJson() {
	person := &Person{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
	}
	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b)) // to Json string

	var personb Person
	res := json.Unmarshal(b, &personb)
	if res != nil {
		log.Fatal(err)
	}

	fmt.Println(personb) // to Structure
}

package gosample

import (
	"encoding/json"
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

	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString("Write String must be a byte[]\n")
	_, err = file.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}

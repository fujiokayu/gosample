package gosample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	file, _ := os.Open("./file.txt")
	message, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(message)
	fmt.Print(string("\n"))
}

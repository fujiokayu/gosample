package gosample

import (
	"fmt"
	"io/ioutil"
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

func Write() {
	message := []byte("hello world\n")
	// 3rd arg is permission
	err := ioutil.WriteFile("./file.txt", message, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func Read() {
	message, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(message))
	fmt.Print(string("\n"))
}
